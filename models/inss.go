package models

func (s Salario) AliquotaEBaseINSS() (aliquota, base float32) {
	basesAliquotaINSS := []float32{1556.94, 2594.92, 5189.82}
	salario := float32(s)
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

func (s Salario) INSS() float32 {
	aliquota, base := s.AliquotaEBaseINSS()

	return aliquota * base / 100
}
