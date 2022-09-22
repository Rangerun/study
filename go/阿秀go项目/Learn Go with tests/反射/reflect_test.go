package main

import (
	"testing"
	"reflect"
)



// 嵌套结构

type Person struct {
    Name    string
    Profile Profile
}

type Profile struct {
    Age  int
    City string
}



func TestWalk(t *testing.T) {

    cases := []struct{
        Name string
        Input interface{}
        ExpectedCalls []string
    } {
        {
			"Slices",
			[]Profile {
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
    }

    for _, test := range cases {
        t.Run(test.Name, func(t *testing.T) {
            var got []string
            walk(test.Input, func(input string) {
                got = append(got, input)
            })

            if !reflect.DeepEqual(got, test.ExpectedCalls) {
                t.Errorf("got %v, want %v", got, test.ExpectedCalls)
            }
        })
    }
}


func walk(x interface{}, fn func(input string)) {
    val := getValue(x)

    numberOfValues := 0
    var getField func(int) reflect.Value

    switch val.Kind() {
    case reflect.String:
        fn(val.String())
    case reflect.Struct:
        numberOfValues = val.NumField()
        getField = val.Field
    case reflect.Slice, reflect.Array:
        numberOfValues = val.Len()
        getField = val.Index
    case reflect.Map:
        for _, key := range val.MapKeys() {
            walk(val.MapIndex(key).Interface(), fn)
        }
    }

    for i:=0; i< numberOfValues; i++ {
        walk(getField(i).Interface(), fn)
    }
}

// 指针类型的 Value 不能使用 NumField 方法，在执行此方法前需要调用 Elem() 提取底层值。
func getValue(x interface{}) reflect.Value {
    val := reflect.ValueOf(x)

    if val.Kind() == reflect.Ptr {
        val = val.Elem()
    }

    return val
}