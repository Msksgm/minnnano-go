package main

import (
	"fmt"
	"reflect"
)

type Point struct {
	X int
	Y int
}

func main() {
	p := &Point{X: 10, Y: 5}
	fmt.Println(p)
	rv := reflect.ValueOf(p)
	fmt.Printf("rv.Type = %v\n", rv.Type())
	fmt.Printf("rv.Kind = %v\n", rv.Kind())
	fmt.Printf("rv.Interface = %v\n", rv.Interface())

	rv = reflect.ValueOf(p).Elem() // Elemをつけないとエラーになる
	xv := rv.Field(0)              // rv内のX要素を取り出し
	fmt.Printf("xv = %d\n", xv.Int())
	xv.SetInt(100)
	fmt.Printf("xv (after SetInt) = %d\n", xv.Int())

	rv1 := reflect.ValueOf(1)
	// rv2 := reflect.ValueOf("Hello, World")// panic
	// rv3 := reflect.ValueOf([]byte{0xa, 0xd}) // panic
	// rv4 := reflect.ValueOf(make(chan struct{})) // panic

	fmt.Println(rv1.Int())
	// fmt.Println(rv2.Int()) //panic
	// fmt.Println(rv3.Int()) //panic
	// fmt.Println(rv4.Int()) //panic

	rv = reflect.ValueOf(map[string]int{"foo": 1})
	value := rv.MapIndex(reflect.ValueOf("foo"))
	fmt.Println(value)
	rv.SetMapIndex(reflect.ValueOf("foo"), reflect.ValueOf(2))
	// rv.Int() // panic
}
