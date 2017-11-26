package app

import (
	"net/http"

	"github.com/arteev/er-task/storage"

	"github.com/gorilla/mux"
)

type App struct {
	db               storage.Storage
	routes           *mux.Router
	connectionString string
}

func (a *App) init() {
	a.db = storage.GetStorage()
	a.db.Init(a.connectionString)
	a.routes = mux.NewRouter()

	a.routes.HandleFunc("/api/v1/tracking/{car}/{x}/{y}", ErrorHandler(a.Tracking)).Methods("PUT")
	//.Methods("GET")
	http.Handle("/", a.routes)

}
func (a *App) Run(addr, connection string) error {
	a.connectionString = connection
	a.init()
	//TODO: host from env or flag

	//TODO: роуты в отдельный файл

	err := http.ListenAndServe(addr, nil)

	return err
}
