package app

import (
	"encoding/json"
	"log"
	"net/http"
)

//todo: как сделать в mux?
func ErrorHandler(fn func(http.ResponseWriter, *http.Request) (int, error)) func(http.ResponseWriter, *http.Request) {
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
