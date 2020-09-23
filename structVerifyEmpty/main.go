package main

import "fmt"

type User struct {
	id   int32
	name string
}

func main() {

	traveler := &User{}
	viewer := &User{
		id:   123,
		name: "Justin",
	}

	fmt.Println(&User{id: 123, name: "Justin"})

	fmt.Println(traveler)
	fmt.Println(*traveler)
	fmt.Println(&traveler)

	fmt.Println(traveler == &User{})
	fmt.Println(*traveler == User{})

	fmt.Println("----------")

	fmt.Println(viewer)
	fmt.Println(*viewer)
	fmt.Println(&viewer)

	fmt.Println(viewer == &User{})
	fmt.Println(*viewer == User{})
}
