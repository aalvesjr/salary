package taxes

type INSS struct {
	Rate  float32
	Base  float32
	Value float32
}

func NewINSS(v float32) INSS {
	inss := INSS{Base: v}
	inss.baseAndRate()
	inss.Value = inss.Base * inss.Rate
	return inss
}

func (i *INSS) baseAndRate() {
	INSSRates := []float32{1556.94, 2594.92, 5189.82}
	b := i.Base

	switch {
	case b <= INSSRates[0]:
		i.Rate = 0.08
	case b <= INSSRates[1]:
		i.Rate = 0.09
	case b <= INSSRates[2]:
		i.Rate = 0.11
	default:
		i.Rate, i.Base = 0.11, INSSRates[2]
	}
}
