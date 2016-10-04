package errors

import "fmt"

type MqNotConfiguredError struct{}

func (e *MqNotConfiguredError) Error() string {
	return "Message queue configuration not found!"
}

type MqTypeNotSupportedError struct {
	Type string
}

func (e *MqTypeNotSupportedError) Error() string {
	return fmt.Sprintf("Message queue type (%s) not supported.", e.Type)
}
