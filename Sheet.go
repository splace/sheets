package sheets

import "fmt"
import "iter"
import "slices"

type Sheet[T Row[U],U any] struct{
	Row[T]
}

func NewSheet[T Row[U],U any](rs ...T) Sheet[T,U]{
	return Sheet[T,U]{Row[T](slices.Values(rs))}
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
		if !ok1 || !ok2 || !Compare(iter.Seq[U](r1),iter.Seq[U](r2)){
			return false
		}
	}
	return true
}


//func (s Sheet[T,U]) Column(i uint) T{
//	f:=iter.Seq[iter.Seq[U]T](s.Row)
//	g:=iter.Seq[T](f)
//	return T(Ats[U](g,i))
//}

type HeadedSheet[U any] struct{
	Row[string]
	Sheet[Row[U],U]
}

func (ht HeadedSheet[T]) Format(s fmt.State, _ rune ){
//	//fmt.Fprintf(s,fmt.Sprintf("%%%c",verb)+"\n",ht.Row)
//	//fmt.Fprintf(s,"%\n",ht.Sheet)
//	fmt.Fprintf(s,"%\n",ConcatRows[string](NewRow(ht.Row).Sprintf("%10v","|%v",nil),ht.Sheet.Sprintf("%v","|%v",nil)))
	fmt.Fprintf(s,"%\n",Row[string](Concat[string](iter.Seq[string](NewRow(ht.Row).Sprintf("%10v","|%v",nil)),iter.Seq[string](ht.Sheet.Sprintf("%v","|%v",nil)))))
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

