package main

import (
	"flag"
	"os/signal"
	"os"
	"syscall"
	"github.com/MiteshSharma/Sarthi/creator/utils"
	"github.com/MiteshSharma/Sarthi/creator/api"
)

var configFileName string

func main() {

	parseCmdParams()

	utils.LoadConfig(configFileName)

	api.InitServer()
	api.StartServer()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c

	api.StopServer()
}

func parseCmdParams()  {
	flag.StringVar(&configFileName, "config", "config.json", "")
	flag.Parse()
}
