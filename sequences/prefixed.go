package sequences

import "iter"
//import . "golang.org/x/exp/constraints"


func Prefix[T any](prefix iter.Seq[T], s iter.Seq[T]) iter.Seq[T] {
	next, stop := iter.Pull(prefix)
	var ok bool
	var f T
	return func(yield func(T) bool) {
		defer stop()
		for e := range s {
			f, ok = next()
			if !ok{
				return
			}
			if !yield(f) {
				return
			}
			if !yield(e) {
				return
			}
		}
	}
}

func Index[T Number](s iter.Seq[T]) iter.Seq[T] {
    return func(yield func(T) bool) {
    	var i T
        for e := range s {
			i++
			if !yield(i) {
				return
			}
			if !yield(e) {
				return
			}
        }
    }
}

