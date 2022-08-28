package demo

import (
	"fmt"

	"github.com/gohade/hade/framework"
)

type DemoService struct {
	Service
	c framework.Continer
}

func NewDemoService(params ...interface{}) (interface{}, error) {
	//c := params[0].(framework.Continer)
	fmt.Println("new demo service")
	return &DemoService{}, nil
}

func (s *DemoService) GetFoo() Foo {
	return Foo {
		Name :"i am foo",
	}
}

