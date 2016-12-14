package models

func AliquotaINSS(salario float32, basesAliquotaINSS []float32) int {
	switch {
	case salario <= basesAliquotaINSS[0]:
		return 8
	case salario <= basesAliquotaINSS[1]:
		return 9
	default:
		return 11
	}
}

func INSS(salario float32, aliquota int, maximaBaseINSS float32) float32 {
	percentual := float32(aliquota) / 100

	if salario > maximaBaseINSS {
		return maximaBaseINSS * percentual
	} else {
		return salario * percentual
	}
}
