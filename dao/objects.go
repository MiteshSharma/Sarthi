package dao

import (
	"errors"
	"net/http"
	"strings"
)

type Task struct {
	Id              string `json:"id" bson:"_id"`
	State           string `json:"state" bson:"state"`
	CallbackUrl     string `json:"callback_url" bson:"callback_url"`
	CallbackMethod  string `json:"callback_method" bson:"callback_method"`
	CallbackPayload string `json:"callback_payload" bson:"callback_payload"`
	Schedule        string `json:"schedule" bson:"schedule"`
	ScheduledAt     int64  `json:"scheduled_at" bson:"scheduled_at"`
}

func (t *Task) IsValid() error {
	if t.CallbackUrl == "" {
		return errors.New("Callback url not defined.")
	}
	if t.Schedule == "" && t.ScheduledAt == 0 {
		return errors.New("Schedule time is not specified.")
	}
	return nil
}

func (t *Task) Execute() {
	task, error := GetTaskLock(t.Id)
	// We found a task which means lock is acquired
	if error == nil {
		// This means task found and lock is acquired. Don't run it until lock is freed.
		return
	}

	if err:= CreateTaskLock(&task); err != nil {
		println(err.Error())
	}

	client := &http.Client{}
	request, error := http.NewRequest(t.CallbackMethod, t.CallbackUrl, strings.NewReader(t.CallbackPayload))
	if error != nil {
		println("Ok we got paniced"+ error.Error())
	}

	response, error := client.Do(request)
	if error != nil {
		println("Error received during request: "+ error.Error())
	} else {
		println("Response received: "+response.Status)
	}

	if err:= DeleteTaskLock(task.Id); err != nil {
		println(err.Error())
	}

}