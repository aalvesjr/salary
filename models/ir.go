package models

func BaseIR(t tributavel) float32 {
	return float32(t.Valor()) - INSS(t)
}

func AliquotaEDescontoIR(t tributavel) (aliquota, desconto float32) {
	basesAliquotaIR := []float32{1903.98, 2826.65, 3751.05, 4664.68}
	parcelasDescontoIR := []float32{142.8, 354.8, 636.13, 869.36}
	base := BaseIR(t)

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

func IRSemParcelaDesconto(t tributavel) float32 {
	aliquota, _ := AliquotaEDescontoIR(t)

	return BaseIR(t) * aliquota / 100
}

func IR(t tributavel) float32 {
	_, desconto := AliquotaEDescontoIR(t)

	return IRSemParcelaDesconto(t) - desconto
}
