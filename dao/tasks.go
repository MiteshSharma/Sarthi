package dao

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"github.com/MiteshSharma/Sarthi/database"
	errorz "github.com/MiteshSharma/Sarthi/errors"
)

// task states
const TASK_STATE_SCHEDULED = "SCHEDULED"
const TASK_STATE_WAITING = "WAITING"
const TASK_STATE_EXECUTING = "EXECUTING"
const TASK_STATE_COMPLETED = "COMPLETED"

const TASKS_TYPE = "tasks"

func GetPendingTasks(timestamp int64) ([]Task, error) {
	fmt.Println(fmt.Sprintf("Getting pending tasks at %d.", timestamp))
	db := database.GetDatabaseManager()

	fmt.Println(fmt.Sprintf("Getting tasks -> timestamp %v | state %v ", timestamp, TASK_STATE_SCHEDULED))
	query := &bson.M{
		"state":        TASK_STATE_SCHEDULED,
		"scheduled_at": &bson.M{"$lte": timestamp},
	}

	result := []Task{}
	if err := db.GetAllByQuery(TASKS_TYPE, query, &result); err != nil {
		return result, err
	}

	return result, nil
}

func CreateTask(task *Task) error {
	db := database.GetDatabaseManager()
	if err := db.Create(TASKS_TYPE, task); err != nil {
		return err
	}

	return nil
}

func GetTask(taskId string) (Task, error) {
	task := Task{}
	db := database.GetDatabaseManager()
	if err := db.Get(TASKS_TYPE, taskId, &task); err != nil {
		return task, err
	}

	return task, nil
}

func UpdateTaskState(task *Task, state string) error {
	switch state {
	// break for known cases
	case TASK_STATE_WAITING:
	case TASK_STATE_SCHEDULED:
	case TASK_STATE_EXECUTING:
	case TASK_STATE_COMPLETED:
	default:
		return &errorz.InvalidTaskStateError{
			State: state,
		}
	}
	task.State = state

	db := database.GetDatabaseManager()
	if err := db.Save(TASKS_TYPE, task.Id, task); err != nil {
		return err
	}

	return nil
}
