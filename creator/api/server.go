package api

import (
	"net/http"
	"github.com/MiteshSharma/Sarthi/creator/repository"
	"github.com/julienschmidt/httprouter"
)

type Server struct  {
	Repository repository.Repository
	Router *httprouter.Router
}

var ServerObj *Server

func InitServer()  {
	ServerObj = &Server{}

	// Creating DB connection and keeping it
	//ServerObj.Repository = Get repository object

	ServerObj.Router = InitApi()
}

func StartServer()  {
	go func() {
		error := http.ListenAndServe(":9001", ServerObj.Router)
		if error != nil {
			panic("Server start failed.")
		}
	}()
}

func StopServer()  {
	// Closing DB connection
}