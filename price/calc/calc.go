package calc

func Calc(userChoice string, number1 float64, number2 float64){
  switch userChoice{
  case "add":
	Add(int(number1),int(number2));
  case "sub":
	Sub(int(number1),int(number2));
  }
}

func Add(valuea, valueb int) int{
	return valuea + valueb;
}
func Sub(valuea, valueb int)int{
	return valuea - valueb;
}
