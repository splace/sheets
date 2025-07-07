package sheets

import "fmt"
import "iter"
import "slices"
import "./sequences"

import "log"


type Sheet[U any,T Row[U]] struct{
	Row[T]
}

func NewSheet[U any,T Row[U]](rs ...T) Sheet[U,T]{
	return Sheet[U,T]{Row[T](slices.Values(rs))}
}

func (s Sheet[U,T]) Column(i uint) T{
	return func(yield func(U) bool) {
		for r:=range s.Row{
			if !yield(Row[U](r).At(i-1)){
				return
			}
		}
	}
}

func (s Sheet[U,T]) SelectColumns(cs ...uint) Sheet[U,T]{	
	return Sheet[U,T]{
		func(yield func(T) bool) {
			for r:=range s.Row{
				if !yield(T(Row[U](r).Select(cs...))){
					return
				}
			}
		},
	}
}

func CompareSheets[U comparable,T Row[U]](s1,s2 Sheet[U,T]) bool{
	next1, stop1 := iter.Pull(iter.Seq[T](s1.Row))
	next2, stop2 := iter.Pull(iter.Seq[T](s2.Row))
	defer stop1()
	defer stop2()
	for{
		r1, ok1 := next1()
		r2, ok2 := next2()
		if !ok1 && !ok2 {
			return true
		}
		if !ok1 || !ok2 || !sequences.Same(iter.Seq[U](r1),iter.Seq[U](r2)){
			return false
		}
	}
	return true
}


func SelectMatchedRows[U comparable,T Row[U]](s Sheet[U,T], match Row[U] ) Sheet[U,T]{	
	return SelectRowsFunc(s,
		func(r Row[U])bool{
			return CompareRows(r,match)
		},
	)
}


func SelectRows[U comparable,T Row[U]](s Sheet[U,T], col uint, v U ) Sheet[U,T]{
	return SelectRowsFunc(s,func(r Row[U])bool{return r.At(col)==v})
}

func SelectRowsFunc[U comparable,T Row[U]](s Sheet[U,T], match func(Row[U])bool) Sheet[U,T]{
	return Sheet[U,T]{
		func(yield func(T) bool) {
			for r:=range s.Row{
				if match(Row[U](r)) && !yield(r){
					return
				}
			}
		},
	}
}


type HeadedSheet[U any] struct{
	Row[string]
	Sheet[U,Row[U]]
}

func NewHeadedSheet[T any](ss iter.Seq2[string,T]) HeadedSheet[T]{
	r,s:=sequences.Split(ss)
	return HeadedSheet[T]{Row[string](r),NewSheet[T,Row[T]](Row[T](s))}
}

func (ht HeadedSheet[T]) Format(s fmt.State, _ rune ){
//	//fmt.Fprintf(s,fmt.Sprintf("%%%c",verb)+"\n",ht.Row)
//	//fmt.Fprintf(s,"%\n",ht.Sheet)
//	fmt.Fprintf(s,"%\n",ConcatRows[string](NewRow(ht.Row).Sprintf("%10v","|%v",nil),ht.Sheet.Sprintf("%v","|%v",nil)))
	fmt.Fprintf(s,"%\n",Row[string](sequences.Concat[string](iter.Seq[string](NewRow(ht.Row).Sprintf("%10v","|%v",nil)),iter.Seq[string](ht.Sheet.Sprintf("%v","|%v",nil)))))
}

type FormattedSheet[U any] struct{
	HeadedSheet[U]
	Row[Formatter]
}

//// need all Rows the same type (can be any)
//func ConcatRows[E any](rs ...Row[E]) Row[E] {
//    return func(yield func(E) bool) {
//        for _, r := range rs {
//            for e := range r {
//                if !yield(e) {
//                    return
//                }
//            }
//        }
//    }
//}



func SelectRowsFrom[U comparable,G any,T Row[U],F Row[G]](f Sheet[G,F],s Sheet[U,T], col uint, v U ) Sheet[G,F]{
	return SelectRowsFromFunc(f,s,func(r Row[U])bool{return r.At(col)==v})
}


func SelectMatchedRowsFrom[U comparable,G any,T Row[U],F Row[G]](f Sheet[G,F],s Sheet[U,T], match Row[U] ) Sheet[G,F]{	
	return SelectRowsFromFunc(f,s,
		func(r Row[U])bool{
			return CompareRows(r,match)
		},
	)
}

func SelectRowsFromFunc[U comparable,G any,T Row[U],F Row[G]](f Sheet[G,F],s Sheet[U,T], match func(Row[U])bool) Sheet[G,F]{
	var i uint
	return Sheet[G,F]{
		func(yield func(F) bool) {
			for r:=range s.Row{
				log.Print(r)
				log.Print(match(Row[U](r)))
				if match(Row[U](r)) && !yield(f.At(i)){
					return
				}
				i++
			}
		},
	}
}


