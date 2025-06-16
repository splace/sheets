package sheets

import "iter"

// only for the same type (can be any)
func Concat[T any](rs ...iter.Seq[T]) iter.Seq[T] {
    return func(yield func(T) bool) {
        for _, r := range rs {
            for e := range r {
                if !yield(e) {
                    return
                }
            }
        }
    }
}

// starts after a number of elements from the provided sequence
func After[T any](ts iter.Seq[T],start uint) iter.Seq[T] {
	return func(yield func(T) bool) {
		next, stop := iter.Pull(ts)
		defer stop()
		var ok bool
		for range start{
			_, ok = next()
			if !ok{
				return
			}
		}
		for {
			v, ok := next()
			if !ok || !yield(v) {
				return
			}
		}
	}
}

// return at most a number of elements, from the provided sequence
func Limit[T any](ts iter.Seq[T],n uint) iter.Seq[T] {
	return func(yield func(T) bool) {
		next, stop := iter.Pull(ts)
		defer stop()
		for range n{
			v, ok := next()
			if !ok || !yield(v) {
				return
			}
		}
	}
}


// sequence of elements from stepping through the provided sequence
func Step[T any](ts iter.Seq[T],step uint) iter.Seq[T] {
	return func(yield func(T) bool) {
		next, stop := iter.Pull(ts)
		defer stop()
		for {
			for i:=range step{
				v, ok := next()
				if !ok || i==0 && !yield(v) {
					return
				}
			}
		}
	}
}

// returns a sequence composed of all the i'th elements from each sequence in the provded sequence of sequences.
func Ats[T any](ts iter.Seq[iter.Seq[T]],i uint) iter.Seq[T] {
	return func(yield func(T) bool) {
		for r:=range ts{
			for c:=range After(r,i){
				if !yield(c){
					return
				}
				break
			}
		}
	}
}

// returns a selection from a sequence of sequences, where the sequences i'th element matches a func.
func Select[T any](ts iter.Seq[iter.Seq[T]],i uint,isMatch func(T)bool) iter.Seq[iter.Seq[T]] {
	return func(yield func(iter.Seq[T]) bool) {
		for r:=range ts{
			for t:= range After(r,i){
				if isMatch(t) && !yield(r) {
					return
				}
				break
			}
		}
	}
}
