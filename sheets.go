package sheets

import "fmt"
import "strings"

type Sheets[T Sheet[Row[U],U], U any] map[string]T

func (ss Sheets[T,U]) String() string{
	var sb strings.Builder
	for n,s:=range ss{
		fmt.Fprintf(&sb,"%q\n%v",n,s)
	}
	return sb.String()
}



func GroupBy[T Row[U],U comparable](s Sheet[T,U], group func(Row[U])bool) Sheets[Sheet[Row[U],U],U]{
	return Sheets[Sheet[Row[U],U],U]{
		"1":Sheet[Row[U],U]{
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



