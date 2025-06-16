package table

import (
	"fmt"
	"os"
)

import . "../../sheet"


func ExampleScanner(){
	f,_:=os.Open("./testing/test.csv")
	r:=Row[string](LineScanner(f))
	fmt.Printf("%\n",r)
	// Output:
	// 1,2,3,4
	// 1,2,3,4
	// 1,2,3,4
	// 1,2,3,4
}

