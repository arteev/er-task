package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arteev/er-task/model"
)

//TODO: test it
//TODO: do it
//Departments - handler for api. Returns list of the departments
func (a *App) Departments(w http.ResponseWriter, r *http.Request) (int, error) {
	/*cars, err := a.db.GetCars()
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, err
	}*/
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewEncoder(w).Encode(&model.DepartmentsResponse{
		Response: model.Response{
			Message:     "success",
			ContentType: "departments",
		},
		Data: []model.Department{{0, "Амстердам"}, {0, "Афины"}},
	})
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
