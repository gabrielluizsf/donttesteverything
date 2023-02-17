package main

import (
  "fmt"
)

func main(){
  var userChoice string;
  var aVista bool;
  fmt.Printf("Pagamento a vista?");
  fmt.Scanf("%s",&userChoice);
  switch userChoice{
    case "s":
    aVista = true
    case "n":
    aVista = false
    default:
    fmt.Println("Invalid Value");
  }
  GetPrice(100,10,aVista);
}

func GetPrice(preco float64, margem float64, aVista bool) float64{
  var valorASerPago float64;
  if aVista{
      if preco < 200{
      desconto := 10.0;
      valorASerPago = preco - desconto;
      PrintValorASerPago(valorASerPago);
      }
      if preco >= 200{
      desconto := 14.0;
      valorASerPago = preco - desconto;
      PrintValorASerPago(valorASerPago);
      }
  }else{
    fmt.Println("[Sem Desconto]: O valor a ser pago será de ",preco);  
  }
  PrintLucro(valorASerPago,margem);
    return valorASerPago;
} 

func PrintValorASerPago(valorASerPago float64){
  fmt.Println("O valor a ser pago será de ", valorASerPago);
}
func PrintLucro(pagamento , margem float64){
  fmt.Println("O lucro foi de ", pagamento * margem);
}
