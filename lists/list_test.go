package lists

import "fmt"
//import . "golang.org/x/exp/constraints"
//import "math"
import . "../sequences"
import "slices"
//import "log"
import "strings"

func ExampleList(){
	fmt.Println(List[int](slices.Values([]int{1,2,3})))
	lc:=List[List[int]](
		slices.Values([]List[int]{
			List[int](Limit(Odds[int](),3)),
			List[int](Limit(Evens[int](),3)),
		},
	))
	fmt.Println(lc)
	// Output:
	// 1 2 3
}

func ExampleStrings() {
	var ss Strings
	fmt.Println(fmt.Sscan("2 3 Hi",&ss))
	fmt.Println(ss)
	l:=To[int](ss...)
	fmt.Println(l)
	// Output:
	// 1 <nil>
	// [2 3 Hi]
	// 2 3
}



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
	_,err=fmt.Sscanf("2+0i,3-6i","%,",&lc)
	if err!=nil{
		fmt.Println(err)	
	}
	fmt.Println(lc)
	_,err=fmt.Sscanf("2+0i	3-6i","%	",&lc)
	if err!=nil{
		fmt.Println(err)	
	}
	fmt.Println(lc)
	_,err=fmt.Sscanf("2+0i\n3-6i","%\n",&lc)
	if err!=nil{
		fmt.Println(err)	
	}
	fmt.Println(lc)
	s:=strings.NewReader("2+0i\n3-6i\n2+0i\n3-6i")
	_,err=fmt.Fscanln(s,&lc)
	_,err=fmt.Fscanln(s,&lc)
	_,err=fmt.Fscanln(s,&lc)
	_,err=fmt.Fscanln(s,&lc)
	if err!=nil{
		fmt.Println(err)	
	}
	fmt.Println(lc)
	// Output:
	// 2 3 Hi
	// 2 3
	// (2+0i) (3-6i)
	// (2+0i) (3-6i)
	// (2+0i) (3-6i)
	// (3-6i)
}

func ExampleScan_array() {
	s:=strings.NewReader("2+0i 3-6i\n3+0i 4-6i")
	var lc List[List[complex64]]
	fmt.Println(fmt.Fscanln(s,&lc))
	fmt.Println(lc)
	fmt.Println(fmt.Fscanln(s,&lc))
	fmt.Println(lc)
	// Output:
	// (2+0i) (3-6i)

}



//func ExampleScan_b() {
//	rc:=make(chan(List[string]))
//	go func(){
//		var ls List[string]
//		fmt.Sscan("2 3 Hi",&ls)
//		log.Printf("%p %[1]T",ls)
//		log.Print(ls)
//		rc <- ls
//	}()
////	fmt.Println(ls)
////	
////	var li List[int]
////	_,err=fmt.Sscan("2 3 Hi",&li)
////	if err!=nil{
////		fmt.Println(err)		
////	}
////	fmt.Println(li)
////	var lc List[complex64]
////	_,err=fmt.Sscan("2+0i 3-6i",&lc)
////	if err!=nil{
////		fmt.Println(err)		
////	}
//	fmt.Printf("%T\n",rc )
//	fmt.Print(<-rc )
//	// Output:
//	// 2 3 Hi
//	// 2 3
//	// (2+0i) (3-6i)

//}


