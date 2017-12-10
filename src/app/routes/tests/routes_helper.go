package tests

import (
	"net/http"

	"github.com/arteev/er-task/src/app/routes"
	"github.com/arteev/er-task/src/storage"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

//GetRoutes для тестирование встраивает в контекст storage
func GetRoutes(db storage.Storage) *mux.Router {
	ctxdb := func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			context.Set(r, "storage", db)
			f(w, r)
		}
	}
	routes, _ := routes.GetHandler(nil, ctxdb)
	return routes
}
