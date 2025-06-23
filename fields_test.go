package sheets

import "fmt"

func ExampleFields() {
	fmt.Print(NewHeadedSheet(Fields(struct{Name string;Age uint}{"simon",60})))
	// Output:
	// 
}

func ExampleFieldsTags() {
	fmt.Print(NewHeadedSheet(FieldsTags(FieldsStructTags(FieldsValues(struct{Name string `type:"Given"`;Age uint `unit:"Years"`}{"simon",60})),"unit")))
	// Output:
	// 
}





