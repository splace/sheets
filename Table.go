package sheets

//import "fmt"
//import "iter"
import "slices"
//import "./sequences"
import "./lists"


type Table[U any,T Row[U]] lists.List[T]

func NewTable[U any,T Row[U]](rs ...T) Table[U,T]{
	return Table[U,T](lists.List[T](slices.Values(rs)))
}


//type IndexedRow[T] Row[T]

//func (ir IndexedRow[T])  struct{





