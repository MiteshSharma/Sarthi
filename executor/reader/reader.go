package reader

import (
	"time"
	"github.com/MiteshSharma/Sarthi/executor/logs"
	"github.com/MiteshSharma/Sarthi/executor/work"
	"github.com/MiteshSharma/Sarthi/executor/source"
)

type Reader struct  {
	pingInterval time.Duration
	response chan WorkResponse
	Quit chan bool
}

type WorkResponse struct  {
	Work work.Work
	IsSuccess bool
}

func NewReader(pingInterval time.Duration, response chan WorkResponse) *Reader  {
	reader := &Reader{
		pingInterval: pingInterval,
		response: response,
		Quit: make(chan bool)}
	return reader
}

var TaskQueue = make(chan work.Work, 100)

func (r *Reader) Start()  {
	go func() {
		for {
			select {
			case <-time.After(time.Second * 1):
				logs.Logger.Debug("Reading data from message queue.")
				source := source.GetSource()
				// Fetch task from queue
				var work work.Work
				work = source.Get()
				if (work.GetId() != "") {
					TaskQueue <- work
				} else {
					logs.Logger.Debug("No message found to execute in reader.")
				}
				//work = &dao.Task{Id: "123", CallbackUrl: "https://www.google.com", CallbackMethod: "GET"}
			case taskResponse:= <-r.response:
				logs.Logger.Debug("Response of task received with id : "+taskResponse.Work.GetId())
				source.GetSource().Delete(taskResponse.Work.GetId())
			}
		}
	}()
}

func (r *Reader) Stop()  {
	go func() {
		r.Quit <- true
	}()
}