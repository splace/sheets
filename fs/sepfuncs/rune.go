package sepfunc


import "unicode/utf8"



// rune sep including EOF + post processing
func Rune(sep rune,procs ...func([]byte)[]byte) func([]byte,bool) (int,[]byte,error) {
	sf:=EOF(RuneSep(sep))
	switch len(procs){
	case 0:
		return sf
	case 1:
		return Post(sf,procs[0])
	default:
		return Post(sf,combine(procs...))
	}
}

func Post(sf func([]byte,bool) (int,[]byte,error), pro func([]byte)[]byte) func([]byte,bool) (int,[]byte,error){
	return func(bs []byte, atEOF bool) (int,[]byte,error) {
		i,bs,e:=sf(bs,atEOF)
		return i,pro(bs),e
	}
}

func RuneSep(sep rune) func([]byte) (int,[]byte) {
	return RuneFunc(func(r rune)bool{return r==sep})
}

func RuneFunc(sep func(rune)bool) func([]byte) (int,[]byte) {
	return func(data []byte) (int,[]byte){
		td:=data
		for len(data)>0{
			r,s:=utf8.DecodeRune(data)
			if sep(r){
				at:=len(td)-len(data)
				return at+1,td[:at]
			}
			data=data[s:]
		}
		return 0, nil
	}
}


// returns a sep func that returns upto a sep AND also from last sep to EOF, as its tokens
func EOF(sf func(data []byte) (int,[]byte))func(data []byte, atEOF bool) (int,[]byte,error) {
	return func(data []byte, atEOF bool) (int,[]byte,error) {
		if !atEOF{
			i,data:=sf(data)
			return i,data,nil
		}
		return len(data),emptyAsNil(data),nil
	}
}

func emptyAsNil(b []byte) []byte{
	if len(b)==0{
		return nil
	}
	return b
}


func combine(fs ...func([]byte)[]byte)func([]byte)[]byte{
	return func(in []byte)[]byte{
		for _,f:=range fs{
			in=f(in)	
		}
		return in
	}
}

