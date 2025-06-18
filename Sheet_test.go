package sheets

import "fmt"

func ExampleSheet(){
	t:=NewSheet(NewRow(1,2,3),NewRow(4,5,6))
	fmt.Printf("%v\n",t)
	t=NewSheet(NewRow[int](1,2,3),NewRow[int](4,5,6))
	fmt.Printf("%v\n",t)
	t2:=NewSheet(NewRow[int](1,2,3),NewRow[int](4,5,6))
	fmt.Printf("%v\n",t2)
	t3:=NewSheet(NewRow[any](1,2,"hi"),NewRow[any](1,2,"hi"))
	fmt.Printf("%v\n",t3)
	// Output:
	//
}

func ExampleCompareSheets(){
	t1:=NewSheet(NewRow(1,2,3),NewRow(4,5,6))
	t2:=NewSheet(NewRow(1,2,3),NewRow(4,5,6))
	fmt.Println(CompareSheets(t1,t2))
	t3:=NewSheet(NewRow(1,2,3),NewRow(4,5,7))
	fmt.Println(CompareSheets(t1,t3))
	// Output:
	// true
	// false
}


func ExampleHeadedSheet(){
	ht:=HeadedSheet[int]{
		NewRow("age","height","weight"),
		NewSheet[Row[int]](
			NewRow(1,2,3),
			NewRow(4,5,6),
			NewRow(7,8,9),
		),
	}
	fmt.Printf("%\t\n",ht)
	// Output:
	//
}

//func ExampleHeadedSheet(){
//	ht:=HeadedSheet[int]{
//		NewRow("age","height","weight"),
//		NewSheet[Row[int]](
//			NewRow(1,2,3),
//			NewRow(4,5,6),
//			NewRow(7,8,9),
//		),
//	}
//	fmt.Printf("%\t\n",ht)
//	// Output:
//	//
//}

