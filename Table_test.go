package sheets

import "fmt"
import "./lists"


func ExampleTable(){
	t:=NewTable(NewRow(1,2,3),NewRow(4,5,6))
	fmt.Printf("%\n",Sheet[int,Row[int]]{Row[Row[int]](lists.List[Row[int]](t))})
	// Output:
	// 1 2 3
	// 4 5 6
}

