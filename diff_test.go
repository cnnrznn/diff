package diff

import (
	"fmt"
	"testing"
)

func TestSameString(t *testing.T) {
	s1 := "hello"
	s2 := "hello"

	Diff(s1, s2)
}

func TestDifferentString(t *testing.T) {
	s1 := "hello"
	s2 := "world"

	seq := Diff(s1, s2)
	fmt.Println(seq.Steps)
}

func TestMiddleEdit(t *testing.T) {
	s1 := "something"
	s2 := "someINTHEMIDDLEthing"

	seq := Diff(s1, s2)
	fmt.Println(seq.Steps)
}

func TestDocs(t *testing.T) {
	s1 := `This is a document discussing many important topics.

	Such as how to bake cakes; it's easy!
	
	1. step one
	2. step two
	3. step three	
	`
	s2 := `This is a document discussing many important topics.

	Such as how to make cakes; it's easy!
	2. step two
	3. step three	
	`

	seq := Diff(s1, s2)
	fmt.Println(seq.Steps)
}
