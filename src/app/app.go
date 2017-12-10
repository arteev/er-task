package app

import (
	"html/template"
	"log"
	"net/http"
	"os"

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
type App struct {
	db storage.Storage
	//routes           *mux.Router
	connectionString string
	preroutes        []routes.Route
}

func (a *App) cachehit(name string, hit bool) {
	if hit {
		log.Printf("Cache hit: %s", name)
	} else {
		log.Printf("Cache missing: %s", name)
	}
}

func (a *App) initStorage() error {
	if val, ok := os.LookupEnv("CACHE"); ok && val == "true" {
		if rs, ok := os.LookupEnv("REDIS"); !ok {
			rs = "127.0.0.1:6379"
		} else {
			a.db = cache.NewCacheRedis(rs, storage.GetStorage(), a.cachehit)
		}
	} else {
		a.db = storage.GetStorage()
	}
	err := a.db.Init(a.connectionString, true)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) Init() error {
	err := a.initStorage()
	if err != nil {
		return err
	}
	a.initroutes()
	return nil
}
func (a *App) initroutes() http.Handler {
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
			Handler: ws.GetServer(a.db.Notify()).Handler,
		},
	}
	return a.regroutes()
}

func (a *App) regroutes() (handler http.Handler) {
	/*a.Routes = mux.NewRouter()
	a.Routes.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	subrouter := a.Routes.PathPrefix("/api/v1/").Subrouter()
	regroutes := a.preroutes
	regroutes = append(regroutes, routes.Routes...)
	for _, r := range regroutes {
		if r.IsAPI {
			rnew := subrouter.HandleFunc(r.Path, a.ContextInit(r.Handler))
			if len(r.Methods) != 0 {
				rnew.Methods(r.Methods...)
			}
		} else {
			a.Routes.HandleFunc(r.Path, a.ContextInit(r.Handler))
		}
	}
	logroutes := handlers.LoggingHandler(os.Stdout, a.Routes)
	return logroutes*/
	_, handler = routes.GetHandler(a.preroutes, a.ContextInit)
	return
}

//Run run application. Retruns  a error when failure
func (a *App) Run(addr, connection string) error {
	a.connectionString = connection
	err := a.Init()
	if err != nil {
		return err
	}
	err = http.ListenAndServe(addr, a.initroutes())
	return err
}
