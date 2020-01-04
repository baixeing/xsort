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
		xs := reflect.MakeSlice(left.Type(), length, length)

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
				j = n
			}
		}
	}
}

func Selection(xs interface{}, f func(interface{}, interface{}) bool) {
	v := reflect.ValueOf(xs)

	for i := 0; i < v.Len()-1; i++ {
		n := i
		for j := i + 1; j < v.Len(); j++ {
			if f(v.Index(j).Interface(), v.Index(n).Interface()) {
				n = j
			}
		}

		x, y := v.Index(i).Interface(), v.Index(n).Interface()
		v.Index(i).Set(reflect.ValueOf(y))
		v.Index(n).Set(reflect.ValueOf(x))
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
