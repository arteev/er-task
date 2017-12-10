package routes

import (
	"net/http"
	"strconv"

	"github.com/arteev/er-task/src/storage"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

//Tracking Handler для трекинга ТС
func Tracking(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	//TODO : //через worker????
	db := context.Get(r, "storage").(storage.Storage)
	vars := mux.Vars(r)
	carnum := vars["car"]
	x, err := strconv.ParseFloat(vars["x"], 64)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	y, err := strconv.ParseFloat(vars["y"], 64)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	if err = db.Track(carnum, x, y); err != nil {
		return nil, http.StatusNotFound, err
	}

	return &struct {
		Message string `json:"message"`
	}{
		"success",
	}, http.StatusOK, nil
}
