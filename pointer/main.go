package main

import "fmt"

func main() {

	user := &User{
		ID:     123456789,
		Name:   "Justin",
		Hight:  177,
		Weight: 67,
	}
	UserInfoPointer(user)
	fmt.Println("----------")
	UserInfo(*user)
}

type User struct {
	ID     uint32
	Name   string
	Hight  uint32
	Weight uint32
}

func UserInfoPointer(userInfo *User) {
	fmt.Println(userInfo)
	userInfo.Weight = 72
	fmt.Println(userInfo)
	// UserBind(userInfo)
}

func UserInfo(userInfo User) {
	fmt.Println(userInfo)
	userInfo.Weight = 72
	fmt.Println(userInfo)
	// UserBind(userInfo)
}

func UserBind(args interface{}) {
	fmt.Println(args)
}
