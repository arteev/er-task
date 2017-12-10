package routes

import (
	"html/template"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

//Index - /index Стартовая страница просмотр истории проката автомобилей
func Index(w http.ResponseWriter, r *http.Request) {
	templs := context.Get(r, "templates").(*template.Template)
	err := templs.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

//Car - обработчик для страницы /car просмотр истории и действия с ТС
func Car(w http.ResponseWriter, r *http.Request) {
	var car struct {
		RN string
	}
	vars := mux.Vars(r)
	if rn, ok := vars["rn"]; ok {
		car.RN = rn
	}
	templs := context.Get(r, "templates").(*template.Template)
	err := templs.ExecuteTemplate(w, "car.gohtml", car)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

//Stats - обработчик для страницы просмотра статистики
func Stats(w http.ResponseWriter, r *http.Request) {
	templs := context.Get(r, "templates").(*template.Template)
	err := templs.ExecuteTemplate(w, "stats.gohtml", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}
