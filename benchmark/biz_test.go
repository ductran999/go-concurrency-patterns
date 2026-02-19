package benchmark_test

import (
	"testing"

	"go-concurrency-patterns/benchmark"
)

func BenchmarkGreeting(b *testing.B) {
	for b.Loop() {
		benchmark.Greeting()
	}
}
