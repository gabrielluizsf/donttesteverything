package price

import (
	"fmt"
)

func Value(userChoice string) float64 {
	var aVista bool
	switch userChoice {
	case "s":
		aVista = true
	case "n":
		aVista = false
	default:
		fmt.Println("Invalid Value")
	}
	return price(100, 10, aVista)
}

func price(preco float64, margem float64, aVista bool) float64 {
	var valorASerPago float64
	if aVista {
		if preco < 200 {
			desconto := 10.0
			valorASerPago = preco - desconto
			printValorASerPago(valorASerPago)
		}
		if preco >= 200 {
			desconto := 14.0
			valorASerPago = preco - desconto
			printValorASerPago(valorASerPago)
		}
	} else {
		fmt.Println("[Sem Desconto]: O valor a ser pago será de ", preco)
	}
	printLucro(valorASerPago, margem)
	return valorASerPago
}

func printValorASerPago(valorASerPago float64) {
	fmt.Println("O valor a ser pago será de ", valorASerPago)
}
func printLucro(pagamento, margem float64) {
	fmt.Println("O lucro foi de ", pagamento*margem)
}
