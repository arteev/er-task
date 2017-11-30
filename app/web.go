package app

import (
	"html/template"
	"net/http"
	"sync"
)

// Стартовая страница просмотр истории проката автомобилей
func (a *App) Index(w http.ResponseWriter, r *http.Request) {
	templs.ExecuteTemplate(w, "index.gohtml", nil)
}

var muTemp sync.Mutex

func (a *App) AutoReloadTemplates(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		muTemp.Lock()
		defer muTemp.Unlock()
		templs = template.Must(template.ParseGlob("_template/*"))
		fn(w, r)
	}
}
