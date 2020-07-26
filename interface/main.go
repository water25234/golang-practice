package main

import (
	"fmt"
)

// https://mgleon08.github.io/blog/2018/05/12/golang-interfaces/
type People interface {
	Speak(string) string
	Move(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func (stu *Student) Move(feet string) (foot string) {
	if feet == "right" {
		foot = "ok! right foot"
	} else {
		foot = "ok! left foot"
	}
	return
}

func main() {
	var peo People = &Student{}
	think := "bitch"
	feet := "right"
	fmt.Println(peo.Speak(think))
	fmt.Println(peo.Move(feet))
}
