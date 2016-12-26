package impostos

func AliquotaEBaseINSS(s float32) (aliquota, base float32) {
	basesAliquotaINSS := []float32{1556.94, 2594.92, 5189.82}

	switch {
	case s <= basesAliquotaINSS[0]:
		return 8.0, s
	case s <= basesAliquotaINSS[1]:
		return 9.0, s
	case s <= basesAliquotaINSS[2]:
		return 11.0, s
	default:
		return 11.0, basesAliquotaINSS[2]
	}
}

func INSS(s float32) float32 {
	aliquota, base := AliquotaEBaseINSS(s)

	return aliquota * base / 100
}
