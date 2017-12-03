package app

import (
	"html/template"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type cardata struct {
	RN string
}

// Стартовая страница просмотр истории проката автомобилей
func (a *App) Index(w http.ResponseWriter, r *http.Request) {
	templs.ExecuteTemplate(w, "index.gohtml", nil)
}

//TODO: //test api with rn
func (a *App) Car(w http.ResponseWriter, r *http.Request) {
	var car cardata
	vars := mux.Vars(r)
	if rn, ok := vars["rn"]; ok {
		car.RN = rn
	}
	//TODO: читать var
	templs.ExecuteTemplate(w, "car.gohtml", car)
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
