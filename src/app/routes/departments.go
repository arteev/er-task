package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arteev/er-task/src/model"
	"github.com/arteev/er-task/src/storage"
	"github.com/gorilla/context"
)

//Departments - handler for api. Returns list of the departments
func Departments(w http.ResponseWriter, r *http.Request) (int, error) {
	db := context.Get(r, "storage").(storage.Storage)
	deps, err := db.GetDepartments()
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, err
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(&model.DepartmentsResponse{
		Response: model.Response{
			Message:     "success",
			ContentType: "departments",
		},
		Data: deps,
	})
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
