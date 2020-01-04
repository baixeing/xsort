package xsort

import (
	"math/rand"
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

		Bubble(xs, f)
	}
}

func BenchmarkInsertion(b *testing.B) {
	xs := make([]int, Xsize)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		shuffle(xs)
		b.StartTimer()

		Insertion(xs, f)
	}
}

func BenchmarkMerge(b *testing.B) {
	xs := make([]int, Xsize)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		shuffle(xs)
		b.StartTimer()

		xs = Merge(xs, f).([]int)
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

		Selection(xs, f)
	}
}
