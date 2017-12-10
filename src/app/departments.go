package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arteev/er-task/src/model"
)

//TODO: test it
//Departments - handler for api. Returns list of the departments
func (a *App) Departments(w http.ResponseWriter, r *http.Request) (int, error) {
	deps, err := a.db.GetDepartments()
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
