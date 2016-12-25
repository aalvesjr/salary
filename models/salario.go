package models

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
	salario.AliquotaINSS, salario.BaseINSS = AliquotaEBaseINSS(salario)
	salario.INSS = INSS(salario)

	// Calculo do IR
	salario.BaseIR = BaseIR(salario)
	salario.AliquotaIR, salario.DescontoIR = AliquotaEDescontoIR(salario)
	salario.IRSemDesconto = IRSemParcelaDesconto(salario)
	salario.IR = IR(salario)

	salario.AplicaDescontos()
	return salario
}

func (s *Salario) AplicaDescontos() {
	s.Liquido = s.Bruto - (s.INSS + s.IR + s.Descontos)
}
