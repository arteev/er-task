package routes

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/arteev/er-task/src/storage"

	"github.com/arteev/er-task/src/model"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

//Cars - handler for api. Returns list of the cars
func Cars(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	db := context.Get(r, "storage").(storage.Storage)
	cars, err := db.GetCars()
	if err != nil {
		log.Println(err)
		return nil, http.StatusInternalServerError, err
	}
	return &model.CarsResponse{
		Response: model.Response{
			Message:     "success",
			ContentType: "cars",
		},
		Data: cars,
	}, http.StatusOK, nil
}

func CarInfo(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	db := context.Get(r, "storage").(storage.Storage)
	vars := mux.Vars(r)
	rn := vars["rn"]
	if rn == "" {
		return nil, http.StatusBadRequest, errors.New("Не задан рег.номер ТС")
	}
	car, err := db.GetCarInfo(rn)
	if err != nil {
		if sql.ErrNoRows == err {
			return nil, http.StatusNotFound, fmt.Errorf("Car %q not found", rn)
		}
		return nil, http.StatusInternalServerError, err
	}
	return &model.CarInfoResponse{
		Response: model.Response{
			Message:     "success",
			ContentType: "carinfo",
		},
		Data: *car,
	}, http.StatusOK, nil
}
