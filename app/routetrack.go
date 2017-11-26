package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//TODO: mux ?? toJSON() ????

func (a *App) Index(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Index"))

	templs.Execute(w, nil)
}

//Handler для трекинга ТС
func (a *App) Tracking(w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)

	carnum := vars["car"]

	x, err := strconv.ParseFloat(vars["x"], 64)
	if err != nil {
		return http.StatusBadRequest, err
	}

	y, err := strconv.ParseFloat(vars["y"], 64)
	if err != nil {
		return http.StatusBadRequest, err
	}

	/*car, err := a.db.FindCarByID(carID)
	if err != nil {
		return http.StatusNotFound, err
	}*/

	if err := a.db.Track(carnum, x, y); err != nil {
		return http.StatusNotFound, err
	}

	//TODO: refactor this. Middeware
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(&struct {
		Message string `json:"message"`
	}{
		"success",
	})

	if err != nil {
		return http.StatusInternalServerError, err
	}
	//	w.WriteHeader(http.StatusOK)
	//	w.Write(b)
	return http.StatusOK, nil
}
