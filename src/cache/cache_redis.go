package cache

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/arteev/er-task/src/model"
	"github.com/arteev/er-task/src/storage"

	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
)

type rediscache struct {
	mu sync.RWMutex
	storage.Storage

	codec    *cache.Codec
	conn     *redis.Client
	callback HitMissCallback

	getcars []model.Car
}

//HitMissCallback callback function. Для подсчета статистики
type HitMissCallback func(name string, hit bool)

//NewCacheRedis - возвращает обертку над хранилищем и кэширует данные
func NewCacheRedis(addr string, s storage.Storage, f HitMissCallback) storage.Storage {
	return &rediscache{
		Storage:  s,
		callback: f,
		codec: &cache.Codec{
			Redis: redis.NewClient(&redis.Options{
				Addr: addr,
			}),
			Marshal: func(v interface{}) ([]byte, error) {
				return json.Marshal(v)
			},
			Unmarshal: func(b []byte, v interface{}) error {
				return json.Unmarshal(b, v)
			},
		},
	}
}

func (r *rediscache) Rent(rn string, dep string, agn string) (int, error) {
	keycache := "rentjournal"
	id, err := r.Storage.Rent(rn, dep, agn)
	if err != nil {
		return 0, err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.codec.Delete(keycache + rn)
	r.codec.Delete(keycache)
	return id, nil
}

func (r *rediscache) Return(rn string, dep string, agn string) (int, error) {
	keycache := "rentjournal"
	id, err := r.Storage.Return(rn, dep, agn)
	if err != nil {
		return 0, err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.codec.Delete(keycache + rn)
	r.codec.Delete(keycache)
	return id, nil
}

func (r *rediscache) GetRentJornal(rn string) ([]model.RentData, error) {
	keycache := "rentjournal"
	if rn != "" {
		keycache += rn
	}
	rds := make([]model.RentData, 0)

	r.mu.Lock()
	if err := r.codec.Get(keycache, &rds); err == nil {
		r.mu.Unlock()
		r.doCallback(keycache, true)
		return rds, nil
	}
	r.mu.Unlock()

	r.doCallback(keycache, false)

	rds, err := r.Storage.GetRentJornal(rn)
	if err != nil {
		return nil, err
	}

	r.mu.Lock()
	err = r.codec.Set(&cache.Item{
		Key:        keycache,
		Object:     rds,
		Expiration: time.Hour,
	})
	r.mu.Unlock()
	if err != nil {
		log.Println(err)
	}

	return rds, nil

}

func (r *rediscache) GetDepartments() ([]model.Department, error) {
	key := "deps"

	deps := make([]model.Department, 0)
	r.mu.Lock()
	if err := r.codec.Get(key, &deps); err == nil {
		r.mu.Unlock()
		r.doCallback(key, true)
		return deps, nil
	}
	r.mu.Unlock()

	r.doCallback(key, false)
	deps, err := r.Storage.GetDepartments()
	if err != nil {
		return nil, err
	}
	r.mu.Lock()
	err = r.codec.Set(&cache.Item{
		Key:        key,
		Object:     deps,
		Expiration: time.Hour,
	})
	r.mu.Unlock()
	if err != nil {
		log.Println(err)
	}
	return deps, nil
}

func (r *rediscache) Done() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.conn.Close()
	return r.Storage.Done()
}

func (r *rediscache) doCallback(name string, hit bool) {
	if r.callback != nil {
		r.callback(name, hit)
	}
}
