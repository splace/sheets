package lists

import "iter"
import "strings"
import "fmt"
import "bytes"
import "io"

// List uses whitespace item separation.
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

func (l *List[T]) Scan(state fmt.ScanState, verb rune) (err error){
	// buffer the runes from state since, unfortunately, state gets modified by caller, fmt package, after this call. maybe dont use for large lists with early return.
	var b bytes.Buffer
	if _,err=copyRunes(&b,state);err!=nil{
		return
	}
	*l=List[T](
		func(yield func(T) bool) {
			var v T
			for{
				_,err=fmt.Fscan(&b,&v)
				if err!=nil || !yield(v) {
					return
				}
			}
		},
	)
	return
}

func copyRunes(rw interface{WriteRune(rune)(int,error)},rr io.RuneReader) (n int, err error){
	var r rune
	for{
		r,_,err=rr.ReadRune()
		if err!=nil{
			break
		}
		_,err=rw.WriteRune(r)
		if err!=nil{
			return
		}
		n++
	}
	if err==io.EOF{
		err=nil
	}
	return
}
