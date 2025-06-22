package sequences

import "iter"

import . "golang.org/x/exp/constraints"

type Number interface {
	Integer | Float | Complex
}

type Addable interface {
	Number | ~string
}


func Runes(s string) iter.Seq[rune] {
	return func(yield func(rune) bool) {
		for _, r := range s {
			if !yield(r) {
				break
			}
		}
	}
}


// return a sequence of added items from the given sequence, all produced from different permutaions, ordered and increacing.
// ex  a,b,c.... a+a,a+b,a+c....b+a,b+b,b+c.... a+a+a,a+a+b,a+a+c...a+b+a,a+b+b,a+b+c...
func Permutations[T Addable](s iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for p := range s {
			if !yield(p) {
				break
			}
		}
		for p := range s {
			for q := range s {
				if !yield(p+q) {
					break
				}
			}
		}
	}
}

func Totalise[T Addable](in iter.Seq[T]) iter.Seq[T] {
	var t T
	return func(yield func(T) bool) {
		for p := range in {
			t += p
			if !yield(t) {
				break
			}
		}
	}
}

