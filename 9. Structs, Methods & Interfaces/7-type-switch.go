package main

import "fmt"

type myOwnType struct {
	//created my own type
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case myOwnType:
		fmt.Printf("This is my own created type")
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	ownType := myOwnType{}
	do(21)
	do("hello")
	do(true)
	do(ownType)
}
