package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/arteev/er-task/src/storage"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

//Handler для трекинга ТС
func Tracking(w http.ResponseWriter, r *http.Request) (int, error) {
	db := context.Get(r, "storage").(storage.Storage)
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

	if err = db.Track(carnum, x, y); err != nil {
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
	return http.StatusOK, nil
}
