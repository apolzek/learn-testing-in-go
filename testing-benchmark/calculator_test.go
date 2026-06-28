package main

import (
	"fmt"
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	fmt.Println(b.N)
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}
