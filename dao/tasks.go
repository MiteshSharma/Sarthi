package dao

import (
	"fmt"
	"time"
)

// task states
const TASK_STATE_SCHEDULED = "SCHEDULED"
const TASK_STATE_WAITING = "WAITING"
const TASK_STATE_EXECUTING = "EXECUTING"
const TASK_STATE_COMPLETED = "COMPLETED"

// task
type Task struct {
	id               string
	state            string
	callback_url     string
	callback_method  string
	callback_payload string
	schedule         string
	scheduled_at     int64
}

func GetPendingTasks(timestamp int64) []Task {
	fmt.Println(fmt.Sprintf("Getting pending tasks at %d.", timestamp))
	// TODO get tasks from the db where scheduled_at < timestamp and state == SCHEDULED
	tasks := []Task{
		Task{
			id:              "id1",
			callback_url:    "http://google.com",
			callback_method: "GET",
			schedule:        "1234567890",
			scheduled_at:    time.Now().Unix(),
			state:           "SCHEDULED",
		},
		Task{
			id:              "id2",
			callback_url:    "http://google.com",
			callback_method: "GET",
			schedule:        "1234567890",
			scheduled_at:    time.Now().Unix(),
			state:           "SCHEDULED",
		},
	}
	return tasks
}
