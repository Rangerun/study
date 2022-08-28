package demo 

import (
	"fmt"

	"github.com/gohade/hade/framework"
)

type DemoServiceProvider struct {
}

func (sp *DemoServiceProvider) Name() string {
	return Key
}


func (sp *DemoServiceProvider) Register(c framework.Continer) framework.NewInstance {
	return NewDemoService
}

func (sp *DemoServiceProvider) Params(c framework.Continer) []interface{} {
	return []interface{}{c}
}

func (sp *DemoServiceProvider) Boot(c framework.Continer) error {
	fmt.Println("demo service boot")
	return nil
}

func (sp *DemoServiceProvider) IsDefer() bool {
	return true
}
