package calc

import "testing"

func TestAdd(test *testing.T){
	result   := Add(10,10);
	expected := 20;
	if result != expected{
		test.Errorf("Resultado: %d, Esperado: %d",result,expected);
	}
}

func  TestSub(test *testing.T){
	result   := Sub(100,50);
	expected := 50;
	if result != expected{
		test.Errorf("Resultado: %d, Esperado: %d",result,expected);
	}
}