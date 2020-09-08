package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	execStringByID()

	// fmt.Println(GetCheckIDs("8TtAq4imIgLOaJGQZEbpmbEtTRLG14"))
	// fmt.Println(GetCheckIDs("NLZR2KgcRyqfxj8n9Zj77nl8uGlrLN"))
	// fmt.Println(GetCheckIDs("water25234@gmail.com"))
	// fmt.Println(GetCheckIDs("justin.huang@kkday.com"))
	// fmt.Println(GetCheckIDs("justin.huang@newtype.games"))
	// fmt.Println(GetCheckIDs("wei.shun.huang771210@hotmail.games"))
	// fmt.Println(GetCheckIDs("wei.shun.huang771210@hotmail.gamesjustin.huang@newtype.games"))
	// fmt.Println(GetCheckIDs("A"))
	// fmt.Println(GetCheckIDs("AB"))
	// fmt.Println(GetCheckIDs("ABC"))
	// fmt.Println(GetCheckIDs("ABCD"))
	// fmt.Println(GetCheckIDs("ABCDE"))
	// fmt.Println(GetCheckIDs("test@qq.com"))
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

var (
	Field_mapping = map[string]int64{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"A": 10,
		"B": 11,
		"C": 12,
		"D": 13,
		"E": 14,
		"F": 15,
		"G": 16,
		"H": 17,
		"I": 18,
		"J": 19,
		"K": 20,
		"L": 21,
		"M": 22,
		"N": 23,
		"O": 24,
		"P": 25,
		"Q": 26,
		"R": 27,
		"S": 28,
		"T": 29,
		"U": 30,
		"V": 31,
		"X": 32,
		"Y": 33,
		"Z": 34,
		"a": 35,
		"b": 36,
		"c": 37,
		"d": 38,
		"e": 39,
		"f": 40,
		"g": 41,
		"h": 42,
		"i": 43,
		"j": 44,
		"k": 45,
		"l": 46,
		"m": 47,
		"n": 48,
		"o": 49,
		"p": 50,
		"q": 51,
		"r": 52,
		"s": 53,
		"t": 54,
		"u": 55,
		"v": 56,
		"w": 57,
		"x": 58,
		"y": 59,
		"z": 60,
	}
)

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

func Arr() []string {
	arr := []string{
		"",
		"X",
		"XX",
		"XXX",
		"XXXX",
		"XXXXX",
		"XXXXXX",
		"XXXXXXX",
		"XXXXXXXX",
		"XXXXXXXXX",
		"XXXXXXXXXX",
		"XXXXXXXXXXX",
		"XXXXXXXXXXXX",
		"XXXXXXXXXXXXX",
		"XXXXXXXXXXXXXX",
		"XXXXXXXXXXXXXXX",
		"XXXXXXXXXXXXXXXX", // 0
		"XXXXXXXXXXXXXXXXX",
		"XXXXXXXXXXXXXXXXXX",
		"XXXXXXXXXXXXXXXXXXX",
		"justin.huang@kkday.com",
		"justin.huang@NewType.games",
		"water25234@gmail.com",
		"water25234@hotmail.com",
		"abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijk.com",
		"",
		"A",
		"AA",
		"AAA",
		"AAAA",
		"AAAAA",
		"AAAAAA",
		"AAAAAAA",
		"AAAAAAAA",
		"AAAAAAAAA",
		"AAAAAAAAAA",
		"AAAAAAAAAAA",
		"AAAAAAAAAAAA",
		"AAAAAAAAAAAAA",
		"AAAAAAAAAAAAAA",
		"AAAAAAAAAAAAAAA",
		"AAAAAAAAAAAAAAAA",
		"AAAAAAAAAAAAAAAAA",
	}
	return arr
}
