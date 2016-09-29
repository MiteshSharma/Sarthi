package main

import (
	"flag"
	"../creator/utils/"
	"../creator/api/"
	"os/signal"
	"os"
	"syscall"
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
