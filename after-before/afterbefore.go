package afterbefore

import "github.com/gabrielluizsf/donttesteverything/price/calc"

func DoubleNumber(number int) int {
	return calc.Execute("mult", float64(number), float64(2))
}

func IsNullOrEmpty(str string) bool {
	return str == ""
}

func Add(num, num2 int) int {
	return calc.Execute("add", float64(num), float64(num2))
}
