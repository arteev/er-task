package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//TODO: mux ?? toJSON() ????

//Handler для трекинга ТС
func (a *App) Tracking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	carID, err := strconv.Atoi(vars["car"])
	if err != nil {
		errorResponseError(err, w, r, http.StatusBadRequest)
	}

	//x := vars["x"]
	//y := vars["y"]

	_, err = a.db.FindCarByID(carID)
	if err != nil {
		errorResponseError(err, w, r, http.StatusNotFound)
		return
	}

	//TODO: refactor this
	b, err := json.Marshal(&struct {
		Message string `json:"message"`
	}{"success"})
	if err != nil {
		errorResponseError(err, w, r, http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)

	//	a.db.Track()

	/*b, err := json.Marshal(&vars)
	if err != nil {
		panic(err)
	}*/
	//errorHandler(vars, w, r, http.StatusBadRequest)
}

func errorResponseError(err error, w http.ResponseWriter, r *http.Request, status int) {
	s := struct{ error string }{err.Error()}
	errorHandler(s, w, r, status)
}

func errorHandler(data interface{}, w http.ResponseWriter, r *http.Request, status int) {
	//TODO: переделать ошибки м/б через миддлваре
	w.WriteHeader(status)
	b, err := json.Marshal(&data)
	if err != nil {
		log.Printf("could not marshal into errorHandler: %s", err)
	}
	w.Write(b)
}
