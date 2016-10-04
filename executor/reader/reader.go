package reader

import (
	"github.com/MiteshSharma/Sarthi/dao"
	"time"
	"github.com/MiteshSharma/Sarthi/executor/logs"
	"github.com/MiteshSharma/Sarthi/executor/work"
)

type Reader struct  {
	pingInterval time.Duration
	Quit chan bool
}

func NewReader(pingInterval time.Duration) *Reader  {
	reader := &Reader{
		pingInterval: pingInterval,
		Quit: make(chan bool)}
	return reader
}

var TaskQueue = make(chan work.Work, 100)

func (r *Reader) Start()  {
	go func() {
		for {
			logs.Logger.Debug("Reading data from message queue.")
			time.Sleep(1000*time.Millisecond)
			// Fetch task from queue
			var work work.Work
			work = &dao.Task{Id: "123", CallbackUrl: "https://www.google.com", CallbackMethod: "GET"}
			TaskQueue <- work
			time.Sleep(1000*time.Millisecond)
		}
	}()
}

func (r *Reader) Stop()  {
	go func() {
		r.Quit <- true
	}()
}