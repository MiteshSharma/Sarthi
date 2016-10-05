package main

import (
	"time"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"github.com/MiteshSharma/Sarthi/utils"
	"github.com/MiteshSharma/Sarthi/executor/reader"
	"github.com/MiteshSharma/Sarthi/executor/dispatcher"
	"github.com/MiteshSharma/Sarthi/executor/logs"
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

	logs.Logger = logs.NewExecutorLogger("Executor")

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

var communicationChan chan reader.WorkResponse

func (e *Executor) Start() {
	logs.Logger.Debug("Executor starting.")
	communicationChan = make(chan reader.WorkResponse)
	// Start task reader
	Reader = reader.NewReader(e.pingInterval, communicationChan)
	Reader.Start()
	logs.Logger.Debug("Started reader during executor start.")

	// Start dispatcher, dispatcher will start needed workers
	Dispatcher = dispatcher.NewDispatcher(e.numWorker, communicationChan)
	Dispatcher.Start()
	logs.Logger.Debug("Started dispatcher during executor start.")
	logs.Logger.Debug("Executor started.")
}

func (e *Executor) Stop() {
	logs.Logger.Debug("Executor stopping.")
	// Stop task reader
	Reader.Stop()

	// Stop dispatcher
	Dispatcher.Stop()
	logs.Logger.Debug("Executor stopped.")
}