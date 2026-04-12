// https://pjchender.dev/golang/slice-and-array/
// https://geektutu.com/post/hpg-slice.html
package main

import "fmt"

func main() {
	example1()

	fmt.Println("----- end -----")

	example2()

	fmt.Println("----- end -----")

	example3()
}

func example1() {
	// Example 1
	names := []string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	a := names[0:2]
	b := names[1:3]

	fmt.Println(a, b) // [John Paul] [Paul George

	b[0] = "XXX"       // a 和 b 這兩個 slice 參照到的是底層相同的 array
	fmt.Println(a, b)  // [John XXX] [XXX George]
	fmt.Println(names) // [John XXX George Ringo]

	fmt.Println("----------")

	// Example 2
	words := []string{"Hello", "every", "one"}
	fmt.Println(words) // [Hello every one]
	changeSliceItem(words)
	fmt.Println(words) // [Hi every one]

	fmt.Println("----------")

	// Example 3
	nums := make([]int, 0, 8)
	nums = append(nums, 1, 2, 3, 4, 5)
	nums2 := nums[2:4]
	printLenCap(nums)  // len: 5, cap: 8 [1 2 3 4 5]
	printLenCap(nums2) // len: 2, cap: 6 [3 4]

	nums2 = append(nums2, 50, 60)
	printLenCap(nums)  // len: 5, cap: 8 [1 2 3 4 50]
	printLenCap(nums2) // len: 4, cap: 6 [3 4 50 60]

	fmt.Println("----------")
	x := []int{1, 2}
	foo(x)
	fmt.Println(x)
}

func example2() {
	a := []int{10, 20, 30, 40, 50}

	fmt.Println(a[2:3]) // 從 2 到 3 → [30]
	fmt.Println(a[2:4]) // 從 2 到 4 → [30, 40]
	fmt.Println(a[:3])  // 省略左邊 = 從 0 開始 → [10, 20, 30]
	fmt.Println(a[2:])  // 省略右邊 = 到結尾   → [30, 40, 50]
	fmt.Println(a[:])   // 全部省略 = 整個 slice → [10, 20, 30, 40, 50]
	fmt.Println(a[2:2]) // 同一條線，夾不到東西 → []（空 slice）
}

func example3() {
	a := []int{10, 20, 30, 40, 50}
	b := a[1:3] // b → [20, 30]

	fmt.Println(a, b) // [10 20 30 40 50] [20 30]

	b[0] = 999
	// a 變成 [10, 999, 30, 40, 50]，因為 b 改的是同一塊記憶體 (slice 出來的結果跟原本的 array 共用底層記憶體)
	fmt.Println(a, b) // [10 999 30 40 50] [999 30]
}

func changeSliceItem(words []string) {
	words[0] = "Hi"
}

func printLenCap(nums []int) {
	fmt.Printf("len: %d, cap: %d %v\n", len(nums), cap(nums), nums)
}

func foo(a []int) {
	a = append(a, 1, 2, 3, 4, 5, 6, 7, 8)
	a[0] = 200
	fmt.Println(a)
}
