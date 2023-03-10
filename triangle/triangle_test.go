package triangle

import "testing";

func TestCheckIsTriangle(test *testing.T){
	result   := CheckIsTriangle(10,12,13);
	expected := true;

	if result != expected{
		test.Errorf("Result: %v  Expected: %v",result,expected);
	}
}