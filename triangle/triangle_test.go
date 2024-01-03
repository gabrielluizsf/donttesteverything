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
}
