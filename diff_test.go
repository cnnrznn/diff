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
