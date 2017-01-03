package impostos

type IR struct {
	Aliquota         float32
	Base             float32
	ValorSemDesconto float32
	Desconto         float32
	Valor            float32
}

func NewIR(v float32) IR {
	ir := IR{}
	inss := NewINSS(v)
	ir.BaseIR(v, inss)
	ir.AliquotaEDescontoIR()
	ir.ValorSemDesconto = ir.Base * ir.Aliquota
	ir.Valor = ir.ValorSemDesconto - ir.Desconto
	return ir
}

func (i *IR) BaseIR(v float32, inss INSS) {
	i.Base = v - inss.Valor
}

func (i *IR) AliquotaEDescontoIR() {
	basesAliquotaIR := []float32{1903.98, 2826.65, 3751.05, 4664.68}
	parcelasDescontoIR := []float32{142.8, 354.8, 636.13, 869.36}
	base := i.Base

	switch {
	case base <= basesAliquotaIR[0]:
		i.Aliquota, i.Desconto = 0, 0
	case base <= basesAliquotaIR[1]:
		i.Aliquota, i.Desconto = 0.075, parcelasDescontoIR[0]
	case base <= basesAliquotaIR[2]:
		i.Aliquota, i.Desconto = 0.15, parcelasDescontoIR[1]
	case base <= basesAliquotaIR[3]:
		i.Aliquota, i.Desconto = 0.225, parcelasDescontoIR[2]
	default:
		i.Aliquota, i.Desconto = 0.275, parcelasDescontoIR[3]
	}
}
