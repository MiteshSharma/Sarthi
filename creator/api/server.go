package api

import (
	"github.com/MiteshSharma/Sarthi/creator/repository"
	"github.com/julienschmidt/httprouter"
	"github.com/MiteshSharma/Sarthi/creator/utils"
	"github.com/urfave/negroni"
)

type Server struct  {
	Repository repository.Repository
	Router *httprouter.Router
}

var ServerObj *Server

func InitServer()  {
	ServerObj = &Server{}
	ServerObj.Repository = repository.NewMongoRepository()
	ServerObj.Router = InitApi()
}

func StartServer()  {
	go func() {
		negroni := negroni.Classic()
		negroni.UseHandler(ServerObj.Router)
		negroni.Run(utils.Config.ServerConfig.Port)
	}()
}

func StopServer()  {
	// Closing DB connection
}