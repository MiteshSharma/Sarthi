package api

import (
	"github.com/julienschmidt/httprouter"
)

// Init all routes here
func InitApi() *httprouter.Router {
	m := httprouter.New()

	return m
}
