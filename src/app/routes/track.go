package routes

import (
	"net/http"
	"strconv"

	"github.com/arteev/er-task/src/storage"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

/*Exp1:*/
/*
type data struct {
	car string
	x   float64
	y   float64
}
var datachan = make(chan data, 100000)
var once sync.Once
func DoData(db storage.Storage) func() {

	return func() {
		go func() {
			log.Println("Worker tracker started")
			for {
				newdata := <-datachan
				if err := db.Track(newdata.car, newdata.x, newdata.y); err != nil {
					log.Println("Track err: %s", err)
				}
			}
		}()
	}
}*/

//Tracking Handler для трекинга ТС
func Tracking(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	//TODO : //через worker????
	db := context.Get(r, "storage").(storage.Storage)

	/*Exp1:*/
	/*
		once.Do(DoData(db))
	*/

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

	/*Exp1:*/
	/*
		go func() {
			datachan <- data{carnum, x, y}
		}()
		// закомментировать ниже db.Track vvvvvvv
	*/

	if err = db.Track(carnum, x, y); err != nil {
		return nil, http.StatusNotFound, err
	}

	return &struct {
		Message string `json:"message"`
	}{
		"success",
	}, http.StatusOK, nil
}
