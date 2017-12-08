package cache

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/arteev/er-task/model"
	"github.com/arteev/er-task/storage"

	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
)

type rediscache struct {
	mu sync.RWMutex
	storage.Storage

	codec    *cache.Codec
	conn     *redis.Client
	callback CacheHitMissCallback

	getcars []model.Car
}
type CacheHitMissCallback = func(name string, hit bool)

//NewCacheRedis - возвращает обертку над хранилищем и кэширует данные
func NewCacheRedis(addr string, s storage.Storage, f CacheHitMissCallback) storage.Storage {
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

func (r *rediscache) Rent(rn string, dep string, agn string) error {
	keycache := "rentjournal"
	err := r.Storage.Rent(rn, dep, agn)
	if err != nil {
		return err
	}
	r.codec.Delete(keycache)
	r.codec.Delete(keycache + rn)

	return nil
}

func (r *rediscache) Return(rn string, dep string, agn string) error {
	keycache := "rentjournal"
	err := r.Storage.Return(rn, dep, agn)
	if err != nil {
		return err
	}
	r.codec.Delete(keycache)
	r.codec.Delete(keycache + rn)
	return nil
}

func (r *rediscache) GetRentJornal(rn string) ([]model.RentData, error) {
	keycache := "rentjournal"
	rds := make([]model.RentData, 0)

	checkkey := keycache
	if rn != "" {
		checkkey += rn
	}
	if r.codec.Exists(checkkey) {
		if err := r.codec.Get(checkkey, &rds); err == nil {
			r.docallback(checkkey, true)
			return rds, nil
		}
	}

	r.docallback(keycache, false)
	if rn != "" {
		r.docallback(keycache+rn, false)
	}
	rds, err := r.Storage.GetRentJornal(rn)
	if err != nil {
		return nil, err
	}

	err = r.codec.Set(&cache.Item{
		Key:        keycache,
		Object:     rds,
		Expiration: time.Hour,
	})
	if err != nil {
		log.Println(err)
	}

	if rn != "" {
		err = r.codec.Set(&cache.Item{
			Key:        keycache + rn,
			Object:     rds,
			Expiration: time.Hour,
		})
		if err != nil {
			log.Println(err)
		}
	}

	return rds, nil

}

func (r *rediscache) GetDepartments() ([]model.Department, error) {
	key := "deps"
	if r.codec.Exists(key) {
		deps := make([]model.Department, 0)
		if err := r.codec.Get(key, &deps); err == nil {
			r.docallback(key, true)
			return deps, nil
		}
	}
	r.docallback(key, false)
	deps, err := r.Storage.GetDepartments()
	if err != nil {
		return nil, err
	}
	err = r.codec.Set(&cache.Item{
		Key:        key,
		Object:     deps,
		Expiration: time.Hour,
	})
	if err != nil {
		log.Println(err)
	}
	return deps, err
}

func (r *rediscache) Done() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.conn.Close()
	return r.Storage.Done()
}

func (r *rediscache) docallback(name string, hit bool) {
	if r.callback != nil {
		r.callback(name, hit)
	}
}
