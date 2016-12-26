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
	salario := Salario{Bruto: valor, Descontos: descontos}

	salario.calculaINSS()
	salario.calculaIR()
	salario.aplicaDescontos()
	return salario
}

func (s *Salario) aplicaDescontos() {
	s.Liquido = s.Bruto - (s.INSS + s.IR + s.Descontos)
}

func (s *Salario) calculaINSS() {
	s.AliquotaINSS, s.BaseINSS = impostos.AliquotaEBaseINSS(s.Bruto)
	s.INSS = impostos.INSS(s.Bruto)
}

func (s *Salario) calculaIR() {
	s.BaseIR = impostos.BaseIR(s.Bruto)
	s.AliquotaIR, s.DescontoIR = impostos.AliquotaEDescontoIR(s.Bruto)
	s.IRSemDesconto = impostos.IRSemParcelaDesconto(s.Bruto)
	s.IR = impostos.IR(s.Bruto)
}
