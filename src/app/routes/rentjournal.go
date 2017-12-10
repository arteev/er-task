package routes

import (
	"fmt"
	"net/http"

	"github.com/arteev/er-task/src/model"
	"github.com/arteev/er-task/src/storage"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

//RentJournal - handler for api. Returns list of the rental cars
func RentJournal(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	db := context.Get(r, "storage").(storage.Storage)
	vars := mux.Vars(r)
	rncar := vars["rn"]
	rds, err := db.GetRentJornal(rncar)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if (rncar != "") && (len(rds) == 0) {
		return nil, http.StatusNotFound, fmt.Errorf("Car %q not found", rncar)
	}
	return &model.RentDataResponse{
		Response: model.Response{
			Message:     "success",
			ContentType: "rentjournal",
		},
		Data: rds,
	}, http.StatusOK, nil
}
