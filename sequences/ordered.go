package sheets

import . "golang.org/x/exp/constraints"


func Between[T Ordered](lower, upper T) func(T) bool {
	return func(i T) bool {
		return i > lower && i < upper
	}
}


