package lists

import "fmt"
//import . "golang.org/x/exp/constraints"
//import "math"
//import "testing"
//import "iter"


func ExampleScan() {
	var ls List[string]
	_,err:=fmt.Sscan("2 3 Hi",&ls)
	if err!=nil{
		fmt.Println(err)		
	}
	fmt.Println(ls)
	var li List[int]
	_,err=fmt.Sscan("2 3 Hi",&li)
	if err!=nil{
		fmt.Println(err)		
	}
	fmt.Println(li)
	var lc List[complex64]
	_,err=fmt.Sscan("2+0i 3-6i",&lc)
	if err!=nil{
		fmt.Println(err)		
	}
	fmt.Println(lc)
	// Output:
	// 2 3 Hi
	// 2 3
	// (2+0i) (3-6i)

}

