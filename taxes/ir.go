package taxes

type IR struct {
	Rate                 float32
	Base                 float32
	ValueWithoutDiscount float32
	Discount             float32
	Value                float32
}

func NewIR(v float32) IR {
	ir := IR{}
	ir.base(v, NewINSS(v))
	ir.rateAndDiscount()
	ir.ValueWithoutDiscount = ir.Base * ir.Rate
	ir.Value = ir.ValueWithoutDiscount - ir.Discount
	return ir
}

func (i *IR) base(v float32, inss INSS) {
	i.Base = v - inss.Value
}

func (i *IR) rateAndDiscount() {
	IRBases := []float32{1903.98, 2826.65, 3751.05, 4664.68}
	IRInstallmentsDiscount := []float32{142.8, 354.8, 636.13, 869.36}
	base := i.Base

	switch {
	case base <= IRBases[0]:
		i.Rate, i.Discount = 0, 0
	case base <= IRBases[1]:
		i.Rate, i.Discount = 0.075, IRInstallmentsDiscount[0]
	case base <= IRBases[2]:
		i.Rate, i.Discount = 0.15, IRInstallmentsDiscount[1]
	case base <= IRBases[3]:
		i.Rate, i.Discount = 0.225, IRInstallmentsDiscount[2]
	default:
		i.Rate, i.Discount = 0.275, IRInstallmentsDiscount[3]
	}
}
