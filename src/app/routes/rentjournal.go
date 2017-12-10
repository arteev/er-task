package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/arteev/er-task/src/model"
	"github.com/arteev/er-task/src/storage"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

//RentJournal - handler for api. Returns list of the rental cars
func RentJournal(w http.ResponseWriter, r *http.Request) (int, error) {
	db := context.Get(r, "storage").(storage.Storage)
	vars := mux.Vars(r)
	rncar := vars["rn"]
	rds, err := db.GetRentJornal(rncar)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if (rncar != "") && (len(rds) == 0) {
		return http.StatusNotFound, fmt.Errorf("Car %q not found", rncar)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(&model.RentDataResponse{
		Response: model.Response{
			Message:     "success",
			ContentType: "rentjournal",
		},
		Data: rds,
	})
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
