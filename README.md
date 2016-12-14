# Calculo do Imposto de Renda (IR)
![alt Programa executado no terminal](https://s23.postimg.org/hu8yae1zv/Captura_de_tela_de_2016_12_14_01_15_39.png)

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

#=> Salário         => R$ 5168.78
#=> Base INSS       => R$ 5168.78
#=> Aliquota INSS   => 11.00
#=> INSS            => R$ 568.57
#=> Base IR         => R$ 4600.21
#=> Aliquota IR     => 22.50
#=> IR sem desconto => R$ 1035.05
#=> Desconto do IR  => R$ 636.13
#=> valor IR        => R$ 398.92
```
