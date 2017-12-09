package app

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/arteev/er-task/cache"
	"github.com/arteev/er-task/storage"
	"github.com/arteev/er-task/ws"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var templs *template.Template

func init() {

	//TODO: dir template from env | flags
	var err error
	templs, err = template.ParseGlob("_template/*")
	if err != nil {
		///
	}

}

//App - Application
type App struct {
	db               storage.Storage
	routes           *mux.Router
	connectionString string
	preroutes        []route
}

func (a *App) init() http.Handler {

	if val, ok := os.LookupEnv("CACHE"); ok && val == "true" {
		rs, ok := os.LookupEnv("REDIS")
		if !ok {
			rs = "127.0.0.1:6379"
		}
		a.db = cache.NewCacheRedis(rs, storage.GetStorage(), func(name string, hit bool) {
			if hit {
				log.Printf("Cache hit: %s", name)
			} else {
				log.Printf("Cache missing: %s", name)
			}
		})
	} else {
		a.db = storage.GetStorage()
	}

	err := a.db.Init(a.connectionString, true)
	if err != nil {
		panic(err)
	}
	//ADD ROUTE HERE!
	a.preroutes = []route{
		{
			IsAPI:   true,
			Path:    "/tracking/{car}/{x}/{y}",
			Methods: []string{"PUT"},
			Handler: ErrorHandler(a.Tracking),
		},
		{
			IsAPI:   true,
			Path:    "/rent",
			Methods: []string{"POST"},
			Handler: ErrorHandler(a.Rent),
		},
		{
			IsAPI:   true,
			Path:    "/return",
			Methods: []string{"POST"},
			Handler: ErrorHandler(a.Return),
		},

		{
			IsAPI:   true,
			Path:    "/rentjournal",
			Methods: []string{"GET"},
			Handler: ErrorHandler(a.RentJournal),
		},
		{
			IsAPI:   true,
			Path:    "/rentjournal/{rn}",
			Methods: []string{"GET"},
			Handler: ErrorHandler(a.RentJournal),
		},
		{
			IsAPI:   true,
			Path:    "/cars",
			Methods: []string{"GET"},
			Handler: ErrorHandler(a.Cars),
		},
		{
			IsAPI:   true,
			Path:    "/departments",
			Methods: []string{"GET"},
			Handler: ErrorHandler(a.Departments),
		},
		{
			IsAPI:   true,
			Path:    "/car/{rn}",
			Methods: []string{"GET"},
			Handler: ErrorHandler(a.CarInfo),
		},

		{
			IsAPI:   true,
			Path:    "/stats/deps/model",
			Methods: []string{"GET"},
			Handler: ErrorHandler(a.StatsByModel),
		},
		{
			IsAPI:   true,
			Path:    "/stats/deps/type",
			Methods: []string{"GET"},
			Handler: ErrorHandler(a.StatsByType),
		},

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
			Handler: a.autoReloadTemplates(a.Stats),
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

func (a *App) regroutes() http.Handler {
	a.routes = mux.NewRouter()
	a.routes.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("./_static/"))))

	subrouter := a.routes.PathPrefix("/api/v1/").Subrouter()
	for _, r := range a.preroutes {
		if r.IsAPI {
			rnew := subrouter.HandleFunc(r.Path, r.Handler)
			if len(r.Methods) != 0 {
				rnew.Methods(r.Methods...)
			}
		} else {
			a.routes.HandleFunc(r.Path, r.Handler)
		}

	}
	logroutes := handlers.LoggingHandler(os.Stdout, a.routes)
	//http.Handle("/", logroutes)
	return logroutes
}

//Run run application. Retruns  a error when failure
func (a *App) Run(addr, connection string) error {

	a.connectionString = connection

	//TODO: host from env or flag
	//TODO: роуты в отдельный файл

	err := http.ListenAndServe(addr, a.init())

	return err
}
