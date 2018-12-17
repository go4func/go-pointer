package main

import (
	"fmt"
	"reflect"
)

type testStruct struct {
	Name string
	Nest *testNested
}

type testNested struct {
	Name string
}

func (t *testStruct) Clone() testStruct {
	ret := testStruct{}
	rFields := reflect.ValueOf(&ret).Elem()
	tFields := reflect.ValueOf(t).Elem()
	for i := 0; i < tFields.NumField(); i++ {
		f := tFields.Field(i)
		r := rFields.Field(i)
		switch f.Kind() {
		case reflect.Ptr:
			// r.Set(reflect.New(f.Type().Elem()))
			r.Set(f)
		default:
			r.Set(f)
		}
	}

	return ret
}

func main() {
	t := testStruct{
		Name: "test struct",
		Nest: &testNested{
			Name: "nested struct",
		},
	}
	t2 := t.Clone()
	fmt.Println(t2)

	// fmt.Println(t.Nest.Name)

	// t2 := t
	// t2.Nest = new(testNested)
	// *t2.Nest = *t.Nest
	// t2.Nest.Name = "new name"

	// fmt.Println(t.Nest.Name)
}
