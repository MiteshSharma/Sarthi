package api

import (
	"github.com/julienschmidt/httprouter"
	"github.com/MiteshSharma/Sarthi/utils"
	"github.com/urfave/negroni"
)

type Server struct  {
	Router *httprouter.Router
}

var ServerObj *Server

func InitServer()  {
	ServerObj = &Server{}
	ServerObj.Router = InitApi()
}

func StartServer()  {
	go func() {
		negroni := negroni.Classic()
		negroni.UseHandler(ServerObj.Router)
		negroni.Run(utils.ConfigParam.ServerConfig.Port)
	}()
}

func StopServer()  {
	// Closing DB connection
}