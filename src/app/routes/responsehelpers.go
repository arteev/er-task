package routes

import (
	"encoding/json"
	"log"
	"net/http"
)

//ErrorHandler - returns error in json
func ErrorHandler(fn func(http.ResponseWriter, *http.Request) (int, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status, err := fn(w, r)
		if err != nil {
			s := struct {
				Error string `json:"error"`
			}{err.Error()}
			b, err := json.Marshal(&s)
			if err != nil {
				log.Printf("ErrorHandler: could not marshal %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(status)
			w.Write(b)
		}
	}
}

//JSONHandler - returns response in json or error
func JSONHandler(next func(http.ResponseWriter, *http.Request) (interface{}, int, error)) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var (
			v      interface{}
			status int
			err    error
		)
		merror := ErrorHandler(func(w1 http.ResponseWriter, r1 *http.Request) (int, error) {
			v, status, err = next(w1, r1)
			return status, err
		})
		merror(w, r)
		if err == nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			err := json.NewEncoder(w).Encode(v)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
		}
	}
}
