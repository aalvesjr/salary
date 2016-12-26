package models

import "github.com/aalvesjr/salario/models"

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
	salario.AliquotaINSS, salario.BaseINSS = models.AliquotaEBaseINSS(salario.Bruto)
	salario.INSS = models.INSS(salario.Bruto)

	// Calculo do IR
	salario.BaseIR = models.BaseIR(salario.Bruto)
	salario.AliquotaIR, salario.DescontoIR = models.AliquotaEDescontoIR(salario.Bruto)
	salario.IRSemDesconto = models.IRSemParcelaDesconto(salario.Bruto)
	salario.IR = models.IR(salario.Bruto)

	salario.AplicaDescontos()
	return salario
}

func (s *Salario) AplicaDescontos() {
	s.Liquido = s.Bruto - (s.INSS + s.IR + s.Descontos)
}
