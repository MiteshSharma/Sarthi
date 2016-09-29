package repository

import "github.com/MiteshSharma/Sarthi/database"

type MongoRepository struct  {
	databaseManager database.DatabaseManager
	task TaskRepository
}

func NewMongoRepository() Repository  {
	mongoRepository := &MongoRepository{}
	mongoRepository.databaseManager = database.GetDatabaseManager()
	mongoRepository.task = NewMongoTaskRepository(mongoRepository.databaseManager)
	return mongoRepository
}

func (m MongoRepository) Task() TaskRepository  {
	return m.task
}