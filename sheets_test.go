package sheets

import "fmt"

func ExampleSheets(){
	ss:=Sheets[int,Sheet[int,Row[int]]]{}
	ss["test"]=NewSheet(NewRow(1,2,3),NewRow(4,5,6))
	fmt.Printf("%\n",ss["test"])
	// Output:
	// 1 2 3
	// 4 5 6
}

func ExampleGroupBy(){
	ss:=NewSheet(NewRow(1,2,3),NewRow(4,5,6))
	fmt.Printf("%v",GroupBy(ss,func(_ Row[int])bool{return true} ))
	// Output:
	// 1 2 3
	// 4 5 6
}

func ExampleGroupBy2(){
	ss:=NewSheet(NewRow(1,2,3),NewRow(4,5,6))
	fmt.Printf("%,\n",GroupBy2(ss,
		NewRow(
			func(r Row[int])bool{return r.At(0)==1},
			func(r Row[int])bool{return r.At(0)==2},
			func(r Row[int])bool{return r.At(0)==3},
			func(r Row[int])bool{return r.At(0)==4},
		),
	))
	// Output:
	// 1 2 3,,,4 5 6
}

func ExampleGroupBy2_concat(){
	ss:=NewSheet(NewRow(1,2,3),NewRow(4,5,6))
	fmt.Printf("%,\n",GroupBy2(ss,
		NewRow(
			func(r Row[int])bool{return r.At(0)==1},
			func(r Row[int])bool{return r.At(0)==2},
			func(r Row[int])bool{return r.At(0)==3},
			func(r Row[int])bool{return r.At(0)==4},
		),
	))
	// Output:
	// 1 2 3,,,4 5 6
}



