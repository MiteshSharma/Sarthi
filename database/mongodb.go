package database

import (
	"gopkg.in/mgo.v2"

	"github.com/MiteshSharma/Sarthi/errors"
)

type MongodbManager struct {
	host    string
	db_name string
}

func NewMongodbManager(dbConfig map[string]string) (*MongodbManager, error) {
	var host string
	var db_name string
	var ok bool
	var mdb *MongodbManager

	if host, ok = dbConfig["host"]; !ok {
		return mdb, &errors.DatabaseInitializationError{
			Message: "Missing key in mongodb manager configuration : host",
		}
	}

	if db_name, ok = dbConfig["db_name"]; !ok {
		return mdb, &errors.DatabaseInitializationError{
			Message: "Missing key in mongodb manager configuration : db_name",
		}
	}

	return &MongodbManager{
		host:    host,
		db_name: db_name,
	}, nil
}

func (m *MongodbManager) Create(docType string, doc interface{}) error {
	session, err := mgo.Dial(m.host)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB(m.db_name).C(docType)
	err = c.Insert(doc)

	lastErr, ok := err.(*mgo.LastError)
	if ok && lastErr.Code == 11000 {
		return &errors.DuplicateError{Type: docType}
	} else {
		return err
	}
}

func (m *MongodbManager) Get(docType string, id string, result interface{}) error {
	session, err := mgo.Dial(m.host)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB(m.db_name).C(docType)
	err = c.FindId(id).One(result)

	if err != nil && err.Error() == "not found" {
		return &errors.NotFoundError{Type: docType}
	} else {
		return err
	}
}

func (m *MongodbManager) GetByQuery(docType string, query interface{}, result interface{}) error {
	session, err := mgo.Dial(m.host)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB(m.db_name).C(docType)
	err = c.Find(query).One(result)

	if err != nil && err.Error() == "not found" {
		return &errors.NotFoundError{Type: docType}
	} else {
		return err
	}
}

func (m *MongodbManager) GetAllByQuery(docType string, query interface{}, result interface{}) error {
	session, err := mgo.Dial(m.host)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB(m.db_name).C(docType)
	found := c.Find(query).Limit(FETCH_LIMIT + 1)
	count, err := found.Count()

	if err == nil {
		if count > FETCH_LIMIT {
			panic(&errors.FetchLimitExceededError{})
		}
		err = found.All(result)
	}

	return err
}

func (m *MongodbManager) Save(docType string, docId string, doc interface{}) error {
	session, err := mgo.Dial(m.host)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB(m.db_name).C(docType)
	err = c.UpdateId(docId, doc)
	return err
}

func (m *MongodbManager) Delete(docType string, docId string) error {
	session, err := mgo.Dial(m.host)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB(m.db_name).C(docType)
	err = c.RemoveId(docId)
	return err
}

func (m *MongodbManager) DeleteByQuery(docType string, query interface{}) error {
	session, err := mgo.Dial(m.host)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB(m.db_name).C(docType)
	err = c.Remove(query)
	return err
}

// func (m *MongodbManager) GetQueryResult(docType string, query interface{}, start int, size int, result *QueryResult) ([]bson.Raw, error) {
// 	session, err := mgo.Dial(m.host)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer session.Close()

// 	result.Start = start
// 	result.Size = size

// 	c := session.DB(m.db_name).C(docType)
// 	found := c.Find(query)

// 	var values []bson.Raw
// 	count, err := found.Count()
// 	if err == nil {
// 		result.Total = count
// 		if size == 0 {
// 			values = []bson.Raw{}
// 		} else {
// 			err = found.Skip(start).Limit(size).All(&values)
// 		}
// 	}

// 	return values, err
// }
