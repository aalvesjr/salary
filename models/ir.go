package models

func (s Salario) BaseIR() float32 {
	return float32(s) - s.INSS()
}

func (s Salario) AliquotaEDescontoIR() (aliquota, desconto float32) {
	basesAliquotaIR := []float32{1903.98, 2826.65, 3751.05, 4664.68}
	parcelasDescontoIR := []float32{142.8, 354.8, 636.13, 869.36}

	switch {
	case s.BaseIR() <= basesAliquotaIR[0]:
		return 0, 0
	case s.BaseIR() <= basesAliquotaIR[1]:
		return 7.5, parcelasDescontoIR[0]
	case s.BaseIR() <= basesAliquotaIR[2]:
		return 15, parcelasDescontoIR[1]
	case s.BaseIR() <= basesAliquotaIR[3]:
		return 22.5, parcelasDescontoIR[2]
	default:
		return 27.5, parcelasDescontoIR[3]
	}
}

func (s Salario) IRSemParcelaDesconto() float32 {
	aliquota, _ := s.AliquotaEDescontoIR()

	return s.BaseIR() * aliquota / 100
}

func (s Salario) IR() float32 {
	_, desconto := s.AliquotaEDescontoIR()

	return s.IRSemParcelaDesconto() - desconto
}
