package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arteev/er-task/model"
)

//RentJournal - handler for api. Returns list of the rental cars
func (a *App) RentJournal(w http.ResponseWriter, r *http.Request) (int, error) {
	rds, err := a.db.GetRentJornal()
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, err
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(&model.RentDataResponse{Message: "success", Data: rds})
	return http.StatusOK, nil
}
