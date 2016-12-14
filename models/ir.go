package models

func BaseIR(salario float32, inss float32) float32 {
	return salario - inss
}

func AliquotaEDescontoIR(baseIR float32, parcelasDescontoIR, basesAliquotaIR []float32) (aliquota, desconto float32) {
	switch {
	case baseIR <= basesAliquotaIR[0]:
		return 0, 0
	case baseIR <= basesAliquotaIR[1]:
		return 7.5, parcelasDescontoIR[0]
	case baseIR <= basesAliquotaIR[2]:
		return 15, parcelasDescontoIR[1]
	case baseIR <= basesAliquotaIR[3]:
		return 22.5, parcelasDescontoIR[2]
	default:
		return 27.5, parcelasDescontoIR[3]
	}
}

func IRSemParcelaDesconto(baseIR, aliquotaIR float32) float32 {
	return baseIR * aliquotaIR / 100
}

func IR(IRSemParcela, descontoIR float32) float32 {
	return IRSemParcela - descontoIR
}
