package afterbefore

import "testing"

func TestDoubleNumber(test *testing.T){
	result   :=   DoubleNumber(2);
	expected :=	4;

	if result != expected{
		test.Errorf("Result: %d , Expected: %d",result, expected);
	}
}

func TestIsNullOrEmpty(test *testing.T){
	result   :=  IsNullOrEmpty("ahfjsaf");
	expected := false;

	if result != expected{
		test.Errorf("Result: %t, Expected: %t", result, expected);
	}
}

func TestAdd(test *testing.T){
	result   := Add(10,10);
	expected := 20;

	if result != expected{
		test.Errorf("Result: %d, Expected: %d",result,expected)
	}
}