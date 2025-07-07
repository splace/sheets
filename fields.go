package sheets

import "reflect"
import "iter"
//import "fmt"
import "./lists"

func ValueList[E any, Slice ~[]T, T any](s Slice) lists.List[Row[E]]{
	return func(yield func(Row[E])bool){
		for _,t:=range s {
			if !yield(FieldValues(Fields[E](t))) {
				return
			}
		}
	}	
}

func TypedSheet[T any,U any](t Sheet[U,Row[U]]) Sheet[T,Row[T]]{
	return Sheet[T,Row[T]]{
		func(yield func(Row[T])bool){
			for r:=range t.Row {
				if !yield(FieldValues(Fields[T](r))) {
					return
				}
			}
		},
	}
}


// NOTICE: if type isn't whats needed its skipped.
// so index of value will generally be changed
// Note: can be used to select a typed 
func Fields[T any](a any) iter.Seq2[string,T] {
	t := reflect.TypeOf(a)
	if t.Kind()!=reflect.Struct{
		return nil
	}
	v := reflect.ValueOf(a)
	return func(yield func(string,T)bool){
		for i := 0; i < v.NumField(); i++ {
			if tt,is:=v.Field(i).Interface().(T);is{
				if !yield(t.Field(i).Name,tt)  {
					return
				}
			}
		}
	}
}

func FieldNames(fs iter.Seq2[string,any]) Row[string]{
	return func(yield func(string)bool){
		for n,_:=range fs{
			if !yield(n) {
				return
			}
		}
	}
}

func FieldValues[T any](fs iter.Seq2[string,T]) Row[T]{
	return func(yield func(T)bool){
		for _,v:=range fs{
			if !yield(v) {
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

func FieldsStructure(a any) iter.Seq[reflect.StructField] {
	t := reflect.TypeOf(a)
	if t.Kind()!=reflect.Struct{
		return nil
	}
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

