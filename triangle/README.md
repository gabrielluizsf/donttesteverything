# Teste de Software

## **Triangulo**

```bash
    A
  /    \ 
B ----- C

```

```go

a + b > c && c + b > a && c + a > b

```

### Equilátero 

**Se todos os lados possuem a mesma medida**

```go

 a == b && b == c && c == a

```


### Isósceles
**Se dois lados possuem a mesma medida**

```go

a == b || c == a || b == c

```

### Escaleno
**Se todos os lados possuem medidas diferentes**

```go

a != b && b != c && a != c

```

## Tipos de testes

- Teste de Unidade
- Teste de Integração
- Teste de Sistema   --> E2E = o testador não entende da regra de negócio
- Teste de Aceitação --> E2E = o testador é o clente 
- Teste de Regressão  = Testa se as alterações não regridem a qualidade do sistema  
- Caixa-Branca = Tem que olhar o código para testar
- Caixa-Preta  =  Não olha o código fonte --> Normalmente E2E = Teste de Sistema || Teste de Aceitação

### Exercício
2112332
