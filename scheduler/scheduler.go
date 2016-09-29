package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/MiteshSharma/Sarthi/dao"
)

type Scheduler struct {
	ping_interval time.Duration
}

func main() {
	// start the scheduler
	scheduler := Scheduler{
		ping_interval: 2, // TODO externalize this setting, possibly read from command line param
	}
	scheduler.Start()
}

func (s *Scheduler) Start() {
	// start up
	fmt.Println(fmt.Sprintf("Starting scheduler with ping interval : %d seconds", s.ping_interval))

	// channels and waitgroup to ensure clean exit
	var wg sync.WaitGroup
	interrupted := false
	interruptChannel := make(chan bool)
	workChannel := make(chan bool)

	// handle SIGINT
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			// sig is a ^C, handle it
			interruptChannel <- true
		}
	}()

	go sleepAndWork(s.ping_interval, workChannel)

	// work, sleep, repeat by listening to workChannel
	// interrupt and exit when interruptChannel receives message
	for !interrupted {
		select {
		case <-interruptChannel:
			interrupted = true
			break
		case <-workChannel:
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.work()
			}()
			go sleepAndWork(s.ping_interval, workChannel)
		}
	}

	fmt.Println("Waiting for processes to complete.")
	wg.Wait()
	fmt.Println("Scheduler stopped.")
}

func sleepAndWork(duration time.Duration, ch chan<- bool) {
	fmt.Println(fmt.Sprintf("Waiting for %d seconds.", duration))
	<-time.After(duration * time.Second)
	ch <- true
}

func (s *Scheduler) work() {
	// defer recovery function
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering from error -> ", r)
		}
	}()

	// get pending tasks
	pending := dao.GetPendingTasks(time.Now().Unix())
	fmt.Println(fmt.Sprintf("Got %d pending tasks.", len(pending)))
	for _, task := range pending {
		fmt.Println(fmt.Sprintf("task -> %+v", task))
		// TODO
		// push for execution
		// update task state
	}
}
