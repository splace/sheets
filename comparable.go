package sheets

import "iter"


func Compare[T comparable](r1,r2 iter.Seq[T]) bool{
	next1, stop1 := iter.Pull(r1)
	next2, stop2 := iter.Pull(r2)
	defer stop1()
	defer stop2()
	for{
		v1, ok1 := next1()
		v2, ok2 := next2()
		if !ok1 && !ok2 {
			return true
		}
		if !ok1 || !ok2 || v1!=v2{
			return false
		}
	}
	return true
}


func UntilValue[T comparable](i T, in iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for p := range in {
			if !yield(p) || p == i {
				break
			}
		}
	}
}

