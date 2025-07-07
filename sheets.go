package sheets

import "fmt"
import "strings"

type Sheets[U any,T Sheet[U,Row[U]]] map[string]T

func (ss Sheets[U,T]) String() string{
	var sb strings.Builder
	for n,s:=range ss{
		fmt.Fprintf(&sb,"%q\n%v",n,s)
	}
	return sb.String()
}



func GroupBy[U comparable,T Row[U]](s Sheet[U,T], group func(Row[U])bool) Sheets[U,Sheet[U,Row[U]]]{
	return Sheets[U,Sheet[U,Row[U]]]{
		"1":Sheet[U,Row[U]]{
			func(yield func(Row[U]) bool) {
				for r:=range s.Row{
					if group(Row[U](r)) && !yield(Row[U](r)){
						return
					}
				}
			},
		},
	}
}

func GroupBy2[U comparable,T Row[U]](s Sheet[U,T], matches Row[func(Row[U])bool]) Row[Sheet[U,T]]{
	return Row[Sheet[U,T]](
		func(yield func(Sheet[U,T]) bool) {
			for match:=range matches{
				if !yield(SelectRowsFunc[U,T](s,match)){
					return
				}
			}
		},
	)
}


