package models

import (
	impostos "github.com/aalvesjr/salario/impostos"
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

func (s Salario) Valor() float32 {
	return s.Bruto
}

func NewSalario(valor, descontos float32) Salario {
	salario := Salario{Bruto: valor, Descontos: descontos}

	// Calculo do INSS
	salario.AliquotaINSS, salario.BaseINSS = impostos.AliquotaEBaseINSS(salario.Bruto)
	salario.INSS = impostos.INSS(salario.Bruto)

	// Calculo do IR
	salario.BaseIR = impostos.BaseIR(salario.Bruto)
	salario.AliquotaIR, salario.DescontoIR = impostos.AliquotaEDescontoIR(salario.Bruto)
	salario.IRSemDesconto = impostos.IRSemParcelaDesconto(salario.Bruto)
	salario.IR = impostos.IR(salario.Bruto)

	salario.AplicaDescontos()
	return salario
}

func (s *Salario) AplicaDescontos() {
	s.Liquido = s.Bruto - (s.INSS + s.IR + s.Descontos)
}
