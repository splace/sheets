package sheets

import "reflect"
import "iter"


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


