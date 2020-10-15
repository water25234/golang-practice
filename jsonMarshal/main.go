package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	GetGames()
}

var (
	// ErrRedisKey means key is not exist
	ErrRedisKey = fmt.Errorf("redis key is not exist")
	// ErrRedisValue means value is not exist
	ErrRedisValue = fmt.Errorf("redis value is not exist")
	// ErrRedisTime means time is not exist
	ErrRedisTime = fmt.Errorf("redis time is not exist")
)

func GetGames() {
	type GameDetails struct {
		Name            string `json:"nm"`
		Id              int    `json:"-"`
		ThumbUrl        string `json:"tmb_ul"`
		HardDiskSpace   string `json:"Hard disk space"`
		Memory          string `json:"Memory"`
		OperatingSystem string `json:"Operating system"`
	}

	type UserDetails struct {
		ID   string `json:""`
		Name string `json:"name"`
	}

	games := &GameDetails{
		Name:     "Stronghold Crusader",
		Id:       1120,
		ThumbUrl: "picture/eeee.jpg",
	}

	user := &UserDetails{
		ID:   "12345",
		Name: "Justin",
	}

	var jsonCombin map[string]interface{}

	jsonCombin = make(map[string]interface{})
	jsonCombin["Game"] = *games
	jsonCombin["User"] = *user

	Str, err := json.Marshal(jsonCombin)
	if err != nil {
		fmt.Println("bbb")
		return
	}
	userValue := string(Str)
	fmt.Println(userValue)
}
