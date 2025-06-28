package sheets

import "fmt"

type Person struct{
	Name string `type:"Given"`
	Age uint `unit:"Years"`
}


func ExampleFields() {
	fmt.Print(NewHeadedSheet(Fields(Person{"simon",60})))
	// Output:
	// 
}

func ExampleFieldsTags() {
	fmt.Print(NewHeadedSheet(
		FieldsTags(
			FieldsStructTags(
				FieldsValues(Person{"simon",60}),
			),
		"unit",
		),
	))
	// Output:
	// 
}





