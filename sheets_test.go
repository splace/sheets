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

