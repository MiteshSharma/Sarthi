package work

type Work interface {
	GetId() string
	Execute()
}