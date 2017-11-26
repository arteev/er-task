package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//TODO: mux ?? toJSON() ????

//Handler для трекинга ТС
func (a *App) Tracking(w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)

	carID, err := strconv.Atoi(vars["car"])
	if err != nil {
		return http.StatusBadRequest, err
	}

	x, err := strconv.ParseFloat(vars["x"], 64)
	if err != nil {
		return http.StatusBadRequest, err
	}

	y, err := strconv.ParseFloat(vars["x"], 64)
	if err != nil {
		return http.StatusBadRequest, err
	}

	car, err := a.db.FindCarByID(carID)
	if err != nil {
		return http.StatusNotFound, err
	}

	if err := a.db.Track(*car, x, y); err != nil {
		return http.StatusNotFound, err
	}

	//TODO: refactor this. Middeware
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(&struct {
		Message string `json:"message"`
	}{"success"})

	if err != nil {
		return http.StatusInternalServerError, err
	}
	//	w.WriteHeader(http.StatusOK)
	//	w.Write(b)
	return http.StatusOK, nil
}
