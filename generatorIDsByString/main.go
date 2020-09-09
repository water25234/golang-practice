package main

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"time"

	// https://github.com/bingoohuang/blog/issues/143 compare
	"github.com/cespare/xxhash"
	"github.com/dgryski/go-farm"
	"github.com/spaolacci/murmur3"
)

func main() {

	mur64 := murmur3New64("water25234@gmail.com")
	fmt.Println("murmur3 64: ", mur64)

	farm64 := farm.Hash64([]byte("water25234@gmail.com"))
	fmt.Println("farm 64: ", farm64)

	//prometheus & lnfluxDB use this xxhash package
	xxhash := xxhash.Sum64([]byte("water25234@gmail.com"))
	fmt.Println("xxhash 64: ", xxhash)

	fnv64 := fnv64("water25234@gmail.com")
	fmt.Println("fnv 64: ", fnv64)

	IDs := GetCheckIDs("water25234@gmail.com")
	fmt.Println("IDs: ", IDs)
}

func murmur3New32(str string) uint32 {
	murmur3New32 := murmur3.New32()
	murmur3New32.Write([]byte(str))
	return murmur3New32.Sum32()
}

func murmur3New64(str string) uint64 {
	murmur3New64 := murmur3.New64()
	murmur3New64.Write([]byte(str))
	return murmur3New64.Sum64()
}

// https://golang.org/src/hash/fnv/fnv.go?s=1072:1097#L34
func golFnv32(text string) uint32 {
	a := fnv.New32()
	a.Write([]byte(text))
	return a.Sum32()
}

func golFnv64(text string) uint64 {
	a := fnv.New64()
	a.Write([]byte(text))
	return a.Sum64()
}

func execStringByID() {
	arr := make(map[int64]string)

	for i := 0; i < 100000; i++ {
		str := String(20)
		CheckIDs := GetCheckIDs(str)
		if _, ok := arr[CheckIDs]; ok {
			if arr[CheckIDs] != str {
				fmt.Println(CheckIDs, "previous: "+arr[CheckIDs], "currently: "+str)
			}
		}
		arr[CheckIDs] = str
	}
}

func execStringByID_fnv32() {
	arr := make(map[uint32]string)
	for i := 0; i < 100000; i++ {
		str := String(150)
		CheckIDs := fnv32(str)
		if _, ok := arr[CheckIDs]; ok {
			if arr[CheckIDs] != str {
				fmt.Println(CheckIDs, "previous: "+arr[CheckIDs], "currently: "+str)
			}
		}
		arr[CheckIDs] = str
	}
}

func execStringByID_fnv64() {
	arr := make(map[uint64]string)
	for i := 0; i < 10000000; i++ {
		str := String(150)
		CheckIDs := fnv64(str)
		if _, ok := arr[CheckIDs]; ok {
			if arr[CheckIDs] != str {
				fmt.Println(CheckIDs, "previous: "+arr[CheckIDs], "currently: "+str)
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

	if sum > 2147483647 {
		sum = sum - 4294967296
	} else if sum >= 32768 && sum <= 65535 {
		sum = sum - 65536
	} else if sum >= 128 && sum <= 255 {
		sum = sum - 256
	}

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

const (
	BIG_M = 0xc6a4a7935bd1e995
	BIG_R = 47
	SEED  = 0x1234ABCD
)

func MurmurHash64A(str string) int64 {
	var k int64
	var data []byte = []byte(str)
	h := SEED ^ int64(uint64(len(data))*BIG_M)

	var ubigm uint64 = BIG_M
	var ibigm = int64(ubigm)

	for l := len(data); l >= 8; l -= 8 {
		k = int64(int64(data[0]) | int64(data[1])<<8 | int64(data[2])<<16 | int64(data[3])<<24 |
			int64(data[4])<<32 | int64(data[5])<<40 | int64(data[6])<<48 | int64(data[7])<<56)
		k := k * ibigm
		k ^= int64(uint64(k) >> BIG_R)
		k = k * ibigm

		h = h ^ k
		h = h * ibigm
		data = data[8:]
	}

	switch len(data) {
	case 7:
		h ^= int64(data[6]) << 48
		fallthrough
	case 6:
		h ^= int64(data[5]) << 40
		fallthrough
	case 5:
		h ^= int64(data[4]) << 32
		fallthrough
	case 4:
		h ^= int64(data[3]) << 24
		fallthrough
	case 3:
		h ^= int64(data[2]) << 16
		fallthrough
	case 2:
		h ^= int64(data[1]) << 8
		fallthrough
	case 1:
		h ^= int64(data[0])
		h *= ibigm
	}

	return h
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
