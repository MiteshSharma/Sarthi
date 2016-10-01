package database

import (
	"fmt"

	"github.com/MiteshSharma/Sarthi/utils"
)

type DatabaseManager interface {
	Create(docType string, doc interface{}) error
	Get(docType string, id string, result interface{}) error
	GetByQuery(docType string, query interface{}, result interface{}) error
	GetAllByQuery(docType string, query interface{}, result interface{}) error
	Save(docType string, docId string, doc interface{}) error
	Delete(docType string, docId string) error
	DeleteByQuery(docType string, query interface{}) error
	// GetQueryResult(docType string, query interface{}, start int, size int, result *QueryResult) ([]bson.Raw, error)
}

// TODO externalize this or move to common config file
const FETCH_LIMIT = 1000

var dbManager DatabaseManager

func GetDatabaseManager() DatabaseManager {
	if dbManager != nil {
		return dbManager
	} else {
		config := utils.ConfigParam.DatabaseConfig
		dbManager, err := NewMongodbManager(map[string]string{
			"host":    config.Host,
			"db_name": config.DbName,
		})
		if err != nil {
			fmt.Println("Error initializing database manager.")
			panic(err)
		}
		return dbManager
	}
}
