package ch5

import "testing"

type Programmer interface {
	WeiterHello() string
}

type GoProgramer struct {}

func (g *GoProgramer) WeiterHello() string {
	return "hello"
}
func TestGo(t *testing.T) {
	var p Programmer
	p = new(GoProgramer)
	t.Log(p.WeiterHello())
}