package dispatcher

import (
	"github.com/satori/go.uuid"
	"github.com/MiteshSharma/Sarthi/executor/worker"
	"github.com/MiteshSharma/Sarthi/executor/reader"
	"github.com/MiteshSharma/Sarthi/executor/logs"
	"fmt"
	"github.com/MiteshSharma/Sarthi/executor/work"
)

type Dispatcher struct  {
	NumWorker int
	Quit	chan bool
}

func NewDispatcher(numWorker int) *Dispatcher  {
	dispatcher := &Dispatcher{
		NumWorker: numWorker,
		Quit: make(chan bool)}
	return dispatcher
}

var TaskWorkerQueue chan worker.Worker

func (d *Dispatcher) Start()  {
	TaskWorkerQueue = make(chan worker.Worker, d.NumWorker)

	for count:= 0; count < d.NumWorker; count++ {
		worker := worker.NewWorker(uuid.NewV4().String(), TaskWorkerQueue)
		worker.Start()
	}

	go func() {
		var work work.Work
		for {
			select {
			case work = <- reader.TaskQueue:
				logs.Logger.Debug(fmt.Sprint("Work received by dispatcher to execute with id ", work.GetId()))
				go func() {
					var worker worker.Worker = <-TaskWorkerQueue
					worker.Work <- work
				}()
			case <- d.Quit:
				return
			}
		}
	}()
}

func (d *Dispatcher) Stop() {
	go func() {
		d.Quit <- true
	}()
}
