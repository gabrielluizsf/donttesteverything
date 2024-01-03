package assert

import "testing"

type t struct {
	test *testing.T
}

func New(test *testing.T) *t {
	return &t{test: test}
}

func (a *t) Equals(result, expected interface{}) {
	if result != expected {
		a.test.Errorf("Resultado: %d, Esperado: %d", result, expected)
	}
}