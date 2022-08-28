package framework

type NewInstance func(...interface{}) (interface{}, error)

type ServiceProvider interface {
	Register(Continer) NewInstance
	Boot(Continer) error
	IsDefer() bool
	Params(Continer) []interface{}
	Name() string
}