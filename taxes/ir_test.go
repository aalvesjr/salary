package taxes

import "testing"

var newIRTests = []struct {
	i   float32 // input
	a   float32 // rate expected
	b   float32 // base expected
	vsd float32 // value without discount expected
	d   float32 // discount expected
	v   float32 // value expected
}{
	{0, 0, 0, 0, 0, 0},
	{1000, 0, 920, 0, 0, 0},
	{2500, 0.075, 2275, 170.625, 142.80, 27.824997},
	{3500, 0.15, 3115, 467.25003, 354.80, 112.45004},
	{4500, 0.225, 4005, 901.125, 636.13, 264.995},
	{5500, 0.275, 4895, 1346.125, 869.36, 476.765},
}

func TestNewIR(t *testing.T) {
	for _, n := range newIRTests {
		actual := NewIR(n.i)

		if actual.Rate != n.a {
			t.Errorf("NewIR(%v).Rate: expected %v, actual %v", n.i, n.a, actual.Rate)
		}

		if actual.Base != n.b {
			t.Errorf("NewIR(%v).Base: expected %v, actual %v", n.i, n.b, actual.Base)
		}

		if actual.ValueWithoutDiscount != n.vsd {
			t.Errorf("NewIR(%v).ValueWithoutDiscount: expected %v, actual %v", n.i, n.vsd, actual.ValueWithoutDiscount)
		}

		if actual.Discount != n.d {
			t.Errorf("NewIR(%v).Discount: expected %v, actual %v", n.i, n.d, actual.Discount)
		}

		if actual.Value != n.v {
			t.Errorf("NewIR(%v).Value: expected %v, actual %v", n.i, n.v, actual.Value)
		}
	}
}
