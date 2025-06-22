package sheets

import "iter"
import "strings"
import "fmt"

type List[T any] iter.Seq[T]

func (l List[T]) String()string{
	var sb strings.Builder
	next,stop:=iter.Pull(iter.Seq[T](l))	
	defer stop()
	v,ok:=next()
	if !ok{
		return ""
	}
	fmt.Fprint(&sb,v)
	for {
		v,ok=next()
		if !ok{
			break
		}
		fmt.Fprint(&sb," ",v)
	}
	return sb.String()	
}



