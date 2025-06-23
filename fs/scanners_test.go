package table

import (
	"fmt"
	"bytes"
	"bufio"
	"strings"
)

import "./sepfuncs"

func ExampleRune(){
	scanner := bufio.NewScanner(strings.NewReader(" 123 , 124\t\t,\t234, ")) // bufio.Scanner
	scanner.Split(sepfunc.Rune(','))
	for scanner.Scan() {
		fmt.Printf("%q\n",Trim(scanner.Bytes()))
	}
	// Output:
	// "123"
	// "124"
	// "234"
	// ""
}

//func ExampleRuneSepFunc(){
//	scanner := bufio.NewScanner(strings.NewReader(" 123 , 124 ")) // bufio.Scanner
//	scanner.Split(RuneSepFunc(','))
//	for scanner.Scan() {
//		fmt.Printf("%q\n",Trim(scanner.Bytes()))
//	}
//	// Output:
//	// "123"
//	// "124"
//}

// scan csv with comments
func ExampleBeforeString(){
	r:=strings.NewReader(" 123 , 124 ## comment\n 1,2\n3,  4  ## comment") 	
	lscanner := bufio.NewScanner(r) // std.lib line scan
//	lscanner.Split(SepFunc('\n')) 
	for lscanner.Scan() {
		cscanner := bufio.NewScanner(bytes.NewReader(lscanner.Bytes()))
		cscanner.Split(sepfunc.Rune(',',BeforeString("##"),Trim))
		fmt.Print("|")
		for cscanner.Scan() {
			fmt.Printf("%q\t\t|",cscanner.Text())
		}
		fmt.Println()
	}
	// Output:
	// |"123"		|"124"		|
	// |"1"		|"2"		|
	// |"3"		|"4"		|
}

func ExampleBeforeString_windows(){
	r:=strings.NewReader(" 123 , 124 // comment\r\n 1,2\r\n3,  4  // comment") 	
	lscanner := bufio.NewScanner(r)
	lscanner.Split(Lines)
//	lscanner.Split(SepFunc('\n')) 
	for lscanner.Scan() {
//		fmt.Print(len(lscanner.Text()))
		cscanner := bufio.NewScanner(bytes.NewReader(lscanner.Bytes()))
		cscanner.Split(sepfunc.Rune(',',Trim))
		fmt.Print("|")
		for cscanner.Scan() {
//			fmt.Printf("%q\t\t|",cscanner.Text())
			fmt.Printf("%q\t\t|",cscanner.Text())
		}
		fmt.Println()
	}
	// Output:
	// |"123"		|"124"		|
	// |"1"		|"2"		|
	// |"3"		|"4"		|
}

//// scan csv with comments
//func ExampleBefore_windows(){
//	r:=strings.NewReader(" 123 , 124 // comment\r\n 1,2\r\n3,  4  // comment") 	
//	lscanner := bufio.NewScanner(r)
////	lscanner.Split(SepFunc('\n')) 
//	var b bytes.Buffer
//	cscanner := bufio.NewScanner(&b)
//	cscanner.Split(SepFunc(',',Before("//"),Trim))
//	for lscanner.Scan() {
//		fmt.Println(lscanner.Text())
//		b.Reset()
//		b.Write(lscanner.Bytes())	
//		fmt.Print("|")
//		for cscanner.Scan() {
//			fmt.Printf("%q\t\t|",cscanner.Text())
//		}
//		fmt.Println()
//	}
//	// Output:
//	// |"123"		|"124"		|
//	// |"1"		|"2"		|
//	// |"3"		|"4"		|
//}

// todo bench
