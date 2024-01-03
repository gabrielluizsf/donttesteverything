package assert

import "testing"

func TestAssert(test *testing.T) {
	a := New(test)
	expected := &t{test: test}
	a.Equals(a.test,expected.test)
}
