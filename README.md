# Salario
[![CircleCI](https://circleci.com/gh/aalvesjr/salario.svg?style=svg)](https://circleci.com/gh/aalvesjr/salario)

## Setup

```
mkdir -p $GOPATH/src/github.com/aalvesjr
cd $GOPATH/src/github.com/aalvesjr/salario
git clone https://github.com/aalvesjr/salario.git
cd salario
go build
```

## Using

Adicione o modelo `salario` ao projeto


```
import(
  salario "github.com/aalvesjr/salario"
)
```

E a função `func NewSalario(valor, desconto float32) Salario` estára disponivel para criar uma struct `Salario`

```
s := salario.NewSalario(5000.00, 107.32)
```

Com isso `s` será uma struct `Salario` com os seguintes atributos

```
type Salario struct {
	Bruto         float32
	AliquotaINSS  float32
	BaseINSS      float32
	INSS          float32
	BaseIR        float32
	AliquotaIR    float32
	DescontoIR    float32
	IRSemDesconto float32
	IR            float32
	Liquido       float32
	Descontos     float32
}
```
