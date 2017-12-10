package app

import (
	"net/http"
)

type route struct {
	IsAPI   bool
	Path    string
	Methods []string
	Handler http.HandlerFunc
}
