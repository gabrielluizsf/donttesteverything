package calc

import (
	"testing"

	"github.com/gabrielluizsf/donttesteverything/assert"
)

func TestAdd(test *testing.T) {
	result := Execute("add", 10.0, 10.0)
	expected := 20
	a := assert.New(test)
	a.Equals(result, expected)
}

func TestSub(test *testing.T) {
	result := Execute("sub", 100.0, 50.0)
	expected := 50
	a := assert.New(test)
	a.Equals(result, expected)
}

func TestMult(test *testing.T) {
	result := Execute("mult", 0, 1)
	expected := 0
	a := assert.New(test)
	a.Equals(result, expected)
	result = Execute("mult", 10, 2)
	expected = 20
	a.Equals(result, expected)
}

func TestDiv(test *testing.T) {
	result := Execute("div", 0, 1)
	expected := 0
	a := assert.New(test)
	a.Equals(result, expected)
	result = Execute("div", 10, 2)
	expected = 5
	a.Equals(result, expected)
}
