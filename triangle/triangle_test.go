package triangle

import (
	"testing"

	"github.com/gabrielluizsf/donttesteverything/assert"
)

func TestCheckIsTriangle(test *testing.T) {
	result := CheckIsTriangle(10, 12, 13)
	expected := true
	a := assert.New(test)
	a.Equals(result, expected)
	result = CheckIsTriangle(10, 10, 10)
	expected = true
	a.Equals(result, expected)
	result = CheckIsTriangle(0,0,0)
	expected = false
	a.Equals(result,expected)
	result = CheckIsTriangle(10,10,13)
	expected = true
	a.Equals(result,expected)
}
