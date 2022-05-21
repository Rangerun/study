package ch4

import "testing"

type Student struct {
	Id		int
	name	string
}

func TestGo(t *testing.T) {
	s1 := Student{}
	s2 := Student{Id: 1,name: "zhan"}
	s3 := new(Student)
	t.Logf("type is %T", s1)
	t.Log(s2)
	t.Logf("type is %T", s3)

}