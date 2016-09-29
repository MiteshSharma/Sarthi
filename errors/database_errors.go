package errors

import "fmt"

type UnknownDatabaseError struct {
	Type string
}

func (e *UnknownDatabaseError) Error() string {
	return fmt.Sprintf("Database type %s is not supported.", e.Type)
}

type DatabaseInitializationError struct {
	Message string
}

func (e *DatabaseInitializationError) Error() string {
	return fmt.Sprintf("Initialization error in database. Reason : %s", e.Message)
}
