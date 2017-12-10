package app

import (
	"html/template"
	"net/http"
	"sync"

	"github.com/gorilla/context"
)

var muTemp sync.Mutex

func (a *app) autoReloadTemplates(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		muTemp.Lock()
		defer muTemp.Unlock()
		templs = template.Must(template.ParseGlob("templates/*"))
		fn(w, r)
	}
}

func (a *app) ContextInit(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, "storage", a.db)
		context.Set(r, "templates", templs)
		f(w, r)
	}
}
