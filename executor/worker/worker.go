package worker

import (
	"time"
	"fmt"
	"github.com/MiteshSharma/Sarthi/executor/logs"
	"github.com/MiteshSharma/Sarthi/executor/work"
)

type Worker struct  {
	Id string
	Work chan work.Work
	WorkerQueue chan Worker
	Quit	chan bool
}

func NewWorker(id string, taskWorkerQueue chan Worker) *Worker  {
	worker := &Worker{
		Id: id,
		Work: make(chan work.Work),
		WorkerQueue: taskWorkerQueue,
		Quit: make(chan bool)}
	return worker
}

func (w *Worker) Start()  {
	go func() {
		for {
			// Adding worker in worker queue
			w.WorkerQueue <- *w
			select {
			case task := <- w.Work:
				logs.Logger.Debug(fmt.Sprint("Worker received work to execute with id ", task.GetId()))
				task.Execute()
				time.Sleep(1 * time.Second)
			case <- w.Quit:
				// Stop this worker
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.Quit <- true
	}()
}