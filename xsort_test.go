package xsort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	f = func(x, y interface{}) bool { return x.(int) < y.(int) }
)

const (
	Xsize = 1000
)

func shuffle(xs []int) {
	for i, _ := range xs {
		xs[i] = r.Intn(len(xs))
	}
}

func BenchmarkBubble(b *testing.B) {
	xs := make([]int, Xsize)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		shuffle(xs)
		b.StartTimer()

		Bubble(xs, func(i, j int) bool { return xs[i] < xs[j] })
	}
}

func BenchmarkInsertion(b *testing.B) {
	xs := make([]int, Xsize)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		shuffle(xs)
		b.StartTimer()

		Insertion(xs, func(i, j int) bool { return xs[i] < xs[j] })
	}
}

func BenchmarkMerge(b *testing.B) {
	xs := make([]int, Xsize)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		shuffle(xs)
		b.StartTimer()

		Merge(xs, func(i, j int) bool { return xs[i] < xs[j] })
	}
}

func BenchmarkQuick(b *testing.B) {
	xs := make([]int, Xsize)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		shuffle(xs)
		b.StartTimer()

		xs = Quick(xs, f).([]int)
	}
}

func BenchmarkSelection(b *testing.B) {
	xs := make([]int, Xsize)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		shuffle(xs)
		b.StartTimer()

		Selection(xs, func(i, j int) bool { return xs[i] < xs[j] })
	}
}

func BenchmarkStdSort(b *testing.B) {
	xs := make([]int, Xsize)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		shuffle(xs)
		b.StartTimer()

		sort.Slice(xs, func(i, j int) bool {
			return xs[i] < xs[j]
		})
	}
}
