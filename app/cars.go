package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arteev/er-task/model"
)

//Cars - handler for api. Returns list of the cars
func (a *App) Cars(w http.ResponseWriter, r *http.Request) (int, error) {
	cars, err := a.db.GetCars()
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, err
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(&model.CarsResponse{
		Response: model.Response{
			Message:     "success",
			ContentType: "cars",
		},
		Data: cars,
	})
	return http.StatusOK, nil
}
