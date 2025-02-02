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

	a := true
	b := true
	if true { // want "only one cuddle assignment allowed before if statement"
		fmt.Println("didn't use cuddled variable")
	}

	_ = a
	_ = b
}

func For() {
	x := true
	for {
		fmt.Println("didn't use cuddled variable")
	}

	_ = x

	a := true
	b := true
	for { // want "only one cuddle assignment allowed before for statement"
		fmt.Println("didn't use cuddled variable")
	}

	_ = a
	_ = b
}

func Range() {
	x := true
	for i := range make([]int, 10) {
		fmt.Println("...")
	}

	_ = x

	a := true
	b := true
	for i := range make([]int, 10) { // want "only one cuddle assignment allowed before range statement"
		fmt.Println("didn't use cuddled variable")
	}

	_ = a
	_ = b
}

func Switch() {
	x := true
	switch "a" {
	case "a":
		return 1
	default:
		return 2
	}

	_ = x

	a := true
	b := true
	switch "a" { // want "only one cuddle assignment allowed before switch statement"
	case "a":
		return 1
	default:
		return 2
	}

	_ = a
	_ = b
}

func TypeSwitch() {
	var a any

	x := true
	switch a.(type) {
	case int:
		return 1
	default:
		return 2
	}

	_ = x
	_ = a

	b := true
	c := true
	switch a.(type) { // want "only one cuddle assignment allowed before type switch statement"
	case "a":
		return 1
	default:
		return 2
	}

	_ = b
	_ = c
}

func Defer() {
	x := true
	defer func() {
		fmt.Println("a")
	}()

	_ = x

	a := true
	b := true
	defer func() { // want "only one cuddle assignment allowed before defer statement"
		fmt.Println("a")
	}()

	_ = a
	_ = b
}
