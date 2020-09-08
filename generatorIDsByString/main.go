package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// execStringByID()

	// execStringByID_fnv32()

	// execStringByID_fnv64()

	fmt.Println(fnv64("water25234@gmail.com"))
}

func execStringByID() {
	arr := make(map[int64]string)

	for i := 0; i < 100; i++ {
		str := String(20)
		CheckIDs := GetCheckIDs(str)
		if _, ok := arr[CheckIDs]; ok {
			if arr[CheckIDs] != str {
				fmt.Println(CheckIDs, arr[CheckIDs], str)
			}
		}
		arr[CheckIDs] = str
		fmt.Println(CheckIDs, arr[CheckIDs], str)
	}
}

func execStringByID_fnv32() {
	arr := make(map[uint32]string)
	for i := 0; i < 100000; i++ {
		str := String(20)
		CheckIDs := fnv32(str)
		if _, ok := arr[CheckIDs]; ok {
			if arr[CheckIDs] != str {
				fmt.Println(CheckIDs, arr[CheckIDs], str)
			}
		}
		arr[CheckIDs] = str
	}
}

func execStringByID_fnv64() {
	arr := make(map[uint64]string)
	for i := 0; i < 1000000; i++ {
		str := String(20)
		CheckIDs := fnv64(str)
		if _, ok := arr[CheckIDs]; ok {
			if arr[CheckIDs] != str {
				fmt.Println(CheckIDs, arr[CheckIDs], str)
			}
		}
		arr[CheckIDs] = str
	}
}

func GetCheckIDs(text string) int64 {
	var sum int64

	for i := 0; i < len(text); i++ {

		sum = ((16 * sum) ^ int64(text[i]))

		overflow := sum / 4294967296

		sum = sum - overflow*4294967296

		sum = sum ^ overflow
	}

	fmt.Println(sum)

	if sum > 2147483647 {
		sum = sum - 4294967296
	} else if sum >= 32768 && sum <= 65535 {
		sum = sum - 65536
	} else if sum >= 128 && sum <= 255 {
		sum = sum - 256
	}

	fmt.Println(sum)

	return WithBranch(sum)
}

func WithBranch(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

// http://www.isthe.com/chongo/tech/comp/fnv/
func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

func fnv64(key string) uint64 {
	hash := uint64(14695981039346656037)
	const prime32 = uint64(1099511628211)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint64(key[i])
	}
	return hash
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}
