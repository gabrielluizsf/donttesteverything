package afterbefore

import (
	"testing"

	"github.com/gabrielluizsf/donttesteverything/assert"
)

func TestDoubleNumber(test *testing.T) {
	result := DoubleNumber(2)
	expected := 4
	a := assert.New(test)
	a.Equals(result, expected)
}

func TestIsNullOrEmpty(test *testing.T) {
	result := IsNullOrEmpty("ahfjsaf")
	expected := false
	a := assert.New(test)
	a.Equals(result, expected)
}

func TestAdd(test *testing.T) {
	result := Add(10, 10)
	expected := 20
	a := assert.New(test)
	a.Equals(result, expected)
}
