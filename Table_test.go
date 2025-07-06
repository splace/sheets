package sheets

import "fmt"

func ExampleTable(){
	t:=NewTable(NewRow(1,2,3),NewRow(4,5,6))
	fmt.Printf("%v\n",t)
	// Output:
	// 1 2 3
	// 4 5 6
}

