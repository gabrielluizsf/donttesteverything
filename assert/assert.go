package assert

import "testing"

type Assert struct {
	test *testing.T
}

func New(test *testing.T) *Assert {
	return &Assert{test: test}
}

func (a *Assert) Equals(result, expected interface{}) {
	if result != expected {
		a.test.Errorf("Resultado: %d, Esperado: %d", result, expected)
	}
}