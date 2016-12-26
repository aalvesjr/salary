# Calculo do Imposto de Renda (IR)
![alt Programa executado no terminal](https://s28.postimg.org/q9bid53el/Captura_de_tela_de_2016_12_25_22_49_50.png)

## Setup

```
mkdir -p $GOPATH/src/github.com/aalvesjr
cd $GOPATH/src/github.com/aalvesjr
git clone https://github.com/aalvesjr/calcula_ir.git
cd calcula_ir
go build
```

## Running

```
./calcula_ir 5168.78 127.79

#=> Salário Bruto   => R$ 5168.78
#=> Descontos       => R$ 127.79
#=> ------------INSS-------------
#=> Base INSS       => R$ 5168.78
#=> Aliquota INSS   => 11.00%
#=> Valor INSS      => R$ 568.57
#=> -------------IR--------------
#=> Base IR         => R$ 4600.21
#=> Aliquota IR     => 22.50%
#=> IR sem desconto => R$ 1035.05
#=> Desconto do IR  => R$ 636.13
#=> Valor IR        => R$ 398.92
#=> -----------------------------
#=> Salário Liquido => R$ 4073.51

```
