package sheets

import "fmt"
import "iter"
import "slices"
import "./sequences"



type Sheet[T Row[U],U any] struct{
	Row[T]
}

func NewSheet[T Row[U],U any](rs ...T) Sheet[T,U]{
	return Sheet[T,U]{Row[T](slices.Values(rs))}
}

func (s Sheet[T,U]) Column(i uint) T{
	return func(yield func(U) bool) {
		for r:=range s.Row{
			if !yield(Row[U](r).At(i-1)){
				return
			}
		}
	}
}

func (s Sheet[T,U]) SelectColumns(cs ...uint) Sheet[T,U]{	
	return Sheet[T,U]{
		func(yield func(T) bool) {
			for r:=range s.Row{
				if !yield(T(Row[U](r).Select(cs...))){
					return
				}
			}
		},
	}
}

func CompareSheets[T Row[U],U comparable](s1,s2 Sheet[T,U]) bool{
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


func SelectMatchedRows[T Row[U],U comparable](s Sheet[T,U], match Row[U] ) Sheet[T,U]{	
	return SelectRowsFunc(s,
		func(r Row[U])bool{
			return CompareRows(r,match)
		},
	)
}


func SelectRows[T Row[U],U comparable](s Sheet[T,U], col uint, v U ) Sheet[T,U]{
	return SelectRowsFunc(s,func(r Row[U])bool{return r.At(col)==v})
}

func SelectRowsFunc[T Row[U],U comparable](s Sheet[T,U], match func(Row[U])bool) Sheet[T,U]{
	return Sheet[T,U]{
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
	Sheet[Row[U],U]
}

func NewHeadedSheet[T any](ss iter.Seq2[string,T]) HeadedSheet[T]{
	r,s:=sequences.Split(ss)
	return HeadedSheet[T]{Row[string](r),NewSheet[Row[T],T](Row[T](s))}
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



func SelectRowsFrom[T Row[U],U comparable](f,s Sheet[T,U], col uint, v U ) Sheet[T,U]{
	return SelectRowsFromFunc(f,s,func(r Row[U])bool{return r.At(col)==v})
}

func SelectMatchedRowsFrom[T Row[U],U comparable](f,s Sheet[T,U], match Row[U] ) Sheet[T,U]{	
	return SelectRowsFromFunc(f,s,
		func(r Row[U])bool{
			return CompareRows(r,match)
		},
	)
}

func SelectRowsFromFunc[T Row[U],U comparable,F Row[G],G any](f Sheet[F,G],s Sheet[T,U], match func(Row[U])bool) Sheet[F,G]{
	var i uint
	return Sheet[F,G]{
		func(yield func(F) bool) {
			for r:=range s.Row{
				if match(Row[U](r)) && !yield(f.At(i)){
					return
				}
				i++
			}
		},
	}
}


