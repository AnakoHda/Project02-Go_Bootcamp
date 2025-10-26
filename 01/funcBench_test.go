package main

import (
	"testing"
)

func BenchmarkSortTask(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortTask(10000, 1)
	}
}

func BenchmarkMutexTask(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mutexTask(10000, 1)
	}
}
