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


func ExampleSheet_Column(){
	fmt.Printf("%\t\n",NewSheet(NewRow(1,2,3),NewRow(4,5,6)).Column(2))
	// Output:
	// 2	5
}

func ExampleSheet_SelectColumns(){
	fmt.Printf("%\n\n",NewSheet(NewRow(1,2,3),NewRow(4,5,6)).SelectColumns(2,3))
	// Output:
	// 2 3
	// 5 6
}


func ExampleSelectRows(){
	s:=NewSheet(
		NewRow(1,2,3),
		NewRow(4,5,6),
		NewRow(3,2,1),
	)
	fmt.Printf("%\n\n",SelectRows(s,1,2))
	fmt.Printf("%\n\n",SelectRows(s,0,1))
	// Output:
	// 1 2 3
	// 3 2 1
	// 1 2 3
}

func ExampleSelectRowsFrom(){
	s:=NewSheet(
		NewRow(1,2,3),
		NewRow(4,5,6),
		NewRow(3,2,1),
	)
	fmt.Printf("%\n\n",SelectRowsFrom(s,s.SelectColumns(2),0,2))
	fmt.Printf("%\n\n",SelectRowsFrom(s,s,0,1))
	// Output:
	// 1 2 3
	// 3 2 1
	// 1 2 3
}

func ExampleSelectMatchedRows(){
	s:=NewSheet(
		NewRow(1,2,3),
		NewRow(3,2,1),
		NewRow(4,5,6),
		NewRow(3,2,1),
	)
	
	fmt.Printf("%\n\n",SelectMatchedRows(s,NewRow(3,2,1)))
	// Output:
	// 3 2 1
	// 3 2 1
}

func ExampleSelectMatchedRowsFrom(){
	s:=NewSheet(
		NewRow(1,2,3),
		NewRow(3,2,1),
		NewRow(4,5,6),
		NewRow(3,2,5),
	)
	
	fmt.Printf("%\n\n",SelectMatchedRowsFrom(s,s.SelectColumns(2,1),NewRow(2,3)))
	// Output:
	// 3 2 1
	// 3 2 5
}


func ExampleHeadedSheet(){
	ht:=HeadedSheet[int]{
		NewRow("age","height","weight"),
		NewSheet[int](
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

