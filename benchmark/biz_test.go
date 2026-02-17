package benchmark_test

import (
	"testing"

	"github.com/DucTran999/go-concurrency-patterns/benchmark"
)

func BenchmarkGreeting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmark.Greeting()
	}
}
