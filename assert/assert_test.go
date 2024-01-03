package assert

import "testing"

func TestAssert(test *testing.T) {
	a := New(test)
	expected := &Assert{test: test}
	a.Equals(a.test,expected.test)
}
