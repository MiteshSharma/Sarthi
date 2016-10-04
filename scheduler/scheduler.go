package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/MiteshSharma/Sarthi/dao"
	"github.com/MiteshSharma/Sarthi/mq"
	"github.com/MiteshSharma/Sarthi/utils"
)

type Scheduler struct {
	ping_interval time.Duration
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
	scheduler := Scheduler{
		ping_interval: (time.Duration(ping_interval) * time.Second),
	}

	scheduler.Start()
}

func (s *Scheduler) Start() {
	// start up
	fmt.Println(fmt.Sprintf("Starting scheduler with ping interval : %v", s.ping_interval))

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
	fmt.Println(fmt.Sprintf("Waiting for %v.", duration))
	<-time.After(duration)
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
	pending, err := dao.GetPendingTasks(time.Now().Unix())
	if err != nil {
		fmt.Println("Error fetching pending tasks.")
		panic(err)
	}
	fmt.Println(fmt.Sprintf("Got %d pending tasks.", len(pending)))

	mq_agent := mq.GetAgent()
	for _, task := range pending {
		fmt.Println(fmt.Sprintf("task -> %+v", task))

		// push for execution
		msg, err := json.Marshal(task)
		if err != nil {
			fmt.Println("Error marshalling task -> ", err)
			continue
		}
		mq_agent.Write(msg)

		// update task state
		dao.UpdateTaskState(&task, dao.TASK_STATE_WAITING)
	}
}
