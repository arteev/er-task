package app

import (
	"html/template"
	"net/http"
	"sync"

	"github.com/gorilla/context"

	"github.com/gorilla/mux"
)

//Index - /index Стартовая страница просмотр истории проката автомобилей
func (a *app) Index(w http.ResponseWriter, r *http.Request) {
	err := templs.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

//Car - обработчик для страницы /car просмотр истории и действия с ТС
func (a *app) Car(w http.ResponseWriter, r *http.Request) {
	var car struct {
		RN string
	}
	vars := mux.Vars(r)
	if rn, ok := vars["rn"]; ok {
		car.RN = rn
	}
	err := templs.ExecuteTemplate(w, "car.gohtml", car)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

func (a *app) Stats(w http.ResponseWriter, r *http.Request) {
	err := templs.ExecuteTemplate(w, "stats.gohtml", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

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
