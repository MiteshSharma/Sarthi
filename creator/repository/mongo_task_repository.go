package repository

import (
	"github.com/MiteshSharma/Sarthi/database"
	"github.com/MiteshSharma/Sarthi/dao"
)

type MongoTaskRepository struct {
	dbManager database.DatabaseManager
}

func NewMongoTaskRepository(manager database.DatabaseManager) TaskRepository  {
	taskRepository:= MongoTaskRepository{manager}

	return taskRepository
}

func (t MongoTaskRepository) Get() []dao.Task {
	return nil
}