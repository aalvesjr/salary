package impostos

type INSS struct {
	Aliquota float32
	Base     float32
	Valor    float32
}

func NewINSS(v float32) INSS {
	inss := INSS{Base: v}
	inss.aliquotaEBaseINSS()
	inss.Valor = inss.Base * inss.Aliquota
	return inss
}

func (i *INSS) aliquotaEBaseINSS() {
	basesAliquotaINSS := []float32{1556.94, 2594.92, 5189.82}
	b := i.Base

	switch {
	case b <= basesAliquotaINSS[0]:
		i.Aliquota = 0.08
	case b <= basesAliquotaINSS[1]:
		i.Aliquota = 0.09
	case b <= basesAliquotaINSS[2]:
		i.Aliquota = 0.11
	default:
		i.Aliquota, i.Base = 0.11, basesAliquotaINSS[2]
	}
}
