package table

import (
	"bufio"
//	"io/fs"
//	"fmt"
	"iter"
	"io"
)

//import . "../../sheet"

func LineScanner(r io.Reader) iter.Seq[string]{
	sr := bufio.NewScanner(r)
	sr.Split(Lines)
	return func(yield func(string) bool) {
		for sr.Scan() {
			if !yield(sr.Text()) {
				return
			}
		}
	}
}




//type Scanned struct{
//	io.Reader
//	
//}

//func (r Scanned) All() Row[string]{
//	scanner := bufio.NewScanner(r) // bufio.Scanner
//	return func(yield func(string) bool){
//		for scanner.Scan() {
//			if !yield(scanner.Text()){
//				return
//			}
//		}
//	}
//}

//bufio.ScanWords
//bufio.ScanLines
//scanner.ScanCommas
//scanner.ScanTabs
//scanner.ScanLinesAndCommas
//scanner.ScanLinesAndTabs

//type Dir fs.#


//ReadFileFS  // optimised for all

//ReadDirFile  // n at time




//func NewLined(s string) fs.File{
//	
//}






//type Rows struct{
//	fs.File
//}



//func (r Lines) Get(i int)any{
//	scanner := bufio.NewScanner(r) // bufio.Scanner
//	for scanner.Scan() {
//		i--
//		if i>0 {continue}
//		return scanner.Text()
//	}
//	return fmt.Errorf("Get %w",scanner.Err)
//}


	
//		fields := strings.Fields(line)

