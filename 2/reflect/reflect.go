package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Foo struct {
	A int `tag1:"First Tag" tag2:"Second Tag"`
	B string
}

func main() {
	//example1()
	//example2()
	//example3()
	//example4()
	//example5()
}

// Get info
func example1() {
	sl := []int{1, 2, 3}
	greeting := "hello"
	greetingPtr := &greeting
	f := Foo{A: 10, B: "Salutations"}
	fp := &f

	slType := reflect.TypeOf(sl)
	gType := reflect.TypeOf(greeting)
	grpType := reflect.TypeOf(greetingPtr)
	fType := reflect.TypeOf(f)
	fpType := reflect.TypeOf(fp)

	examiner(slType, 0)
	examiner(gType, 0)
	examiner(grpType, 0)
	examiner(fType, 0)
	examiner(fpType, 0)
}

// Update values
func example2() {
	greeting := "hello"

	gVal := reflect.ValueOf(greeting)
	// not a pointer so all we can do is read it
	fmt.Println(gVal.Interface().(string))

	gpVal := reflect.ValueOf(&greeting)
	// it’s a pointer, so we can change it, and it changes the underlying variable
	gpVal.Elem().SetString("goodbye")
	fmt.Println(greeting)

	fType := reflect.TypeOf(Foo{})
	//ValueOf(&Foo{})
	fVal := reflect.New(fType)
	fVal.Elem().Field(0).SetInt(20)
	fVal.Elem().Field(1).SetString("Greetings")
	f2 := fVal.Elem().Interface().(Foo)
	fmt.Printf("%v, %d, %s\n", f2, f2.A, f2.B)
}

// Create map and slice
func example3() {
	// declaring these vars, so we can make a reflect.Type
	intSlice := make([]int, 0)
	mapStringInt := make(map[string]int)

	// here are the reflect.Types
	sliceType := reflect.TypeOf(intSlice)
	mapType := reflect.TypeOf(mapStringInt)

	// and here are the new values that we are making
	intSliceReflect := reflect.MakeSlice(sliceType, 0, 0)
	mapReflect := reflect.MakeMap(mapType)

	// and here we are using them
	v := 10
	rv := reflect.ValueOf(v)
	intSliceReflect = reflect.Append(intSliceReflect, rv)
	intSlice2 := intSliceReflect.Interface().([]int)
	fmt.Println(intSlice2)

	k := "hello"
	rk := reflect.ValueOf(k)
	mapReflect.SetMapIndex(rk, rv)
	mapStringInt2 := mapReflect.Interface().(map[string]int)
	fmt.Println(mapStringInt2)
}

// Create a struct
func example4() {
	s := makeStruct(0, "", []int{})
	// this returned a pointer to a struct with 3 fields:
	// an int, a string, and a slice of ints
	// but you can’t actually use any of these fields
	// directly in the code; you have to reflect them
	sr := reflect.ValueOf(s)

	// getting and setting the int field
	fmt.Println(sr.Elem().Field(0).Interface().(int))
	sr.Elem().Field(0).SetInt(20)
	fmt.Println(sr.Elem().Field(0).Interface().(int))

	// getting and setting the string field
	fmt.Println(sr.Elem().Field(1).Interface().(string))
	sr.Elem().Field(1).SetString("reflect me")
	fmt.Println(sr.Elem().Field(1).Interface().(string))

	// getting and setting the []int field
	fmt.Println(sr.Elem().Field(2).Interface().([]int))
	v := []int{1, 2, 3}
	rv := reflect.ValueOf(v)
	sr.Elem().Field(2).Set(rv)
	fmt.Println(sr.Elem().Field(2).Interface().([]int))
}

// reflect.DeepEqual
func example5() {
	f1 := Foo{}
	f2 := Foo{A: 10}
	f3 := Foo{A: 10, B: "test"}
	f4 := Foo{A: 20, B: "test"}
	f5 := Foo{A: 20, B: "test"}

	pf1 := &f1
	pf2 := &f2
	pf3 := &f3
	pf4 := &f4
	pf5 := &f5

	fmt.Println(reflect.DeepEqual(f1, f2))
	fmt.Println(reflect.DeepEqual(f2, f3))
	fmt.Println(reflect.DeepEqual(f3, f4))
	fmt.Println(reflect.DeepEqual(f4, f5))

	fmt.Println()

	fmt.Println(reflect.DeepEqual(pf1, pf2))
	fmt.Println(reflect.DeepEqual(pf2, pf3))
	fmt.Println(reflect.DeepEqual(pf3, pf4))
	fmt.Println(reflect.DeepEqual(pf4, pf5))
}

func examiner(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth), "Type is", t.Name(), "and kind is", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth+1), "Contained type:")
		examiner(t.Elem(), depth+1)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "name is", f.Name, "type is", f.Type.Name(), "and kind is", f.Type.Kind())
			if f.Tag != "" {
				fmt.Println(strings.Repeat("\t", depth+2), "Tag is", f.Tag)
				fmt.Println(strings.Repeat("\t", depth+2), "tag1 is", f.Tag.Get("tag1"), "tag2 is", f.Tag.Get("tag2"))
			}
		}
	}
}

func makeStruct(vals ...interface{}) interface{} {
	var sfs []reflect.StructField
	for i, v := range vals {
		t := reflect.TypeOf(v)
		sf := reflect.StructField{
			Name: fmt.Sprintf("F%d", i+1),
			Type: t,
		}
		sfs = append(sfs, sf)
	}
	st := reflect.StructOf(sfs)
	so := reflect.New(st)
	return so.Interface()
}
