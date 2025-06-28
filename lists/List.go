package lists

import "iter"
import "strings"
import "fmt"
//import "bytes"
//import "io"
//import "sync"
//import "bytes"
//import "unicode"
//import "log"

var Sep,LSep = " ","\n"
 
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
	switch any(v).(type){
	case List[int]:
		fmt.Fprintln(&sb,v)
		for {
			v,ok=next()
			if !ok{
				break
			}
			fmt.Fprintln(&sb,v)
		}
	default:
		fmt.Fprint(&sb,v)
		for {
			v,ok=next()
			if !ok{
				break
			}
			sb.WriteString(Sep)
			fmt.Fprint(&sb,v)
		}
		
	}
	return sb.String()	
}



type Strings []string

func not(f func(rune)bool)func(rune)bool{return func(r rune)bool{return !f(r)}}

func match(v rune)func(rune)bool{return func(r rune)bool{return r==v}}

func or(fs ...func(rune)bool)func(r rune)bool {
	return func(r rune)bool{
		for _,f :=range fs{	
			if f(r) {
				return true
			}
		}
		return false
	}
}

func and(fs ...func(rune)bool)func(r rune)bool {
	return func(r rune)bool{
		for _,f :=range fs{	
			if !f(r) {
				return false
			}
		}
		return true
	}
}


func (ls *Strings) Scan(state fmt.ScanState, verb rune) (err error){
	var lss []string
	var t []byte
	switch verb{
	case '\t','\n':
		for {
			t,err=state.Token(false,not(match(verb))) // dont skip whitespace verb here is one, is skipped miss empty items
			if err!=nil || len(t)==0{
				break
			}
			lss=append(lss,string(t))
			state.ReadRune()
		}
	case ',','.','/','\\','|':
		for {
			t,err=state.Token(true,not(match(verb)))
			if err!=nil || len(t)==0{
				break
			}
			lss=append(lss,string(t))
			state.ReadRune()
		}
	case 'v':
		fallthrough
	default:
		for {
			t,err=state.Token(true,nil)
			if err!=nil || len(t)==0{
				break
			}
			lss=append(lss,string(t))
		}
	}
	*ls=lss
	return
}

// store all tokens then parses to type as pulled
func (l *List[T]) Scan(state fmt.ScanState, verb rune) (err error){
	var ss Strings
	err=ss.Scan(state,verb)
	*l=To[T](ss...)
	return
}


func To[T any](ss ...string) List[T]{
	return List[T](
		func(yield func(T) bool) {
			for _,s:=range ss{
				var v T
				_,err:=fmt.Sscan(s,&v)
				if err!=nil || !yield(v) {
					return
				}
			}
		},
	)
}

//func (l *List[T]) Scan2(state fmt.ScanState, verb rune) (err error){
//	// buffer the runes from state since, unfortunately, state gets modified by caller, fmt package, after this call. maybe dont use for large lists with early return.
//	var b bytes.Buffer
//	if _,err=copyRunes(&b,state);err!=nil{
//		return
//	}
//	*l=List[T](
//		func(yield func(T) bool) {
//			var v T
//			for{
//				_,err=fmt.Fscan(&b,&v)
//				if err!=nil || !yield(v) {
//					return
//				}
//			}
//		},
//	)
//	return
//}

//func (l *List[T]) Scan3(state fmt.ScanState, verb rune) (err error){
//	var m sync.Mutex
//	*l=List[T](
//		func(yield func(T) bool) {
//			var v T
//			for{
//				buf,err:=state.Token(true,nil)
//				if err!=nil{
//					break
//				}
//				_,err=fmt.Sscan(string(buf),&v)
////				_,err=fmt.Fscan(state,&v)
//				if err!=nil || !yield(v) {
//					break
//				}
//			}
//			m.Unlock()
//		},
//	)
//	m.Lock()
//	return
//}



//func copyRunes(rw interface{WriteRune(rune)(int,error)},rr io.RuneReader) (n int, err error){
//	var r rune
//	for{
//		r,_,err=rr.ReadRune()
//		if err!=nil{
//			break
//		}
//		_,err=rw.WriteRune(r)
//		if err!=nil{
//			return
//		}
//		n++
//	}
//	if err==io.EOF{
//		err=nil
//	}
//	return
//}
