package impostos

import "testing"

var newIRTests = []struct {
	i   float32 // input
	a   float32 // aliquota expected
	b   float32 // base expected
	vsd float32 // ValorSemDesconto expected
	d   float32 // desconto expected
	v   float32 // valor expected
}{
	{0, 0, 0, 0, 0, 0},
	{1000, 0, 920, 0, 0, 0},
	{2500, 0.075, 2275, 170.625, 142.80, 27.824997},
	{3500, 0.15, 3115, 467.25003, 354.80, 112.45004},
	{4500, 0.225, 4005, 901.125, 636.13, 264.995},
	{5500, 0.275, 4929.1196, 1355.5079, 869.36, 486.14795},
}

func TestNewIR(t *testing.T) {
	for _, n := range newIRTests {
		actual := NewIR(n.i)

		if actual.Aliquota != n.a {
			t.Errorf("NewINSS(%v).Aliquota: expected %v, actual %v", n.i, n.a, actual.Aliquota)
		}

		if actual.Base != n.b {
			t.Errorf("NewINSS(%v).Base: expected %v, actual %v", n.i, n.b, actual.Base)
		}

		if actual.ValorSemDesconto != n.vsd {
			t.Errorf("NewINSS(%v).ValorSemDesconto: expected %v, actual %v", n.i, n.vsd, actual.ValorSemDesconto)
		}

		if actual.Desconto != n.d {
			t.Errorf("NewINSS(%v).Desconto: expected %v, actual %v", n.i, n.d, actual.Desconto)
		}

		if actual.Valor != n.v {
			t.Errorf("NewINSS(%v).Valor: expected %v, actual %v", n.i, n.v, actual.Valor)
		}
	}
}
