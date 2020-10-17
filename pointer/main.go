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
	fmt.Println("----------")

	userVersion2 := User{
		ID:     123456789,
		Name:   "Justin",
		Hight:  177,
		Weight: 67,
	}
	UserInfo(userVersion2)

	// &{123456789 Justin 177 67}
	// &{123456789 Justin 177 72}
	// ----------
	// {123456789 Justin 177 72}
	// {123456789 Justin 177 72}
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

func example1() {
	// https://www.mdeditor.tw/pl/paja/zh-tw
	// 建立指標陣列的時候，不適合用`range`迴圈。請參考`正確程式碼`示例。
	// 這個問題是 range 迴圈的實現邏輯引起的。跟 for 迴圈不一樣的地方在於 range 迴圈中的 x 變數是臨時變數。range迴圈只是將值拷貝到 x 變數中。因此記憶體地址都是一樣的。

	slice := []int{0, 1, 2, 3}
	myMap := make(map[int]*int)

	for index, value := range slice {
		fmt.Printf("%d => %d\n", value, &value)
		myMap[index] = &value
	}

	// for i := 0; i < len(slice); i++ {
	// 	myMap[i] = &slice[i]
	// }

	fmt.Println("=====new map=====")
	for k, v := range myMap {
		// fmt.Printf("%d => %d\n", k, *v)
		fmt.Printf("指標陣列：索引:%d 值:%d 值的記憶體地址:%d\n", k, *v, v)
	}
}
