package reader

import (
	"time"
	"github.com/MiteshSharma/Sarthi/dao"
	"github.com/MiteshSharma/Sarthi/executor/logs"
	"github.com/MiteshSharma/Sarthi/executor/work"
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
				// Fetch task from queue
				var work work.Work
				work = &dao.Task{Id: "123", CallbackUrl: "https://www.google.com", CallbackMethod: "GET"}
				TaskQueue <- work
			case taskResponse:= <-r.response:
				logs.Logger.Debug("Response of task received with id : "+taskResponse.Work.GetId())
			}
		}
	}()
}

func (r *Reader) Stop()  {
	go func() {
		r.Quit <- true
	}()
}