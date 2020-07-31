package main

import "fmt"

// See https://github.com/rpagliuca/golang-diamond-problem
// for demonstration of the problem we're solving here

func main() {
	t := TypeWithoutAmbiguity{}
	t.Method()
}

type Type1 struct{}

type Type2 struct{}

type ComposedType struct {
	Type1
	Type2
}

func (Type1) Method() {
	fmt.Println("Called Type1.Method")
}

func (Type2) Method() {
	fmt.Println("Called Type2.Method")
}

// Below is the fix
type TypeWithoutAmbiguity struct {
	ComposedType
}

func (t TypeWithoutAmbiguity) Method() {
	// Here we remove the ambiguity, chosing
	// which method to call (Type1.Method or Type2.Method)
	t.Type2.Method()
}
