package app

import (
	"fmt"
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
	db storage.Storage
	//routes           *mux.Router
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
		//render routes
		{
			IsAPI:   false,
			Path:    "/",
			Methods: []string{"GET"},
			Handler: a.Index,
		},
		{
			IsAPI:   false,
			Path:    "/car",
			Methods: []string{"GET"},
			Handler: a.Car,
		},
		{
			IsAPI:   false,
			Path:    "/stats",
			Methods: []string{"GET"},
			Handler: a.Stats,
		},
		{
			IsAPI:   false,
			Path:    "/car/{rn}",
			Methods: []string{"GET"},
			Handler: a.Car,
		},
		// websocket
		{
			IsAPI:   false,
			Path:    "/ws",
			Methods: []string{},
			Handler: routes.ErrorHandler(ws.GetServer(a.db.Notify()).Handler),
		},
		{
			IsAPI:   true,
			Path:    "/test",
			Handler: routes.JSONHandler(fn),
		},
	}
	_, handler := routes.GetHandler(a.preroutes, a.ContextInit)
	return handler
}

func fn(q http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	//return &struct{ MessageNew string }{"TEST"}, http.StatusOK, nil
	return &struct{ MessageNew string }{"TEST"}, http.StatusInternalServerError, fmt.Errorf("this is error")
}

//Run run application. Retruns  a error when failure
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
