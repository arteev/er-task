package app

import (
	"html/template"
	"log"
	"net/http"

	"github.com/arteev/er-task/src/app/routes"
	"github.com/arteev/er-task/src/cache"
	"github.com/arteev/er-task/src/storage"
	"github.com/arteev/er-task/src/ws"
)

var templs *template.Template

func init() {
	var err error
	templs, err = template.ParseGlob("./templates/*")
	if err != nil {
		//log.Fatal(err) TODO:fix it when testing
	}
}

//App - Application
type app struct {
	db               storage.Storage
	connectionString string
	redis            string
	preroutes        []routes.Route
}

func (a *app) cachehit(name string, hit bool) {
	if hit {
		log.Printf("Cache hit: %s", name)
	} else {
		log.Printf("Cache missing: %s", name)
	}
}

func (a *app) initStorage() error {
	if a.redis != "" {
		a.db = cache.NewCacheRedis(a.redis, storage.GetStorage(), a.cachehit)
	} else {
		a.db = storage.GetStorage()
	}
	err := a.db.Init(a.connectionString, true)
	return err
}

func (a *app) Init() error {
	err := a.initStorage()
	if err != nil {
		return err
	}
	a.initroutes()
	return nil
}
func (a *app) initroutes() http.Handler {
	a.preroutes = []routes.Route{
		// websocket
		{
			IsAPI:   false,
			Path:    "/ws",
			Handler: routes.ErrorHandler(ws.GetServer(a.db.Notify()).Handler),
		},
	}
	_, handler := routes.GetHandler(a.preroutes, a.ContextInit)
	return handler
}

//Run run application. Retruns a error when failure
func Run(addr, connection string, redis string) error {
	a := &app{}
	a.connectionString = connection
	a.redis = redis
	err := a.Init()
	if err != nil {
		return err
	}
	err = http.ListenAndServe(addr, a.initroutes())
	return err
}
