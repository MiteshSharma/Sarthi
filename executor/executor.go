package main

import (
	"time"
	"flag"
	"github.com/MiteshSharma/Sarthi/utils"
	"os"
	"os/signal"
	"syscall"
	"github.com/MiteshSharma/Sarthi/executor/reader"
	"github.com/MiteshSharma/Sarthi/executor/dispatcher"
)

type Executor struct {
	pingInterval time.Duration
	numWorker int
}

func main() {
	// read command line params
	var config string
	var ping_interval int
	flag.StringVar(&config, "config", "config.json", "JSON configuration file.")
	flag.IntVar(&ping_interval, "ping_interval", 10, "Pinging interval for scheduler in seconds.")
	flag.Parse()

	// load configuration file
	utils.LoadConfig(config)

	// start the scheduler
	executor := Executor{
		pingInterval: (time.Duration(ping_interval) * time.Second),
		numWorker: 1,
	}

	executor.Start()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c

	executor.Stop()
}

var Reader *reader.Reader
var Dispatcher *dispatcher.Dispatcher

func (e *Executor) Start() {
	// Start task reader
	Reader = reader.NewReader(e.pingInterval)
	Reader.Start()

	// Start dispatcher, dispatcher will start needed workers
	Dispatcher = dispatcher.NewDispatcher(e.numWorker)
	Dispatcher.Start()
}

func (e *Executor) Stop() {
	// Stop task reader
	Reader.Stop()

	// Stop dispatcher
	Dispatcher.Stop()
}