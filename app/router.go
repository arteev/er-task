package app

import (
	"net/http"
)

type route struct {
	Path    string
	Methods []string
	Handler http.HandlerFunc
}
