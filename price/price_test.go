package price

import (
	"fmt"
	"testing"

	"github.com/gabrielluizsf/donttesteverything/assert"
)

func TestValue(test *testing.T){
	result := Value("s")
	expected := 90.0
	a := assert.New(test)
	a.Equals(result,expected)
	result = Value("n")
	fmt.Println(result)
}