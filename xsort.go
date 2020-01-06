package xsort

import (
	"reflect"
)

func Merge(xs interface{}, f func(int, int) bool) {
	var sort func(int, int) (int, int)
	var merge func(int, int, int, int)

	v := reflect.ValueOf(xs)
	//	swap := reflect.Swapper(xs)

	sort = func(start, end int) (int, int) {
		length := end - start

		if length < 2 {
			return start, end
		}

		pivot := start + length/2

		lstart, lend := sort(start, pivot)
		rstart, rend := sort(pivot, end)

		merge(lstart, lend, rstart, rend)

		return start, end
	}

	merge = func(lstart, lend, rstart, rend int) {
		start, end := lstart, rend
		ys := make([]interface{}, end-start, end-start)

		for i := 0; i < len(ys); i++ {
			if lstart < lend && (rstart == rend || f(lstart, rstart)) {
				ys[i] = v.Index(lstart).Interface()
				lstart++
			} else if rstart < rend {
				ys[i] = v.Index(rstart).Interface()
				rstart++
			}
		}

		for i := start; i < end; i++ {
			x := v.Index(i)
			x.Set(reflect.ValueOf(ys[i-start]))
		}
	}

	_, _ = sort(0, v.Len())
}

func Insertion(xs interface{}, f func(int, int) bool) {
	v := reflect.ValueOf(xs)
	swap := reflect.Swapper(xs)
	length := v.Len()

	for i := 0; i < length; i++ {
		for j := i; j > 0 && f(j, j-1); j-- {
			swap(j, j-1)
		}
	}
}

func Bubble(xs interface{}, f func(int, int) bool) {
	v := reflect.ValueOf(xs)
	swap := reflect.Swapper(xs)
	length := v.Len()

	for i, j := length, 0; i > 1; i, j = j, 0 {
		for n := 1; n < i; n++ {
			if f(n, n-1) {
				swap(n, n-1)
				j = n
			}
		}
	}
}

func Selection(xs interface{}, f func(int, int) bool) {
	v := reflect.ValueOf(xs)
	swap := reflect.Swapper(xs)
	length := v.Len()

	for i := 0; i < length-1; i++ {
		n := i
		for j := i + 1; j < length; j++ {
			if f(j, n) {
				n = j
			}
		}

		swap(i, n)
	}
}

func Quick(xs interface{}, f func(interface{}, interface{}) bool) interface{} {
	var sort func(reflect.Value) reflect.Value

	sort = func(xs reflect.Value) reflect.Value {
		if xs.Len() < 2 {
			return xs
		}

		left := reflect.MakeSlice(xs.Type(), 0, 0)
		middle := reflect.MakeSlice(xs.Type(), 0, 0)
		right := reflect.MakeSlice(xs.Type(), 0, 0)

		for i := 0; i < xs.Len(); i++ {
			switch {
			case xs.Index(0).Interface() == xs.Index(i).Interface():
				middle = reflect.Append(middle, xs.Index(i))
			case !f(xs.Index(0).Interface(), xs.Index(i).Interface()):
				left = reflect.Append(left, xs.Index(i))
			case f(xs.Index(0).Interface(), xs.Index(i).Interface()):
				right = reflect.Append(right, xs.Index(i))
			}
		}

		return reflect.AppendSlice(
			sort(left),
			reflect.AppendSlice(
				middle,
				sort(right),
			),
		)
	}

	return sort(reflect.ValueOf(xs)).Interface()
}
