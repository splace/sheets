package table

import (
	"bufio"
	"os"
	"io/fs"
	"iter"
	"io"
//	"log"
//	"path"
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

func ReadLineScanner(sr iter.Seq[io.Reader]) iter.Seq[iter.Seq[string]]{
	return func(yield func(iter.Seq[string]) bool) {
		for r:=range sr {
			if !yield(LineScanner(r)) {
				return
			}
		}
	}
}


// os.Root.FS
func FolderScanner(s string) iter.Seq2[string,iter.Seq[string]]{
	//	return DirScanner(os.DirFS(s)) switched to Root for better isolation
	r,_:=os.OpenRoot(s)
	return DirScanner(r.FS())
}


// fs.DirEntry fs.FileInfo fs.File

// fs.ReadDirFile needs to be from root of fs (using fs.Sub()? )

// uses root of the provided fs (so path in fs, no need to pass in)
// NB. no particular order
// NB. opens files and returns a sequence of the lines in them.
func DirScanner(fsys fs.FS) iter.Seq2[string,iter.Seq[string]]{
	de,_:=fsys.Open(".")
	des:=de.(fs.ReadDirFile)
	return func(yield func(string,iter.Seq[string]) bool) {
		for e,err:=des.ReadDir(1);err==nil;e,err=des.ReadDir(1){
//			log.Print(de)
			fi,err:=e[0].Info()
			if err!=nil{
				break
			}
			f,err:=fsys.Open(fi.Name())
			if err!=nil{
				break
			}
			if !yield(fi.Name(),LineScanner(f)) {
				return
			}
		}
	}
}

func DirScanner2(fsys fs.FS) iter.Seq[io.Reader]{
	de,_:=fsys.Open(".")
	des:=de.(fs.ReadDirFile)
	return func(yield func(io.Reader) bool) {
		for e,err:=des.ReadDir(1);err==nil;e,err=des.ReadDir(1){
//			log.Print(de)
			fi,err:=e[0].Info()
			if err!=nil{
				break
			}
			f,err:=fsys.Open(fi.Name())
			if err!=nil{
				break
			}
			if !yield(f) {
				return
			}
		}
	}
}

func RootRegularFiles(fsys fs.FS) iter.Seq[fs.File]{
	des,_:=fs.ReadDir(fsys,".")
	return func(yield func(fs.File) bool) {
		for _,de:=range des{
			if !de.Type().IsRegular(){
				continue
			}
			f,err:=fsys.Open(de.Name())
			if err!=nil{
				break
			}
			if !yield(f) {
				return
			}
		}
	}
}


func FileLineScanner(fs iter.Seq[fs.File]) iter.Seq2[string,iter.Seq[string]]{
	return func(yield func(string,iter.Seq[string]) bool) {
		for f:=range fs{
			s,err:=f.Stat()
			if err!=nil{
				continue
			}
			if !yield(s.Name(),LineScanner(f)) {
				return
			}
		}
	}
}



// reads all in one go, sorts it
func GlobScanner(fsys fs.FS, pattern string) iter.Seq2[string,iter.Seq[string]]{
	matches,err:=fs.Glob(fsys,pattern)
	if err!=nil{
		return nil
	}
	return FilesLineScanner(fsys,matches)
}



//func FilesScanner(fsys fs.FS,fs []string) iter.Seq2[string,iter.Seq[string]]{
//	return func(yield func(string,iter.Seq[string]) bool) {
//		for _,e:=range fs{
//			f,err:=fsys.Open(e)
//			if err!=nil{
//				break
//			}
//			if !yield(e,LineScanner(f)) {
//				return
//			}
//		}
//	}
//}


func FilesLineScanner(fsys fs.FS,fs []string) iter.Seq2[string,iter.Seq[string]]{
	return func(yield func(string,iter.Seq[string]) bool) {
		for _,e:=range fs{
			f,err:=fsys.Open(e)
			if err!=nil{
				break
			}
			if !yield(e,LineScanner(f)) {
				return
			}
		}
	}
}

func DirEntryLineScanner(fsys fs.FS,ds []fs.DirEntry) iter.Seq2[string,iter.Seq[string]]{
	return func(yield func(string,iter.Seq[string]) bool) {
		for _,e:=range ds{
			if !e.Type().IsRegular(){
				continue
			}
			f,err:=fsys.Open(e.Name())
			if err!=nil{
				continue
			}
			if !yield(e.Name(),LineScanner(f)) {
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

//fs.ReadDirFile  // n at time




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

