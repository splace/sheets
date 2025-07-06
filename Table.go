package sheets

//import "fmt"
//import "iter"
import "slices"
//import "./sequences"
import "./lists"


type Table[T Row[U],U any] struct{
	lists.List[T]
}

func NewTable[T Row[U],U any](rs ...T) Table[T,U]{
	return Table[T,U]{lists.List[T](slices.Values(rs))}
}

//type IndexedRow[T] Row[T]

//func (ir IndexedRow[T])  struct{





