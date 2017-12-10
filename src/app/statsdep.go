package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arteev/er-task/src/model"
)

//TODO: test StatsByModel
func (a *App) StatsByModel(w http.ResponseWriter, r *http.Request) (int, error) {
	stats, err := a.db.GetStatsByModel()
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, err
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(&model.StatsDepartmentoResponse{
		Response: model.Response{
			Message:     "success",
			ContentType: "statsbymodel",
		},
		Data: stats,
	})
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

//TODO: test StatsByType
func (a *App) StatsByType(w http.ResponseWriter, r *http.Request) (int, error) {
	stats, err := a.db.GetStatsByType()
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, err
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(&model.StatsDepartmentoResponse{
		Response: model.Response{
			Message:     "success",
			ContentType: "statsbytype",
		},
		Data: stats,
	})
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
