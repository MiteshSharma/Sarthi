package repository

import "github.com/MiteshSharma/Sarthi/dao"

type Repository interface  {
	Task() TaskRepository
}

type TaskRepository interface {
	Get() []dao.Task
}