package sequences

import "runtime"
import "iter"
import . "golang.org/x/exp/constraints"

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


func CompareNotNil(r1,r2 iter.Seq[any]) bool{
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
		if !ok1 || !ok2 || v1==nil || v2==nil || v1!=v2{
			return false
		}
	}
	return true
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

// returns a sequence composed of all the i'th returns from each sequence in the provded sequence of sequences.
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

// returns a selection from a sequence of sequences, where the sequences i'th return matches a func.
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

func Reverse[Slice ~[]T, T any](s Slice) iter.Seq[T]{
	return func(yield func(T) bool) {
		for i:=len(s);i>-0;i--{
			if !yield(s[i]) {
				return
			}
		}
	}
}

// returns the required sequence items, optimised for random order and/or repeating 
func Sub[T any](ts iter.Seq[T],is ...uint) iter.Seq[T] {
	return func(yield func(T) bool) {
		var c uint
		found:=make(map[uint]T)
		next, stop := iter.Pull(ts)
		defer stop()
		for _,i:=range is{
			if m,got:=found[i];got{
				if !yield(m){
					return
				}
				continue
			}
			c++
			for ;;c++{
				t,ok:=next()
				if !ok{
					break
				}
				found[c]=t
				if m,got:=found[i];got{
					if !yield(m){
						return
					}
					break
				}				
			}
		}
	}
}

// returns the sequence items, at the indexes given by the values retuned from another sequence. optimised for random order and/or repeating 
func SubIter[T any, U Unsigned](s iter.Seq[T], is iter.Seq[U]) iter.Seq[T] {
	return func(yield func(T) bool) {
		c := U(1)
		found:=make(map[U]T)
		next, stop := iter.Pull(s)
		defer stop()
		for i:=range is{
			if m,have:=found[i];have{
				if !yield(m){
					return
				}
				continue
			}
			for ;;c++{
				t,ok:=next()
				if !ok{
					break
				}
				found[c]=t
				if m,got:=found[i];got{
					if !yield(m){
						return
					}
					c++
					break
				}				
			}
		}
	}
}


//func Interlace[T any](ss ...iter.Seq[T]) iter.Seq[T] {
//	return func(yield func(T) bool) {
//		nexts := make([]func() (T, bool), len(ss))
//		var stop func()
//		for i := range nexts {
//			nexts[i], stop = it.Pull[T](it.iter.Seq[T](ss[i]))
//			defer stop()
//		}
//		var v T
//		var g bool
//		for {
//			for _, next := range nexts {
//				v, g = next()
//				if !g {
//					return
//				}
//				if !yield(v) {
//					return
//				}
//			}
//		}
//	}
//}

//// returns nil if Row ends before a requested item is arrived at. 
//func Items[T Unsigned](is iter.Seq[T]) iter.Seq[T]{
//	for i:=range is{
//		if i==ci{
//				iis[ii]=t
//				needed--
//				if needed<1{
//					return NewRow(iis...)
//				}
//			}
//		}
//	}
//	return nil
//}	


func Until[T any](stop func(T) bool, in iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for p := range in {
			if !yield(p) || stop(p) {
				break
			}
		}
	}
}

// retains complete history, because stop cant be knowm to not use it all.
func UntilHistory[T any](stop func(...T) bool, in iter.Seq[T])iter.Seq[T] {
	return func(yield func(T) bool) {
		olds := make([]T, 0)
		for p := range in {
			olds = append(olds, p)
			if !yield(p) || stop(olds...) {
				break
			}
		}
	}
}

func FilterUntil[T any](selected, stop func(T) bool, in iter.Seq[T])iter.Seq[T] {
	return func(yield func(T) bool) {
		for p := range in {
			if selected(p) {
				if !yield(p) || stop(p) {
					break
				}
			}
		}
	}
}

func Filter[T any](selected func(T) bool, in iter.Seq[T])iter.Seq[T] {
	return func(yield func(T) bool) {
		for p := range in {
			if selected(p) {
				if !yield(p) {
					break
				}
			}
		}
	}
}


