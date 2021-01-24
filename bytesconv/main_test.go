package main

import (
	"testing"
)

func BenchmarkStoB1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = stob1("1000000")
		// _ = string(x)
		// fmt.Println(x)
	}
}

func BenchmarkStoB2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = stob2("1000000")
		// _ = string(x)
		// fmt.Println(x)
	}
}

// 結果：
// bash
// $ go test -v -bench=. -run=none .
// goos: darwin
// goarch: amd64
// BenchmarkPrintInt2String01
// BenchmarkPrintInt2String01-4    1000000000          0.342 ns/op
// BenchmarkStoB1
// BenchmarkStoB1-4                158921668          8.08 ns/op
// BenchmarkStoB2
// BenchmarkStoB2-4                1000000000          0.370 ns/op
// PASS
// ok   _/Users/tih/localgitrepos/golang-byte-string-convert-benchmark 3.232s
