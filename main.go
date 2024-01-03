package main

import (
	"fmt"

	"github.com/gabrielluizsf/donttesteverything/triangle"
)

var (
	triangleType int
)

func main() {
	x := 10
	y := 12
	z := 13

	isTriangle := triangle.CheckIsTriangle(x, y, z)
	fmt.Printf("is triangle: %v", isTriangle)

	numeroReferenteAoTipoDoTriangulo := GetTypeTriangle(x, y, z)
	printTriangleTypeName(numeroReferenteAoTipoDoTriangulo)
}

func printTriangleTypeName(number int) {
	switch number {
	case 0:
		fmt.Println("Não é um Triângulo")
	case 10:
		fmt.Println("O Triângulo é Equilátero")
	case 20:
		fmt.Println("O Triângulo pe Isósceles")
	case 30:
		fmt.Println("O Triângulo é Escaleno")
	default:
		fmt.Println("Invalid Value")
	}
}

func GetTypeTriangle(x, y, z int) int {

	isTriangle := triangle.CheckIsTriangle(x, y, z)
	if !isTriangle {
		triangleType = 0
	} else if x == y && y == z && z == x {
		triangleType = 10
	} else if x != y && y != z && x != z {
		triangleType = 30
	} else if x == y || z == x || y == z {
		triangleType = 20
	}
	return triangleType
}
