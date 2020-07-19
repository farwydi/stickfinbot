package domain

type Endpoint interface {
	Run() error
	Stop()
}
