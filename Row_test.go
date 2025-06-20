package sheets

import "fmt"
import "slices"

func ExampleRow_At(){
	fmt.Println(NewRow(0,1,2,3,4,5,6,7,8,9).At(6))
	// Output:
	// 6
}

func ExampleRow_Sample(){
	fmt.Println(NewRow(0,1,2,3,4,5,6,7,8,9).Sample([2]uint{4,2}))
	// Output:
	// 468
}

func ExampleNewRow(){
	ns:=[]float32{0,1,2,3,4,5,6,7,8,9}
	n:=NewRow(ns...)
	fmt.Println(n)
	ns[6]=0
	slices.Sort(ns)
	fmt.Println(n)
	// Output:
	// 0123456789
	// 0012345789
}

func ExampleReverse(){
	n:=NewRow(0,1,2,3,4,5,6,7,8,9)
	fmt.Println(n.Sub([2]uint{4,4}))
	fmt.Println(n.Sub([2]uint{4,4}).Reverse())
	// Output:
	// 4567
	// 7654
}

func ExampleSorted(){
	n:=NewRow(0,1,2,3,5,4,6,7,8,9)
	fmt.Println(Sorted(n))
	// Output:
	// 0123456789
}

func ExampleRow_Items(){
	fmt.Println(NewRow(0,1,2,3,4,5,6,7,8,9).Items(2,7))
	// Output:
	// 16
}

func ExampleRow_Select(){
	fmt.Println(NewRow(0,1,2,3,4,5,6,7,8,9).Select(2,7,2,7))
	// Output:
	// 1616
}


func ExampleCompare(){
	s1,s2:=[]int{0,1,2,3,4,5,6,7,8,9},[]int{0,1,2,3,4,5,6,7,8,9}
	r1,r2:=NewRow(s1...),NewRow(s2...)
	fmt.Println(CompareRows(r1,r2))
	s1[5]+=1
	fmt.Println(CompareRows(r1,r2))
	s1[5]=5
	fmt.Println(CompareRows(r1,r2))
	// Output:
	// true
	// false
	// true
}

