package main

import (
	"fmt"

	"github.com/gabrielluizsf/triangle"
)

var (
	triangleType int
);

func main(){
	a := 10;
	b := 12;
	c := 13;
	/*
	isTriangle := triangle.CheckIsTriangle(a,b,c);
		fmt.Printf("is triangle: %v",isTriangle);
	*/
	numeroReferenteAoTipoDoTriangulo := GetTypeTriangle(a,b,c);
		printTriangleTypeName(numeroReferenteAoTipoDoTriangulo);
}

func printTriangleTypeName(number int){
	switch(number){
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

func GetTypeTriangle(a,b,c int)int{	

	isTriangle := triangle.CheckIsTriangle(a,b,c);
	if !isTriangle{ 
			triangleType = 0;	
		} else if a == b && b == c && c == a{
			triangleType = 10;
		}else if a != b && b != c && a != c{
			triangleType = 30;
		}else if a == b || c == a || b == c{
			triangleType = 20
		}
	return triangleType;
}