package calc

func Execute(operation string, number1 float64, number2 float64) int {
	switch operation {
	case "add":
		return add(int(number1), int(number2))
	case "sub":
		return sub(int(number1), int(number2))
	case "mult":
		return mult(int(number1), int(number2))
	case "div":
		return div(int(number1), int(number2))
	default:
		return 0
	}
}

func add(valuea, valueb int) int {
	return valuea + valueb
}
func sub(valuea, valueb int) int {
	return valuea - valueb
}
func div(valuea, valueb int) int {
	if valuea == 0 || valueb == 0 {
		return 0
	}
	return valuea / valueb
}
func mult(valuea, valueb int) int {
	return valuea * valueb
}
