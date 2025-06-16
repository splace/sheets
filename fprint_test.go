package sheets

import "fmt"
import "slices"


func ExampleSprint(){
//	fmt.Printf("[%q]\n",NewRow[any](1,2,"hi"))
	fmt.Println("["+fmt.Sprint(NewRow[any](1,2,"hi").Sprintf("%10v","|%10v",nil))+"]",
//		NewRow(nil,format("\t%d\t"),format("%q")),
	)
	// Output:
	// [         1|         2|        hi]
}

func ExampleSprintf(){
	ns:=[]float32{0,1,2,3}
	n:=NewRow(ns...)
	fmt.Print("║")
	fmt.Print(n.Sprintf("%10.4fk","│%10.4fk",nil))
	fmt.Println("║")
	ns[3]=0
	slices.Sort(ns)
	fmt.Print("║")
	fmt.Print(n.Sprintf("%10.4fk","│%10.4fk",nil))
	fmt.Println("║")
	// Output:
	// ║    0.0000k│    1.0000k│    2.0000k│    3.0000k║
	// ║    0.0000k│    0.0000k│    1.0000k│    2.0000k║

}




