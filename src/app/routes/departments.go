package routes

import (
	"log"
	"net/http"

	"github.com/arteev/er-task/src/model"
	"github.com/arteev/er-task/src/storage"
	"github.com/gorilla/context"
)

//Departments - handler for api. Returns list of the departments
func Departments(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	db := context.Get(r, "storage").(storage.Storage)
	deps, err := db.GetDepartments()
	if err != nil {
		log.Println(err)
		return nil, http.StatusInternalServerError, err
	}
	return &model.DepartmentsResponse{
		Response: model.Response{
			Message:     "success",
			ContentType: "departments",
		},
		Data: deps,
	}, http.StatusOK, nil
}
