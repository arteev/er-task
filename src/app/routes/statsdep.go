package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arteev/er-task/src/model"
	"github.com/arteev/er-task/src/storage"
	"github.com/gorilla/context"
)

//TODO: test StatsByModel
func StatsByModel(w http.ResponseWriter, r *http.Request) (int, error) {
	db := context.Get(r, "storage").(storage.Storage)
	stats, err := db.GetStatsByModel()
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
func StatsByType(w http.ResponseWriter, r *http.Request) (int, error) {
	db := context.Get(r, "storage").(storage.Storage)
	stats, err := db.GetStatsByType()
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
