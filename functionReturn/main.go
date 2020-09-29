package main

import "fmt"

func main() {

	u := &User{}
	name := "Justin"

	result := u.UserInfo(name)
	fmt.Println(result)

	resultV2 := UserInfoV2(name)
	fmt.Println(resultV2)

	resultV3 := UserInfoV3(name)
	fmt.Println(resultV3)
}

type User struct {
	ID     string
	Name   string
	Hight  uint16
	Weight uint16
}

func (user *User) UserInfo(name string) (info *User) {
	info = &User{} // response is pointer, it must be declare new struct pointer.
	info.Name = name
	return
}

func UserInfoV2(name string) (info *User) {
	info = &User{} // response is pointer, it must be declare new struct pointer.
	info.Name = name
	return
}

func UserInfoV3(name string) (info User) {
	info.Name = name // it's work
	return
}

func addMult(a, b int) (add, mul int) {
	add = a + b
	mul = a * b

	return
}
