package routes

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/arteev/er-task/src/storage"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

/*Exp1:*/
type dataTrack struct {
	car string
	x   float64
	y   float64
}

var datachan = make(chan dataTrack, 100000)
var once sync.Once

/*Exp1:*/
func processTrack(db storage.Storage) func() {
	return func() {
		go func() {
			log.Println("Worker tracker started")
			for {
				newdata := <-datachan
				if err := db.Track(newdata.car, newdata.x, newdata.y); err != nil {
					log.Printf("Track err: %s", err)
				}
			}
		}()
	}
}

/*Exp1:*/
//Tracking Handler для трекинга ТС
func TrackingExp(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	db := context.Get(r, "storage").(storage.Storage)
	once.Do(processTrack(db))
	vars := mux.Vars(r)
	carnum := vars["car"]
	x, err := strconv.ParseFloat(vars["x"], 64)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	y, err := strconv.ParseFloat(vars["y"], 64)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	go func() {
		datachan <- dataTrack{carnum, x, y}
	}()

	return &struct {
		Message string `json:"message"`
	}{
		"success",
	}, http.StatusOK, nil
}

//Tracking Handler для трекинга ТС
func Tracking(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	db := context.Get(r, "storage").(storage.Storage)
	vars := mux.Vars(r)
	carnum := vars["car"]
	x, err := strconv.ParseFloat(vars["x"], 64)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	y, err := strconv.ParseFloat(vars["y"], 64)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	if err = db.Track(carnum, x, y); err != nil {
		return nil, http.StatusNotFound, err
	}
	return &struct {
		Message string `json:"message"`
	}{
		"success",
	}, http.StatusOK, nil
}
