package main

import "testing";

func TestGetTypeTriangle(test *testing.T){
	checkExpectedResult := func(test *testing.T, result,expected int){
		test.Helper();
		if result != expected {
			test.Errorf("Result: %d Expected: %d",result,expected)
		}
	}
	test.Run("É Inválido", func(test *testing.T){
		InválidoResult   := GetTypeTriangle(10,12,-13);
		InválidoExpected := 0;
			checkExpectedResult(test,InválidoResult,InválidoExpected);	
	});

	test.Run("Não é um triangulo", func(test *testing.T){
		nãoéumtrianguloResult   := GetTypeTriangle(0,0,0);
		nãoéumtrianguloExpected := 0;
			checkExpectedResult(test,nãoéumtrianguloResult,nãoéumtrianguloExpected);	
	});

	test.Run("É Escaleno", func(test *testing.T){
		escalenoResult   := GetTypeTriangle(10,12,13);
		escalenoExpected := 30;
			checkExpectedResult(test,escalenoResult,escalenoExpected);	
	});
	test.Run("É Isósceles", func(test *testing.T){
		isoscelesResult   := GetTypeTriangle(12,12,13);
		isoscelesExpected := 20;
			checkExpectedResult(test,isoscelesResult,isoscelesExpected);	
	});
	test.Run("É Equilátero", func(test *testing.T){
		equilateroResult   := GetTypeTriangle(13,13,13);
		equilateroExpected := 10;
			checkExpectedResult(test,equilateroResult,equilateroExpected);	
	});
}