package sheets

import "iter"
import "fmt"
import "slices"

type Row[T any] iter.Seq[T]

// source slice is referencef so can be changed inplace
func NewRow[T any](rs ...T) Row[T]{
	return Row[T](slices.Values(rs))
}

func Compare[T comparable](r1,r2 Row[T]) bool{
	return true
}

func (r Row[T]) Cache() Row[T]{
	return Row[T](slices.Values(slices.Collect(iter.Seq[T](r))))
}

func (r Row[T]) Reverse() Row[T]{
	t:=slices.Collect(iter.Seq[T](r))
	slices.Reverse(t)
	return Row[T](slices.Values(t))
}

func (r Row[T]) At(i uint) (d T){
	for d=range After(iter.Seq[T](r),i){
		break
	}
	return 
}	

func (r Row[T]) Sample(d [2]uint) Row[T]{
	return Row[T](Step[T](After(iter.Seq[T](r),d[0]),d[1]))
}	
	
func (r Row[T]) Sub(d [2]uint) Row[T]{
	return Row[T](Limit[T](After(iter.Seq[T](r),d[0]),d[1]))
}


//func (r Row[T]) Sprint() Row[string]{
//	return r.Sprintf(nil)
//}

// produces a Row[string] from Row[any]
// uses a sequence of Formatter's
// defaults to fmt.Formatter when;
// *	Formatter==nil
// *	Row[Formatter]==nil
// *	Row[Formatter] has stopped
func (r Row[T]) Sprintf(f1,f string,fmts Row[Formatter]) Row[string]{
	if fmts==nil{
		return func(yield func(string) bool) {
			next, stop := iter.Pull(iter.Seq[T](r))
			defer stop()
			v, ok := next()
			if !ok || !yield(fmt.Sprintf(f1,v)) {
				return
			}
			for {
				v, ok = next()
				if !ok || !yield(fmt.Sprintf(f,v)) {
					return
				}
			}
		}
	}
	return func(yield func(string) bool) {
		next, stop := iter.Pull(iter.Seq[T](r))
		nextf, stopf := iter.Pull(iter.Seq[Formatter](fmts))
		defer stop()
		defer stopf()
		v, ok := next()
		if !ok {
			return
		}
		fr, ok := nextf()
		if fr!=nil && ok {
			if !yield(fr(v)) {
				return
			}
		}else{
			if !yield(fmt.Sprintf(f1,v)) {
				return
			}
		}
		for {
			v, ok := next()
			if !ok {
				return
			}
			fr, ok := nextf()
			if fr!=nil && ok {
				if !yield(fr(v)) {
					return
				}
				continue
			}else{
				if !yield(fmt.Sprintf(f,v)) {
					return
				}
			}
		}
	}
}


//func (r Row[T]) String()string{
//	return fmt.Sprintf("%v",r)
//}


func (r Row[T]) Format(s fmt.State, verb rune ){
	// if the verb is a sep then use it between outputs with default format
	if slices.Contains([]rune{'\t','\n',',','.','/','\\','|'},verb){
		next, stop := iter.Pull(iter.Seq[T](r))
		defer stop()
		v, ok := next()
		if !ok {
			return
		}
		fmt.Fprint(s,v)  // sep only needed before all but first
		for {
			v, ok := next()
			if !ok {
				return
			}
			fmt.Fprintf(s,"%c%v",verb,v)
		}
		return
	}
	// if not use it as format rune
	for v:=range r{
		fmt.Fprintf(s,fmt.Sprintf("%%%c",verb),v)
	}
}

type Formatter func(any)string

func format(f string) Formatter{
	if f=="%v"{
		return func(s any)string{
			return fmt.Sprint(s)
		}
	}
	return func(s any)string{
		return fmt.Sprintf(f,s)
	}
}

