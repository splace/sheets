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

//func (s Sheet[T,U]) Column(i uint) T{
//	return T(Ats[U](s.Row,i))
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

