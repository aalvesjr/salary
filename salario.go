package salario

import (
	"github.com/aalvesjr/salario/impostos"
)

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

func NewSalario(valor, descontos float32) Salario {
	s := Salario{Bruto: valor, Descontos: descontos}
	inss := impostos.NewINSS(valor)
	ir := impostos.NewIR(valor)

	s.AliquotaINSS, s.BaseINSS, s.INSS = inss.Aliquota, inss.Base, inss.Valor
	s.AliquotaIR, s.BaseIR, s.IRSemDesconto, s.DescontoIR, s.IR = ir.Aliquota, ir.Base, ir.ValorSemDesconto, ir.Desconto, ir.Valor

	s.Liquido = s.Bruto - (s.INSS + s.IR + s.Descontos)
	return s
}
