package impostos

import "testing"

var newINSSTests = []struct {
	i float32 // input
	a float32 // aliquota expected
	b float32 // base expected
	v float32 // valor expected
}{
	{0, 0.08, 0, 0},
	{1000, 0.08, 1000, 80},
	{2000, 0.09, 2000, 180},
	{3000, 0.11, 3000, 330},
	{6000, 0.11, 5189.82, 570.8802},
}

func TestNewINSS(t *testing.T) {
	for _, n := range newINSSTests {
		actual := NewINSS(n.i)

		if actual.Aliquota != n.a {
			t.Errorf("NewINSS(%v).Aliquota: expected %v, actual %v", n.i, n.a, actual.Aliquota)
		}

		if actual.Base != n.b {
			t.Errorf("NewINSS(%v).Base: expected %v, actual %v", n.i, n.b, actual.Base)
		}

		if actual.Valor != n.v {
			t.Errorf("NewINSS(%v).Valor: expected %v, actual %v", n.i, n.v, actual.Valor)
		}
	}
}
