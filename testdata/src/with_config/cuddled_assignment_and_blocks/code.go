package pkg

import (
	"fmt"
)

func If() {
	x := true
	if true {
		fmt.Println("didn't use cuddled variable")
	}

	_ = x
}

func For() {
	x := true
	for {
		fmt.Println("didn't use cuddled variable")
	}

	_ = x
}

func Range() {
	x := true
	for i := range make([]int, 10) {
		fmt.Println("...")
	}

	_ = x
}

func Switch() {
	x := true
	switch "a" {
	case "b":
		return 1
	case "c":
		return 2
	default:
		return 3
	}

	_ = x
}

func TypeSwitch() {
	var a any

	x := true
	switch a.(type) {
	case int:
		return 1
	case uint:
		return 2
	default:
		return 3
	}

	_ = x
	_ = a
}

func Defer() {
	x := true
	defer func() {
		fmt.Println("a")
	}()

	_ = x
}
