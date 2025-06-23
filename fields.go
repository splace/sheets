package sheets

import "reflect"
import "iter"
//import "fmt"

func Fields(a any) iter.Seq2[string,any] {
	v := reflect.ValueOf(a)
	t := v.Type()
	return func(yield func(string,any)bool){
		for i := 0; i < v.NumField(); i++ {
			if !yield(t.Field(i).Name, v.Field(i).Interface()) {
				return
			}
		}
	}
}

func FieldsStructTags(rvs iter.Seq[reflect.StructField]) iter.Seq2[string,reflect.StructTag] {
	return func(yield func(string,reflect.StructTag)bool){
		for v:=range rvs{
			if !yield(v.Name, v.Tag) {
				return
			}
		}
	}
}

func FieldsTags(rvs iter.Seq2[string,reflect.StructTag], tName string) iter.Seq2[string,string] {
	return func(yield func(string,string)bool){
		for name,stag:=range rvs{
			if !yield(name, stag.Get(tName)) {
				return
			}
		}
	}
}

func FieldsValues(a any) iter.Seq[reflect.StructField] {
	t := reflect.TypeOf(a)
	return func(yield func(reflect.StructField)bool){
		for i := 0; i < t.NumField(); i++ {
			if !yield(t.Field(i)) {
				return
			}
		}
	}
}


// Tags(reflect.TypeOf(sSI{})
//func Tags(t reflect.Type,l string) iter.Seq[string]{
//	if l==""|| l=="name" {
//		return func(yield func(string) bool) {
//			for _,f:=range reflect.VisibleFields(t){
//				if !yield(fmt.Sprint(f.Name)){
//					return
//				}
//			}
//		}
//	}
//	return func(yield func(string) bool) {
//		for _,f:=range reflect.VisibleFields(t){
//			if !yield(fmt.Sprint(f.Tag.Get(l))){
//				return
//			}
//		}
//	}
//}

