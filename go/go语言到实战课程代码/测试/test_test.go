package main

import (
	"testing"
	"bytes"
)
/*func Test_testDouble(t *testing.T) {
	type args struct {
		op int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			"test1",
			args{op: 1},
			2,
		},
		{
			"test2",
			args{op: 1},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testDouble(tt.args.op); got != tt.want {
				t.Errorf("testDouble() = %v, want %v", got, tt.want)
			}
		})
	}
}*/

func BenchmarkStrAdd(b *testing.B) {
	elems := []string{"1", "2", "3", "4"}
	b.ResetTimer()
	for i:=0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StopTimer()
}

func BenchmarkStrAddBuf(b *testing.B) {
	elems := []string{"1", "2", "3", "4"}
	b.ResetTimer()
	for i:=0; i < b.N; i++ {
		var buf bytes.Buffer
		for _, elem := range elems {
			buf.WriteString(elem)
		}
	}
	b.StopTimer()
}
