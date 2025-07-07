package sheets

import "fmt"
import "./lists"
//import "./sequences"



type Person struct{
	Name string `type:"Given"`
	Age uint `unit:"Years"`
}

func ExampleFields() {
	fmt.Println(NewHeadedSheet(Fields[any](Person{"sixty",60})))
	// Output:
	// 'Name' 'Age'
	//'sixty' '60'
}

func ExampleFields_typed() {
	fmt.Println(NewHeadedSheet(Fields[uint](Person{"sixty",60})))
	// Output:
	//'Age'
	//60
}

func ExampleTypedSheet() {
	ps:=[]Person{
		Person{"sixty",60},
		Person{"sixty",61},
		Person{"sixty",62},
		Person{"sixty",63},
		Person{"sixty+",60},
	}
	t:=Sheet[any,Row[any]]{Row[Row[any]](ValueList[any](ps))}
	mt:=TypedSheet[int](t)
	fmt.Printf("%\n\n",t)
	fmt.Printf("%\n\n",mt)
	fmt.Printf("%\n\n",SelectRowsFrom(t,mt,0,60))
	// Output:
	//
}

//func ExampleSelectRowsFrom_fields() {
//	ps:=[]Person{
//		Person{"sixty",60},
//		Person{"sixty",61},
//		Person{"sixty",62},
//		Person{"sixty",63},
//		Person{"sixty+",60},
//	}
//	t:=Sheet[Row[any],any]{Row[Row[any]](ValueList(ps))}
//	mt:=Sheet[Row[uint],uint]{Row[Row[uint]](FieldValues(Fields[uint](ps)))}
////	mt:=Sheet[Row[uint],uint]{Row[Row[uint]](sequences.Limit(sequences.Repeat(FieldValues(Fields[uint](ps[0]))),len(ps)))}
//	ValueList()
//	fmt.Printf("%\n\n",t)
//	fmt.Printf("%\n\n",mt)
//	fmt.Printf("%\n\n",SelectRowsFrom(t,mt,0,60))
//	// Output:
//	//
//}


func ExampleFieldsTags() {
	fmt.Print(NewHeadedSheet(
		FieldsTags(
			FieldsStructTags(
				FieldsStructure(Person{"sixty",60}),
			),
		"unit",
		),
	))
	// Output:
	// 
}


func ExampleValueList() {
	ps:=[]Person{
		Person{"sixty",60},
		Person{"sixty",60},
		Person{"sixty",60},
		Person{"sixty",60},
		Person{"sixty",60},
	}
	t:=Table[any,Row[any]](ValueList[any](ps))
	fmt.Printf("%\n",Sheet[any,Row[any]]{Row[Row[any]](lists.List[Row[any]](t))})
	// Output:
	// 
}

func ExampleValueList_b() {
	ps:=[]Person{
		Person{"sixty",60},
		Person{"sixty",60},
		Person{"sixty",60},
		Person{"sixty",60},
		Person{"sixty",60},
	}
	t:=Table[any,Row[any]](ValueList[any](ps))
	fmt.Printf("%\n",Sheet[any,Row[any]]{Row[Row[any]](lists.List[Row[any]](t))})
	// Output:
	// 
}

func ExampleValueList_typed() {
	ps:=[]Person{
		Person{"sixty",60},
		Person{"sixty",60},
		Person{"sixty",60},
		Person{"sixty",60},
		Person{"sixty",60},
	}
	t:=Table[int,Row[int]](ValueList[int](ps))
	fmt.Printf("%\n",Sheet[int,Row[int]]{Row[Row[int]](lists.List[Row[int]](t))})
	// Output:
	// 
}







