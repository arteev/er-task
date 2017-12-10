package routes

import (
	"log"
	"net/http"

	"github.com/arteev/er-task/src/model"
	"github.com/arteev/er-task/src/storage"
	"github.com/gorilla/context"
)

//TODO: test StatsByModel
func StatsByModel(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	db := context.Get(r, "storage").(storage.Storage)
	stats, err := db.GetStatsByModel()
	if err != nil {
		log.Println(err)
		return nil, http.StatusInternalServerError, err
	}
	return &model.StatsDepartmentoResponse{
		Response: model.Response{
			Message:     "success",
			ContentType: "statsbymodel",
		},
		Data: stats,
	}, http.StatusOK, nil
}

//TODO: test StatsByType
func StatsByType(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	db := context.Get(r, "storage").(storage.Storage)
	stats, err := db.GetStatsByType()
	if err != nil {
		log.Println(err)
		return nil, http.StatusInternalServerError, err
	}
	return &model.StatsDepartmentoResponse{
		Response: model.Response{
			Message:     "success",
			ContentType: "statsbytype",
		},
		Data: stats,
	}, http.StatusOK, nil
}
