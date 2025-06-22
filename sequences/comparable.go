package sequences

import "iter"


// all sequences match whats returned from first.
// extra ignored, not pulled.
func Matched[T comparable](s1 iter.Seq[T], ss ...iter.Seq[T]) (_ bool){
	var nss []func()(T,bool)
	next,stop:=iter.Pull(s1)
	defer stop()
	lt,ok:=next()
	if !ok{return}
	var t T
	for _,s:=range ss{
		n,stop:=iter.Pull(s)
		defer stop()
		t,ok=n()
		if !ok || t!=lt{return}
		nss=append(nss,n)
	}
	for {
		lt,ok=next()
		if !ok{return true}
		for _,n:=range nss{
			t,ok=n()
			if !ok || t!=lt{return}
		}
	}
	return true
}

func Same[T comparable](r1,r2 iter.Seq[T]) bool{
	next1, stop1 := iter.Pull(r1)
	next2, stop2 := iter.Pull(r2)
	defer stop1()
	defer stop2()
	for{
		v1, ok1 := next1()
		v2, ok2 := next2()
		if !ok1 && !ok2 {
			return true
		}
		if !ok1 || !ok2 || v1!=v2{
			return false
		}
	}
	return true
}

//// 
//func Compare[T comparable](s iter.Seq[T],ss ...iter.Seq[T]) bool{
//	for _,os:=range ss{
//		if !compare(s,os){
//			return false
//		}
//	}
//	return true
//}


func UntilValue[T comparable](i T, in iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for p := range in {
			if !yield(p) || p == i {
				break
			}
		}
	}
}

