package app

import (
	"encoding/json"
	"net/http"

	"github.com/arteev/er-task/model"
	"github.com/gorilla/mux"
)

//RentJournal - handler for api. Returns list of the rental cars
func (a *App) RentJournal(w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	rncar := vars["rn"]
	rds, err := a.db.GetRentJornal(rncar)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(&model.RentDataResponse{
		Response: model.Response{
			Message:     "success",
			ContentType: "rentjournal",
		},
		Data: rds,
	})
	return http.StatusOK, nil
}
