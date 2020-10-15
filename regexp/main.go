package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	phone := "0929291786"
	result := regexp.MustCompile(`^[0-9]+$`).MatchString(phone)
	fmt.Println(result)

	fmt.Println("------")
	email := "justin.huang@gmail.com"
	result = regexp.MustCompile(`gmail.com$`).MatchString(email)
	fmt.Println(result)

	fmt.Println("------")

	x := "justin.huang@gmail.com"

	i := strings.Index(x, "@")
	fmt.Println("Index: ", i)
	if i > -1 {
		chars := x[:i]
		arefun := x[i+1:]
		fmt.Println(chars)
		fmt.Println(arefun)
	} else {
		fmt.Println("Index not found")
		fmt.Println(x)
	}
}
