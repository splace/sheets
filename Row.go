package sheets

import "iter"
import "fmt"
import "slices"
import "cmp"

//import "log"

type Row[T any] List[T]

// source slice is referencef so can be changed inplace
func NewRow[T any](rs ...T) Row[T]{
	return Row[T](slices.Values(rs))
}

func NewRowSeq[T any](rs ...T) Row[T]{
	return Row[T](
		func(yield func(T) bool) {
			for _,r:=range rs{
				if !yield(r){
					return
				}
			}
		},
	)
}

func NewReverseRow[T any](rs ...T) Row[T]{
	return Row[T](Reverse(rs))
}

func CompareRows[T comparable](r1,r2 Row[T]) bool{
	return Same(List[T](r1),List[T](r2))
}

func (r Row[T]) Cache() Row[T]{
	return Row[T](slices.Values(slices.Collect(List[T](r))))
}

func Sorted[T cmp.Ordered](r Row[T]) Row[T]{
	return NewRow(slices.Sorted(List[T](r))...)
}


func (r Row[T]) Reverse() Row[T]{
	t:=slices.Collect(List[T](r))
	slices.Reverse(t)
	return Row[T](slices.Values(t))
}

func (r Row[T]) At(i uint) (d T){
	for d=range After(List[T](r),i){
		break
	}
	return 
}	

func (r Row[T]) Items(is ...uint) Row[T]{
	iis:=make([]T,len(is))
	var ci uint
	for t:=range r{
		ci++
		for ii,i:=range is{
			if i==ci{
				iis[ii]=t
			}
		}
	}
	return NewRow(iis...)
}	



func (r Row[T]) Sample(d [2]uint) Row[T]{
	return Row[T](Step[T](After(List[T](r),d[0]),d[1]))
}	
	
func (r Row[T]) Sub(d [2]uint) Row[T]{
	return Row[T](Limit[T](After(List[T](r),d[0]),d[1]))
}

func (r Row[T]) Select(cs ...uint) Row[T]{
	return Row[T](Sub(List[T](r),cs...))
}



//func (r Row[T]) Sprint() Row[string]{
//	return r.Sprintf(nil)
//}

// produces a Row[string] from Row[any]
// obtains the sting using a Row of Formatter's
// if not available, string comes from a fmt.Formatter using the provided format strings (f0 for first,f for the rest), this is when;
// *	fmts==nil
// *	fmts has stopped
// *	Formatter==nil
func (r Row[T]) Sprintf(f0,f string,fmts Row[Formatter]) Row[string]{
	if fmts==nil{
		return func(yield func(string) bool) {
			next, stop := iter.Pull(List[T](r))
			defer stop()
			v, ok := next()
			if !ok || !yield(fmt.Sprintf(f0,v)) {
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
		next, stop := iter.Pull(List[T](r))
		nextf, stopf := iter.Pull(List[Formatter](fmts))
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
			if !yield(fmt.Sprintf(f0,v)) {
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
	// if the verb is a sep then use it as one
	//log.Printf("%T %q",r,verb)
	if slices.Contains([]rune{'\t','\n',',','.','/','\\','|'},verb){
		next, stop := iter.Pull(List[T](r))
		defer stop()
		v, ok := next()
		if !ok {
			return
		}
		fmt.Fprint(s,v)  // sep not needed before first
		for {
			v, ok := next()
			if !ok {
				return
			}
			fmt.Fprintf(s,"%c%v",verb,v)
		}
		return
	}
	// ...if not, pass on the verb as format for each item
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

