package table

import (
	"unicode/utf8"
//	"strings"
//	"log"
	"unicode"
	"bytes"
	"bufio"
)

import "./sepfuncs"

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

var LinesUniversal = sepfunc.Rune('\n',CutSuffix("\r"),BeforeString(commment)) 

var Lines = sepfunc.Rune('\n',BeforeString(commment)) 

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

