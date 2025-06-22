package sheets

import "fmt"

func ExampleSheets(){
	ss:=Sheets[Sheet[Row[int],int],int]{}
	ss["test"]=NewSheet(NewRow(1,2,3),NewRow(4,5,6))
	fmt.Printf("%\n",ss["test"])
	// Output:
	// 123
	// 456
}

func ExampleGroupBy(){
	ss:=NewSheet(NewRow(1,2,3),NewRow(4,5,6))
	fmt.Printf("%v",GroupBy(ss,func(_ Row[int])bool{return true} ))
	// Output:
	// 123
	// 456
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
	// 123,,,456
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
	// 123,,,456
}


func ExampleFields() {
	fmt.Printf("%\t", NewHeadedSheet(Fields(struct{Name string;Age uint}{"simon",60})))
	// Output:
	// 1	4	9	16	25	36	49	64	81
}

