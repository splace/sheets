package table

import (
	"fmt"
	"os"
//	"iter"
)

import . "../../sheet"


func ExampleLineScanner(){
	f,_:=os.Open("./testing/test.csv") 
	r:=Row[string](LineScanner(f))
	fmt.Printf("%\n",r)
	// Output:
	// 1,2,3,4
	// 1,2,3,4
	// 1,2,3,4
	// 1,2,3,4
}


func ExampleFolderScanner(){
	for n,rs:=range FolderScanner("./testing/people.fsv"){
		fmt.Printf("%q\t",n)
		fmt.Printf("%\n",Row[string](rs))
	}
	// Output:
	// 
}

func ExampleGlobScanner(){
	for n,rs:=range GlobScanner(os.DirFS("./testing/people.fsv"),"*.tsv"){
		fmt.Printf("%q\t",n)
		fmt.Printf("%\n",Row[string](rs))
	}
	// Output:
	// 
}

func ExampleRootRegularFiles(){
	fsys:=os.DirFS("./testing/people.fsv")
	for n,ls:=range FileLineScanner(RootRegularFiles(fsys)){
		fmt.Printf("%q\n",n)
		for l:=range ls{
			fmt.Printf("\t%s\n",l)
		}
	}
	// Output:
	// 
}


