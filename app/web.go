package app

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type cardata struct {
	RN string
}

//Index - /index Стартовая страница просмотр истории проката автомобилей
func (a *App) Index(w http.ResponseWriter, r *http.Request) {
	err := templs.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

//Car - обработчик для страницы /car просмотр истории и действия с ТС
func (a *App) Car(w http.ResponseWriter, r *http.Request) {
	var car cardata
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

func (a *App) Stats(w http.ResponseWriter, r *http.Request) {
	err := templs.ExecuteTemplate(w, "stats.gohtml", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

var muTemp sync.Mutex

/*
func (a *App) autoReloadTemplates(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		muTemp.Lock()
		defer muTemp.Unlock()
		templs = template.Must(template.ParseGlob("_template/*"))
		fn(w, r)
	}
}*/
