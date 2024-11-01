package main

import "fmt"

func switchExamples() {
	// switchWhichCheckValue()
	// switchWhichCheckTrue()
	// unnecessaryInitiator()
	resolveDataType()
}

func switchWhichCheckValue() {
	value := 423
	switch value {
	case 1:
		fmt.Println("contain 1")
	case 2:
		fmt.Println("contain 2")
	case 3:
		fmt.Println("contain 3")
	default:
		fmt.Println("more than 3")
	}
}

func switchWhichCheckTrue() {
	value := 5
	switch {
	case value == 1:
		fmt.Println("contain 1")
	case value == 2:
		fmt.Println("contain 2")
	case value == 3, value == 4:
		fmt.Println("contain 3 or 4")
	default:
		fmt.Println("more than 4")
	}
}

func unnecessaryInitiator() {
	switch num := 6; num%2 == 0 {
	case true:
		fmt.Println("num is even")
	case false:
		fmt.Println("num is odd")
	}
}

func resolveDataType() {
	var data interface{}
	data = 333.2222
	fmt.Printf("type %T\n", data)

	switch mytype := data.(type) {
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("boolean")
	case float64:
		fmt.Println("float64 type")
	case float32:
		fmt.Println("float32 type")
	case int:
		fmt.Println("int")
	default:
		fmt.Printf("%T", mytype)
	}
}
