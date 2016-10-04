package dao

import (
	"github.com/MiteshSharma/Sarthi/database"
)

const TASK_LOCKS_TYPE = "task_locks"

func CreateTaskLock(task *Task) error {
	db := database.GetDatabaseManager()
	if err := db.Create(TASK_LOCKS_TYPE, task); err != nil {
		return err
	}

	return nil
}

func GetTaskLock(taskId string) (Task, error) {
	task := Task{}
	db := database.GetDatabaseManager()
	if err := db.Get(TASK_LOCKS_TYPE, taskId, &task); err != nil {
		return task, err
	}

	return task, nil
}

func DeleteTaskLock(taskId string) error {
	db := database.GetDatabaseManager()
	if err := db.Delete(TASK_LOCKS_TYPE, taskId); err != nil {
		return err
	}

	return nil
}