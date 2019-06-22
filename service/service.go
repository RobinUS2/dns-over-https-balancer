package service

type Instance struct {
}

func (instance Instance) Run() error {
	return nil
}

func New() *Instance {
	return &Instance{}
}
