package routes

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Route struct {
	IsAPI   bool
	Path    string
	Methods []string
	Handler http.HandlerFunc
}

//Routes web appication
var Routes = []Route{
	{
		IsAPI:   true,
		Path:    "/cars",
		Methods: []string{"GET"},
		Handler: ErrorHandler(Cars),
	},
	{
		IsAPI:   true,
		Path:    "/car/{rn}",
		Methods: []string{"GET"},
		Handler: ErrorHandler(CarInfo),
	},
	{
		IsAPI:   true,
		Path:    "/departments",
		Methods: []string{"GET"},
		Handler: ErrorHandler(Departments),
	},
	{
		IsAPI:   true,
		Path:    "/stats/deps/model",
		Methods: []string{"GET"},
		Handler: ErrorHandler(StatsByModel),
	},
	{
		IsAPI:   true,
		Path:    "/stats/deps/type",
		Methods: []string{"GET"},
		Handler: ErrorHandler(StatsByType),
	},
	{
		IsAPI:   true,
		Path:    "/rent",
		Methods: []string{"POST"},
		Handler: ErrorHandler(Rent),
	},
	{
		IsAPI:   true,
		Path:    "/return",
		Methods: []string{"POST"},
		Handler: ErrorHandler(Return),
	},
	{
		IsAPI:   true,
		Path:    "/rentjournal",
		Methods: []string{"GET"},
		Handler: ErrorHandler(RentJournal),
	},
	{
		IsAPI:   true,
		Path:    "/rentjournal/{rn}",
		Methods: []string{"GET"},
		Handler: ErrorHandler(RentJournal),
	},
	{
		IsAPI:   true,
		Path:    "/tracking/{car}/{x}/{y}",
		Methods: []string{"PUT"},
		Handler: ErrorHandler(Tracking),
	},
}

func GetHandler(addon []Route, middlewares ...func(http.HandlerFunc) http.HandlerFunc) (*mux.Router, http.Handler) {
	rt := mux.NewRouter()
	rt.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	subrouter := rt.PathPrefix("/api/v1/").Subrouter()
	regroutes := append(addon, Routes...)
	for _, r := range regroutes {

		handler := r.Handler
		for _, h := range middlewares {
			handler = h(handler)
		}

		if r.IsAPI {
			rnew := subrouter.HandleFunc(r.Path, handler)
			if len(r.Methods) != 0 {
				rnew.Methods(r.Methods...)
			}
		} else {
			rt.HandleFunc(r.Path, handler)
		}
	}
	logroutes := handlers.LoggingHandler(os.Stdout, rt)
	return rt, logroutes
}
