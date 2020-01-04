# xsort
Simple implementation of sorting algorithms ([bubble](https://en.wikipedia.org/wiki/Bubble_sort), [insertion](https://en.wikipedia.org/wiki/Insertion_sort), [merge](https://en.wikipedia.org/wiki/Merge_sort), [quick](https://en.wikipedia.org/wiki/Quicksort), [selection](https://en.wikipedia.org/wiki/Selection_sort)) in Go (using [reflect](https://golang.org/pkg/reflect/)).

### example (for []int)
```go
package main

import (
    "fmt"
    "math/rand"
    "time"

    "github.com/baixeing/xsort"
)

const (
    Xsize = 10
)

func main() {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))

    bs := make([]int, Xsize)
    is := make([]int, Xsize)
    ms := make([]int, Xsize)
    qs := make([]int, Xsize)
    ss := make([]int, Xsize)

    for i := 0; i < Xsize; i++ {
        bs[i] = r.Intn(100)
        is[i] = r.Intn(100)
        ms[i] = r.Intn(100)
        qs[i] = r.Intn(100)
        ss[i] = r.Intn(100)
    }

    fmt.Println("--- unsorted ---")
    fmt.Printf("bs = %#+3.1v\n", bs)
    fmt.Printf("is = %#+3.1v\n", is)
    fmt.Printf("ms = %#+3.1v\n", ms)
    fmt.Printf("qs = %#+3.1v\n", qs)
    fmt.Printf("ss = %#+3.1v\n", ss)

    f := func(x, y interface{}) bool { return x.(int) < y.(int) }

    xsort.Bubble(bs, f)
    xsort.Insertion(is, f)
    ms = xsort.Merge(ms, f).([]int)
    qs = xsort.Quick(qs, f).([]int)
    xsort.Selection(ss, f)

    fmt.Println("--- sorted ---")
    fmt.Printf("bs = %#+3.1v\n", bs)
    fmt.Printf("is = %#+3.1v\n", is)
    fmt.Printf("ms = %#+3.1v\n", ms)
    fmt.Printf("qs = %#+3.1v\n", qs)
    fmt.Printf("ss = %#+3.1v\n", ss)
}
```
