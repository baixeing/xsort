package xsort

import "reflect"

func Merge(xs interface{}, f func(interface{}, interface{}) bool) interface{} {
	var sort func(reflect.Value) reflect.Value
	var merge func(reflect.Value, reflect.Value) reflect.Value

	sort = func(xs reflect.Value) reflect.Value {
		if xs.Len() < 2 {
			return xs
		}

		return merge(
			sort(xs.Slice(0, xs.Len()/2)),
			sort(xs.Slice(xs.Len()/2, xs.Len())),
		)
	}

	merge = func(left, right reflect.Value) reflect.Value {
		length := left.Len() + right.Len()
		capacity := left.Cap() + right.Cap()
		xs := reflect.MakeSlice(left.Type(), length, capacity)

		for i := 0; i < length; i++ {
			v := xs.Index(i)
			if left.Len() > 0 && (right.Len() == 0 || f(left.Index(0).Interface(), right.Index(0).Interface())) {
				v.Set(left.Index(0))
				left = left.Slice(1, left.Len())
			} else if right.Len() > 0 {
				v.Set(right.Index(0))
				right = right.Slice(1, right.Len())
			}
		}

		return xs
	}

	return sort(reflect.ValueOf(xs)).Interface()
}

func Insertion(xs interface{}, f func(interface{}, interface{}) bool) {
	v := reflect.ValueOf(xs)

	for i := 0; i < v.Len(); i++ {
		for j := i; j > 0 && f(v.Index(j).Interface(), v.Index(j-1).Interface()); j-- {
			x, y := v.Index(j).Interface(), v.Index(j-1).Interface()
			v.Index(j).Set(reflect.ValueOf(y))
			v.Index(j - 1).Set(reflect.ValueOf(x))
		}
	}
}

func Bubble(xs interface{}, f func(interface{}, interface{}) bool) {
	v := reflect.ValueOf(xs)

	for i, j := v.Len(), 0; i > 1; i, j = j, 0 {
		for n := 1; n < i; n++ {
			x, y := v.Index(n).Interface(), v.Index(n-1).Interface()
			if f(x, y) {
				v.Index(n).Set(reflect.ValueOf(y))
				v.Index(n - 1).Set(reflect.ValueOf(x))
			}
		}
	}
}