func Interlace[T any](ss ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		nexts := make([]func() (T, bool), len(ss))
		var stop func()
		for i := range nexts {
			nexts[i], stop = iter.Pull[T](iter.Seq[T](ss[i]))
			defer stop()
		}
		var v T
		var g bool
		for {
			for _, next := range nexts {
				v, g = next()
				if !g {
					return
				}
				if !yield(v) {
					return
				}
			}
		}
	}
}

//type iter[T any] = func(T) bool

//type iter2[T,U any] func(T,U)bool

//func (s iter.Seq[any]) Apply(with func(T)bool)string{
//	var b strings.Builder
//	for p:= range s{
//		fmt.Fprint(&b,sep,p)
//	}
//	return b.String()[len(sep):]
//}


//func Generate[T any, U any](fn func(i U) T) iter.Seq[T] {
//	return func(yield func(T) bool) {
//		var i U
//		for ; ; i++ {
//			if !yield(fn(i)) {
//				break
//			}
//		}
//	}
//}

//func Combine[T any, U any](apply func(...T) U, ss ...iter.Seq[T]) iter.Seq[U] {
//	return func(yield func(U) bool) {
//		nexts := make([]func() (T, bool), len(ss))
//		var stop func()
//		for i := range nexts {
//			nexts[i], stop = iter.Pull[T](iter.Seq[T](ss[i]))
//			defer stop()
//		}
//		v := make([]T, len(nexts))
//		var g bool
//		for {
//			for i, next := range nexts {
//				v[i], g = next()
//				if !g {
//					return
//				}
//			}
//			if !yield(apply(v...)) {
//				return
//			}
//		}
//	}
//}

// apply a func to all the values returned from all the seqences
func Amalgomate[T any, U any](apply func(...T) U, ss ...iter.Seq[T]) iter.Seq[U] {
	return func(yield func(U) bool) {
		nexts := make([]func() (T, bool), len(ss))
		var stop func()
		for i := range nexts {
			nexts[i], stop = iter.Pull[T](iter.Seq[T](ss[i]))
			defer stop()
		}
		v := make([]T, len(nexts))
		var g bool
		for {
			for i, next := range nexts {
				v[i], g = next()
				if !g {
					return
				}
			}
			if !yield(apply(v...)) {
				return
			}
		}
	}
}

func Repeat[T any](v T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for yield(v) {}
	}
}

func RepeatSequence[T any](s iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			for v := range s {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func Make[T any, U Number](fn func(i U) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		var i U
		for ; ; i++ {
			if !yield(fn(i)) {
				break
			}
		}
	}
}

func Apply[T any, V any](s iter.Seq[T], f func(T) V) iter.Seq[V] {
	cv := make(chan T, runtime.NumCPU())
	go func() {
		for t := range s {
			cv <- t
		}
		close(cv)
	}()
	return func(yield func(V) bool) {
		for v := range cv {
			if !yield(f(v)) {
				return
			}
		}
	}
}

func Append[T any](s iter.Seq[T], a ...T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range s {
			if !yield(v) {
				break
			}
		}
		for _, v := range a {
			if !yield(v) {
				return
			}
		}
	}
}

func Prepend[T any](s iter.Seq[T], p ...T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range p {
			if !yield(v) {
				return
			}
		}
		for v := range s {
			if !yield(v) {
				return
			}
		}
	}
}

func Delimit[T any](s iter.Seq[T], p ...T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range p[:len(p)/2] {
			if !yield(v) {
				return
			}
		}
		for v := range s {
			if !yield(v) {
				return
			}
		}
		for _, v := range p[len(p)/2:] {
			if !yield(v) {
				return
			}
		}
	}
}

func Interleave[T any](s iter.Seq[T], p ...T) iter.Seq[T] {
	return func(yield func(T) bool) {
		var i int
		for v := range s {
			if !yield(v) {
				return
			}
			if !yield(p[i%len(p)]) {
				return
			}
			i++
		}
	}
}

// separates a seq2 into two seq
// notce: reads the provided seq2 twice TODO buffer one? both?
func Split[T any](ss iter.Seq2[string,T]) (iter.Seq[string],iter.Seq[T]){
	return func(yield func(string)bool){
		for s,_:=range ss{
			if !yield(s){
				return
			}
		}
	},
	func(yield func(T)bool){
		for _,s :=range ss{
			if !yield(s){
				return
			}
		}
	}
}

