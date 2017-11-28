package app

import (
	"html/template"
	"net/http"
	"os"
	"path"

	"github.com/arteev/er-task/storage"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var templs *template.Template

func init() {
	var err error
	//TODO: dir template from env | flags
	templs, err = template.ParseGlob(path.Join("./_template", "*.gohtml"))
	if err != nil {
		//panic(err)
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
	a.db = storage.GetStorage()
	err := a.db.Init(a.connectionString)
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
			IsAPI:   false,
			Path:    "/",
			Methods: []string{"GET"},
			Handler: a.Index,
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
