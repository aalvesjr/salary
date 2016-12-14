# Calculo do Imposto de Renda (IR)

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
./calcula_ir 5168.78

# => Salário => 5168.78
# => INSS => 568.57
# => Base IR => 4600.21
# => Aliquota IR => 22.50
# => IR sem desconto => 1035.05
# => Desconto do IR => 636.13
# => valor IR => 398.92
```
