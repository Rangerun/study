package main

import (
	"fmt"
	"proto/user"
	"google.golang.org/protobuf/proto"
)

func main() {
	article := &user.Person{
		Aid : 1,
		Title: "protobuf for golang",
		Views: 100,
	}

	bytes, _ := proto.Marshal(article)
	fmt.Printf("bytes: %v\n", bytes)

	a := &user.Person{}
	proto.Unmarshal(bytes, a)
	fmt.Println(a.GetAid())
	fmt.Println(a.GetTitle())
	fmt.Println(a.Aid, a.Title)
}