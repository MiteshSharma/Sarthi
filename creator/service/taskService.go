package service

import (
	"github.com/satori/go.uuid"
	"github.com/MiteshSharma/Sarthi/dao"
)

func CreateTask(task *dao.Task) error {
	if task.Id == "" {
		task.Id = uuid.NewV4().String()
	}
	if err := dao.CreateTask(task); err != nil {
		return err
	}
	return nil
}

func GetTask(taskId string) (dao.Task, error) {
	return dao.GetTask(taskId)
}