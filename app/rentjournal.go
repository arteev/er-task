package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arteev/er-task/model"
)

//
func (a *App) RentJournal(w http.ResponseWriter, r *http.Request) (int, error) {
	rds, err := a.db.GetRentJornal()
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, err
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(
		&struct {
			Message string           `json:"message"`
			Data    []model.RentData `json:"data"`
		}{
			"success",
			rds,
		})

	return http.StatusOK, nil
}
