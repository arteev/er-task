package app

import (
	"net/http"

	"github.com/arteev/er-task/storage"

	"github.com/gorilla/mux"
)

//App - Application
type App struct {
	db               storage.Storage
	routes           *mux.Router
	connectionString string
	preroutes        []route
}

func (a *App) init() {
	a.db = storage.GetStorage()
	a.db.Init(a.connectionString)
	//ADD ROUTE HERE!
	a.preroutes = []route{
		{
			Path:    "/tracking/{car}/{x}/{y}",
			Methods: []string{"PUT"},
			Handler: ErrorHandler(a.Tracking),
		},
	}
	a.regroutes()
}

func (a *App) regroutes() {
	a.routes = mux.NewRouter()
	subrouter := a.routes.PathPrefix("/api/v1/").Subrouter()
	for _, r := range a.preroutes {
		rnew := subrouter.HandleFunc(r.Path, ErrorHandler(a.Tracking))
		if len(r.Methods) != 0 {
			rnew.Methods(r.Methods...)
		}
	}
	http.Handle("/", a.routes)
}

//Run run application. Retruns  a error when failure
func (a *App) Run(addr, connection string) error {
	a.connectionString = connection
	a.init()
	//TODO: host from env or flag
	//TODO: роуты в отдельный файл

	err := http.ListenAndServe(addr, nil)

	return err
}
