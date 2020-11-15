package main

import "fmt"

// https://mgleon08.github.io/blog/2018/05/12/golang-interfaces/
type People interface {
	Speak(string) string
	Move(string) string
}

type Student struct {
	Name string
	Age  int32
}

type Teacher struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = stu.Name + " You are a good boy"
	} else {
		talk = stu.Name + "hi"
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

// New mean
func New(name string, age int32) People {
	return &Student{
		Name: name,
		Age:  age,
	}
}

// func NewVersion(name string) People {
// 	return &Teacher{}
// }

func main() {
	// var peo People = &Student{}
	// think := "bitch"
	// feet := "right"
	// fmt.Println(peo.Speak(think))
	// fmt.Println(peo.Move(feet))

	peo := New("Justin", 28)

	think := "bitch"
	feet := "right"
	fmt.Println(peo.Speak(think))
	fmt.Println(peo.Move(feet))

}
