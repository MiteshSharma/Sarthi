package errors

import "fmt"

type DuplicateError struct {
	Type string
}

func (e *DuplicateError) Error() string {
	return fmt.Sprintf("An entity of type %s with given ID already exists.", e.Type)
}

type NotFoundError struct {
	Type string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Entity of type %s with requested ID not found.", e.Type)
}

type FetchLimitExceededError struct{}

func (e *FetchLimitExceededError) Error() string {
	return "Fetch limit exceeded. Please use paginated fetches."
}

type InvalidTaskStateError struct {
	State string
}

func (e *InvalidTaskStateError) Error() string {
	return fmt.Sprintf("Invalid task state %s.", e.State)
}
