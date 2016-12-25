package models

func AliquotaEBaseINSS(t tributavel) (aliquota, base float32) {
	basesAliquotaINSS := []float32{1556.94, 2594.92, 5189.82}
	salario := float32(t.Valor())
	switch {
	case salario <= basesAliquotaINSS[0]:
		return 8.0, salario
	case salario <= basesAliquotaINSS[1]:
		return 9.0, salario
	case salario <= basesAliquotaINSS[2]:
		return 11.0, salario
	default:
		return 11.0, basesAliquotaINSS[2]
	}
}

func INSS(t tributavel) float32 {
	aliquota, base := AliquotaEBaseINSS(t)

	return aliquota * base / 100
}
