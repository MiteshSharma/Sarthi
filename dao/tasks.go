package dao

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"github.com/MiteshSharma/Sarthi/database"
	"errors"
)

// task states
const TASK_STATE_SCHEDULED = "SCHEDULED"
const TASK_STATE_WAITING = "WAITING"
const TASK_STATE_EXECUTING = "EXECUTING"
const TASK_STATE_COMPLETED = "COMPLETED"

const TASKS_TYPE = "tasks"

func GetPendingTasks(timestamp int64) []Task {
	fmt.Println(fmt.Sprintf("Getting pending tasks at %d.", timestamp))
	db := database.GetDatabaseManager()

	fmt.Println(fmt.Sprintf("Getting tasks -> timestamp %v | state %v ", timestamp, TASK_STATE_SCHEDULED))
	query := &bson.M{
		"state":        TASK_STATE_SCHEDULED,
		"scheduled_at": &bson.M{"$lte": timestamp},
	}

	result := []Task{}
	err := db.GetAllByQuery(TASKS_TYPE, query, &result)

	if err != nil {
		panic(err)
	}

	return result
}

func CreateTask(task *Task) error {
	db := database.GetDatabaseManager()
	if err:= db.Create(TASKS_TYPE, task); err != nil {
		return errors.New("Error occured during task creation.")
	}
	return nil
}

func GetTask(taskId string) (Task, error) {
	task := Task{}
	db := database.GetDatabaseManager()
	if err:= db.Get(TASKS_TYPE, taskId, &task); err != nil {
		return task, errors.New("Error occured during task creation.")
	}
	return task, nil
}