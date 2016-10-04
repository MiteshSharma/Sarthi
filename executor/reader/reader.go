package reader

import (
	"github.com/MiteshSharma/Sarthi/dao"
	"time"
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

var TaskQueue = make(chan dao.Task, 100)

func (r *Reader) Start()  {
	go func() {
		for {
			// Fetch task from queue
			task := &dao.Task{Id: "Wew"}
			TaskQueue <- *task
			time.Sleep(1000*time.Millisecond)
		}
	}()
}

func (r *Reader) Stop()  {
	go func() {
		r.Quit <- true
	}()
}