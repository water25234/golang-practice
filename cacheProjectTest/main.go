// Question-2.go
// Design and implement a data structure for cache.
// • get(key) - Get the value of the key if the key exists in the cache, otherwise return -1
// • put(key, value, weight) - Set or insert the value, when the cache reaches its capacity, it should invalidate the
// least scored key. The score is calculated as:
// when current_time != last_access_time: weight / ln(current_time - last_access_time)
// else: weight / -100
// Your data structure should be optimized for the computational complexity of get(key) i.e. Average case for
// computational complexity of get(key) could be O(1).
// In your (pseudo-)code, you can assume common data structure such as array, different type of list, hash table are
// available.
// Please explain the computational complexity of get(key) and put(...) in Big-O notation.

// Answers : get(key) = O(1), put(...) = O(1)

package main

import (
	"fmt"
	"strconv"
	"time"
)

var (
	capacity         int
	capacityLimit    int
	cacheList        map[string]map[string]interface{}
	leastScoredKey   string
	leastScoredScore string
	leastScoredValue int
	smallScoredKey   string
	smallScoredScore string
)

func get(key string) int {
	if _, ok := cacheList[key]; ok {
		value := cacheList[key]["value"]
		return value.(int)
	}
	return -1
}

func put(key string, value int, weight int) {

	if capacityLimit < capacity {

		calculateScore(key, value, weight)

	} else {
		calculateScoreCapacityLimit(key, value, weight)
		delete(cacheList, leastScoredKey)
		capacityLimit--
	}

	capacityLimit++
	score := cacheList[key]["score"].(string)
	setLeastScoredParameter(key, value, score)

	leastScoredKey = key
	leastScoredScore = score
	leastScoredValue = value
	if smallScoredKey == "" {
		smallScoredKey = key
		smallScoredScore = score
	}
}

func calculateScore(key string, value int, weight int) {

	newScore := weight / 100
	currentTime := time.Now().Unix()
	score := strconv.Itoa(newScore)
	tempInnerCacheList := make(map[string]interface{})

	tempInnerCacheList["score"] = score
	tempInnerCacheList["leastTime"] = currentTime
	tempInnerCacheList["value"] = value
	cacheList[key] = tempInnerCacheList
}

func calculateScoreCapacityLimit(key string, value int, weight int) {
	tempInnerCacheList := make(map[string]interface{})
	leastTime := cacheList[leastScoredKey]["leastTime"].(int64)
	currentTime := time.Now().Unix()
	calTime := currentTime - leastTime
	var score string
	if calTime != 0 {
		score = strconv.FormatInt(int64(weight)/(currentTime-leastTime), 10)
	} else {
		score = "1"
	}

	tempInnerCacheList["score"] = score
	tempInnerCacheList["leastTime"] = currentTime
	tempInnerCacheList["value"] = value
	cacheList[key] = tempInnerCacheList
}

func setLeastScoredParameter(key string, value int, score string) {

	scoreInt, _ := strconv.Atoi(score)
	smallScoredScoreInt, _ := strconv.Atoi(smallScoredScore)

	if scoreInt < smallScoredScoreInt {
		smallScoredKey = key
		smallScoredScore = score
	}
}

func main() {

	cacheList = make(map[string]map[string]interface{})
	capacity = 5
	put("key_1", 200, 300)
	put("key_2", 500, 999)
	put("key_3", 150, 461)
	put("key_4", 300, 732)
	put("key_5", 700, 232)
	put("key_6", 850, 876)

	fmt.Println(get("key_6"))
}
