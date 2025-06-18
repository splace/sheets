package table

import (
	"unicode/utf8"
//	"strings"
//	"log"
	"unicode"
	"bytes"
	"bufio"
)



//func RuneSepFunc2(sep rune,pro ...func([]byte)[]byte) func([]byte,bool) (int,[]byte,error) {
//	if len(pro)==0{
//		return func(data []byte, atEOF bool) (int,[]byte,error) {
//			if atEOF && len(data) == 0 {
//				return 0, nil, nil
//			}
//			td:=data
//			for len(data)>0{
//				r,s:=utf8.DecodeRune(data)
//				if r==sep{
//					at:=len(td)-len(data)
//					return at+1,td[:at],nil
//				}
//				data=data[s:]
//			}
//			if atEOF {
//				return len(td), td[:len(td)-len(data)], nil
//			}
//			return 0, nil, nil
//		}
//	}	
//	p:=combine(pro...)
//	return func(data []byte, atEOF bool) (int,[]byte,error) {
//		if atEOF && len(data) == 0 {
//			return 0, nil, nil
//		}
//		td:=data
//		for len(data)>0{
//			r,s:=utf8.DecodeRune(data)
//			if r==sep{
//				at:=len(td)-len(data)
//				return at+1,p(td[:at]),nil
//			}
//			data=data[s:]
//		}
//		if atEOF {
//			return len(td), p(td[:len(td)-len(data)]), nil
//		}
//		return 0, nil, nil
//	}
//}


func RuneSepFunc(sep rune,pro ...func([]byte)[]byte) func([]byte,bool) (int,[]byte,error) {
	c:=combine(pro...)
	return func(bs []byte, atEOF bool) (int,[]byte,error) {
		i,bs,e:=EOFSepFunc(RuneSep(sep))(bs,atEOF)
		return i,c(bs),e
	}
}


func RuneSep(sep rune) func([]byte) (int,[]byte) {
	return func(data []byte) (int,[]byte){
		td:=data
		for len(data)>0{
			r,s:=utf8.DecodeRune(data)
			if r==sep{
				at:=len(td)-len(data)
				return at+1,td[:at]
			}
			data=data[s:]
		}
		return 0, nil
	}
}

// returns a sep func that returns upto a sep and also from last sep to EOF, as tokens
func EOFSepFunc(sf func(data []byte) (int,[]byte))func(data []byte, atEOF bool) (int,[]byte,error) {
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

type EOFScanner struct{
	bufio.Scanner
}

func (s EOFScanner) Scan() bool{
//	if atEOF && len(data) == 0 {
//		return 0, nil, nil
//	}
//	if atEOF {
//		return len(td), p(td[:len(td)-len(data)]), nil
//	}
//	return 0, nil, nil
	return false
}

//type Scanner struct{
//	bufio.Scanner
//	postProcessor func([]byte) []byte
//}

//func (s Scanner) Text() string{
//	return string(s.postProcessor(s.Scanner.Bytes()))
//}

type StringScanner struct{
	bufio.Scanner
	stringPostProcessor func(string) string
}

func (s StringScanner) Text() string{
	if s.stringPostProcessor==nil{
		return s.Scanner.Text()
	}
	return s.stringPostProcessor(s.Scanner.Text())
}

const commment ="//"

var LinesUniversal = RuneSepFunc('\n',CutSuffix("\r"),BeforeString(commment)) 

var Lines = RuneSepFunc('\n',BeforeString(commment)) 

func combine(fs ...func([]byte)[]byte)func([]byte)[]byte{
	return func(in []byte)[]byte{
		for _,f:=range fs{
			in=f(in)	
		}
		return in
	}
}

func Trim(d []byte) []byte{
	return TrimLeft(TrimRight(d))
}

func TrimLeft(d []byte) []byte{
	for {
		r,s:=utf8.DecodeRune(d)
		if !unicode.IsSpace(r){
			break
		}
		d=d[s:]
	}
	return d
}

func TrimRight(d []byte) []byte{
	for {
		r,s:=utf8.DecodeLastRune(d)
		if !unicode.IsSpace(r){
			break
		}
		d=d[:len(d)-s]
	}
	return d
}

//func Before(s string) func([]byte) []byte{
//	return func(d []byte) []byte{
//		if i:=strings.Index(string(d),s);i>=0{
//			return d[:i]
//		}
//		return d
//	}
//}

func CutSuffix(s string)  func([]byte) []byte{
	return Exec(s,bytes.CutSuffix)
}

func Exec(s string,op func(s,o []byte) ([]byte,bool)) func([]byte) []byte{
	bs:=[]byte(s)
	return func(d []byte) []byte{
		if rb,ok:=op(d,bs);ok{
			return rb
		}
		return d
	}
}

func MatchString(s string,op func(s,o []byte) bool) func([]byte) []byte{
	bs:=[]byte(s)
	return func(d []byte) (_ []byte){
		if op(d,bs){
			return 
		}
		return d
	}
}

func BeforeString(s string) func([]byte) []byte{
	bs:=[]byte(s)
	return func(d []byte) (_ []byte){
		if b,_,ok:=bytes.Cut(d,bs);ok{
			return b
		}
		return d
	}
}

func AfterString(s string) func([]byte) []byte{
	bs:=[]byte(s)
	return func(d []byte) (_ []byte){
		if _,a,ok:=bytes.Cut(d,bs);ok{
			return a
		}
		return d
	}
}

