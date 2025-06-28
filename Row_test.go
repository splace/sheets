package sheets

import "fmt"
import "slices"

func ExampleRow_At(){
	fmt.Println(NewRow(0,1,2,3,4,5,6,7,8,9).At(6))
	// Output:
	// 6
}

func ExampleRow(){
	fmt.Println(NewRow(0,1,2,3,4,5,6,7,8,9))
	// Output:
	// 0 1 2 3 4 5 6 7 8 9
}

func ExampleRow_Scan(){
	var r = new(Row[uint8])
	fmt.Println(fmt.Sscanf("0 1 2 3 4 5 6 7 8 9\n","%v",r))
	fmt.Printf("%v\n",r)
	// Output:
	// 1 <nil>
	// 0 1 2 3 4 5 6 7 8 9
}



func ExampleRow_Sample(){
	fmt.Println(NewRow(0,1,2,3,4,5,6,7,8,9).Sample([2]uint{4,2}))
	// Output:
	// 4 6 8
}

func ExampleNewRow(){
	ns:=[]float32{0,1,2,3,4,5,6,7,8,9}
	n:=NewRow(ns...)
	fmt.Println(n)
	ns[6]=0
	slices.Sort(ns)
	fmt.Println(n)
	// Output:
	// 0 1 2 3 4 5 6 7 8 9
	// 0 0 1 2 3 4 5 7 8 9
}

func ExampleRow_Cache(){
	ns:=[]float32{0,1,2,3,4,5,6,7,8,9}
	n:=NewRow(ns...)
	cached:=n.Cache()
	fmt.Println(n,cached)
	ns[6]=0
	fmt.Println(n,cached)
	// Output:
	// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9
	// 0 1 2 3 4 5 0 7 8 9 0 1 2 3 4 5 6 7 8 9
}


func ExampleReverse(){
	n:=NewRow(0,1,2,3,4,5,6,7,8,9)
	fmt.Println(n.Sub([2]uint{4,4}))
	fmt.Println(n.Sub([2]uint{4,4}).Reverse())
	// Output:
	// 4 5 6 7
	// 7 6 5 4
}

func ExampleSorted(){
	n:=NewRow(0,1,2,3,5,4,6,7,8,9)
	fmt.Println(Sorted(n))
	// Output:
	// 0 1 2 3 4 5 6 7 8 9
}

func ExampleRow_Items(){
	fmt.Println(NewRow(0,1,2,3,4,5,6,7,8,9).Items(2,7))
	// Output:
	// 1 6
}

func ExampleRow_Items_missing(){
	fmt.Println(NewRow(0,1,2,3,4,5,6,7,8,9).Items(2,27)==nil)
	// Output:
	// true
}


func ExampleRow_Select(){
	fmt.Println(NewRow(0,1,2,3,4,5,6,7,8,9).Select(2,7,2,7))
	// Output:
	// 1 6 1 6
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

