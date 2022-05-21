package ch6

import (
	"errors"
	"testing"

)


func TestGo(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()
	t.Log(00000)
	panic(errors.New("err"))
}