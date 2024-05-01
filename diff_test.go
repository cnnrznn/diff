package diff

import (
	"testing"
)

func TestSameString(t *testing.T) {
	s1 := "hello"
	s2 := "hello"

	Diff(s1, s2)
}

func TestDifferentString(t *testing.T) {

}

func TestMiddleEdit(t *testing.T) {

}
