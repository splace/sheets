package sheets

import "iter"

import . "golang.org/x/exp/constraints"


func Odds[T Integer]() iter.Seq[T] {
	return Geometric[T](1, 2)
}

func Evens[T Integer]() iter.Seq[T] {
	return Geometric[T](0, 2)
}


func Geometric[T Number](Start, Step T) iter.Seq[T] {
	return Totalise(Concat(Limit(Repeat[T](Start),1), Repeat(Step)))
}


func Fibonacci[T Number]() iter.Seq[T] {
	var v1, v2 T = 1, 1
	return func (yield func(T) bool) {
		if yield(v1) && yield(v2) {
			for v3 := v1 + v2; yield(v3); v3, v1, v2 = v3+v2, v2, v3 {}
		}
	}
}

func Multiply[T Number](s iter.Seq[T], scale T) iter.Seq[T] {
	return Apply(s, func(v T) T { return v * scale })
}

func Divide[T Number](s iter.Seq[T], scale T) iter.Seq[T] {
	return Apply(s, func(v T) T { return v / scale })
}

func Modify[T Integer](fn func(T) T, s func(T) bool) func(T) bool {
	return func(c T) bool {
		return s(fn(c))
	}
}

func Invert[T Integer](s func(T) bool) func(T) bool {
	return func(c T) bool {
		return !s(c)
	}
}

func StartStriper[T Integer](d T) func(T) bool {
	return func(c T) bool {
		return c < d
	}
}

func DashStriper[T Integer](d, s T) func(T) bool {
	dd := d / s
	return func(c T) bool {
		return (c % d) < dd
	}
}

func HalfStriper[T Integer](d T) func(T) bool {
	return DashStriper(d, 2)
}

func ThirdStriper[T Integer](d T) func(T) bool {
	return DashStriper(d, 3)
}

func MultiStriper[T Integer](ss ...func(T) bool) func(T) bool {
	return func(c T) bool {
		for _, s := range ss {
			if s(c) {
				return true
			}
		}
		return false
	}
}
