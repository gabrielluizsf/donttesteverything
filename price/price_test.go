package price

import (
	"testing"

	"github.com/gabrielluizsf/donttesteverything/assert"
)

func TestValue(test *testing.T) {
	result := Value("s")
	expected := 90.0
	a := assert.New(test)
	a.Equals(result, expected)
	result = Value("n")
	expected = 0
	a.Equals(result, expected)
	result = Value("no")
	expected = 0
	a.Equals(result, expected)
	result = Value("y")
	expected = 0
	a.Equals(result,expected)
}
