# CONCEITOS de testes

1- Unit tests: teste de unidade

2- Pirâmide de testes

```go
      / \
      E2E
      ---
  /        \
  Integration
  ----------
 /   Unit   \
 ------------

 ```

 3- Testes de perfomance

 Teste de carga: testa se de 500 mil requests suporta 500 mil requests

 Teste de extresse: testa se de 500 mil requests suporta 550 mil requests

 4- teste de usabilidade: testa em SO diferentes
    - teste de portabilidade: testa o sistema em versoes de SO diferentes

 5- teste de caixa preta: não ve o codigo
 
 6- teste de caixa branca: ve o codigo

 7- teste de aceitação: o cliente analiza se os requisitos estão de acordo com a entrega

 8- teste de integração: testa como partes do sistema se integram

 9- teste fim a fim [E2E]: testa a interface gráfica do sistema