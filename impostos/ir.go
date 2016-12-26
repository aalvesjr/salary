package impostos

func BaseIR(s float32) float32 {
	return float32(s) - INSS(s)
}

func AliquotaEDescontoIR(s float32) (aliquota, desconto float32) {
	basesAliquotaIR := []float32{1903.98, 2826.65, 3751.05, 4664.68}
	parcelasDescontoIR := []float32{142.8, 354.8, 636.13, 869.36}
	base := BaseIR(s)

	switch {
	case base <= basesAliquotaIR[0]:
		return 0, 0
	case base <= basesAliquotaIR[1]:
		return 7.5, parcelasDescontoIR[0]
	case base <= basesAliquotaIR[2]:
		return 15, parcelasDescontoIR[1]
	case base <= basesAliquotaIR[3]:
		return 22.5, parcelasDescontoIR[2]
	default:
		return 27.5, parcelasDescontoIR[3]
	}
}

func IRSemParcelaDesconto(s float32) float32 {
	aliquota, _ := AliquotaEDescontoIR(s)

	return BaseIR(s) * aliquota / 100
}

func IR(s float32) float32 {
	_, desconto := AliquotaEDescontoIR(s)

	return IRSemParcelaDesconto(s) - desconto
}
