package dispatcher

import (
	"github.com/satori/go.uuid"
	"github.com/MiteshSharma/Sarthi/executor/worker"
	"github.com/MiteshSharma/Sarthi/executor/reader"
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
		for {
			select {
			case work:= <- reader.TaskQueue:
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
