package salario

import "testing"

var NewSalarioTests = []struct {
	i    float32 // input
	d    float32 // descontos
	b    float32 // bruto expected
	ain  float32 // aliquota inss expected
	bin  float32 // base inss expected
	vin  float32 // valor inss expected
	air  float32 // aliquota ir expected
	bir  float32 // base ir expected
	vdir float32 // valor do desconto do ir expected
	virs float32 // valor do ir sem desconto expected
	vir  float32 // valor do ir expected
	l    float32 // liquido expected
}{
	{0, 0, 0, 0.08, 0, 0, 0, 0, 0, 0, 0, 0},
	{1000, 0, 1000, 0.08, 1000, 80, 0, 920, 0, 0, 0, 920},
	{1000, 20, 1000, 0.08, 1000, 80, 0, 920, 0, 0, 0, 900},
	{2500, 0, 2500, 0.09, 2500, 225.00002, 0.075, 2275, 142.8, 170.625, 27.824997, 2247.175},
	{2500, 47.175, 2500, 0.09, 2500, 225.00002, 0.075, 2275, 142.8, 170.625, 27.824997, 2200},
	{3000, 0, 3000, 0.11, 3000, 330, 0.075, 2670, 142.8, 200.25002, 57.450012, 2612.55},
	{3000, 12.55, 3000, 0.11, 3000, 330, 0.075, 2670, 142.8, 200.25002, 57.450012, 2600},
	{3500, 0, 3500, 0.11, 3500, 385, 0.15, 3115, 354.8, 467.25003, 112.45004, 3002.55},
	{3500, 102.55, 3500, 0.11, 3500, 385, 0.15, 3115, 354.8, 467.25003, 112.45004, 2900},
	{4500, 0, 4500, 0.11, 4500, 495, 0.225, 4005, 636.13, 901.125, 264.995, 3740.005},
	{4500, 40.005, 4500, 0.11, 4500, 495, 0.225, 4005, 636.13, 901.125, 264.995, 3700},
	{5500, 0, 5500, 0.11, 5189.82, 570.8802, 0.275, 4929.1196, 869.36, 1355.5079, 486.14795, 4442.9717},
	{5500, 442.9717, 5500, 0.11, 5189.82, 570.8802, 0.275, 4929.1196, 869.36, 1355.5079, 486.14795, 4000.0002},
	{10000, 0, 10000, 0.11, 5189.82, 570.8802, 0.275, 9429.12, 869.36, 2593.008, 1723.6481, 7705.4717},
	{10000, 705.4717, 10000, 0.11, 5189.82, 570.8802, 0.275, 9429.12, 869.36, 2593.008, 1723.6481, 7000},
}

func TestNewSalario(t *testing.T) {
	for _, n := range NewSalarioTests {
		actual := NewSalario(n.i, n.d)

		if actual.Bruto != n.b {
			t.Errorf("NewSalario(%v).Bruto: expected %v, actual %v", n.i, n.b, actual.Bruto)
		}

		if actual.AliquotaINSS != n.ain {
			t.Errorf("NewSalario(%v).AliquotaINSS: expected %v, actual %v", n.i, n.ain, actual.AliquotaINSS)
		}

		if actual.BaseINSS != n.bin {
			t.Errorf("NewSalario(%v).BaseINSS: expected %v, actual %v", n.i, n.bin, actual.BaseINSS)
		}

		if actual.INSS != n.vin {
			t.Errorf("NewINSS(%v).INSS: expected %v, actual %v", n.i, n.vin, actual.INSS)
		}

		if actual.Descontos != n.d {
			t.Errorf("NewINSS(%v).INSS: expected %v, actual %v", n.i, n.d, actual.Descontos)
		}

		if actual.AliquotaIR != n.air {
			t.Errorf("NewSalario(%v).AliquotaIR: expected %v, actual %v", n.i, n.air, actual.AliquotaIR)
		}

		if actual.BaseIR != n.bir {
			t.Errorf("NewSalario(%v).BaseIR: expected %v, actual %v", n.i, n.bir, actual.BaseIR)
		}

		if actual.IR != n.vir {
			t.Errorf("NewINSS(%v).IR: expected %v, actual %v", n.i, n.vir, actual.IR)
		}

		if actual.IRSemDesconto != n.virs {
			t.Errorf("NewINSS(%v).IRSemDesconto: expected %v, actual %v", n.i, n.virs, actual.IRSemDesconto)
		}

		if actual.DescontoIR != n.vdir {
			t.Errorf("NewINSS(%v).DescontoIR: expected %v, actual %v", n.i, n.vdir, actual.DescontoIR)
		}

		if actual.Liquido != n.l {
			t.Errorf("NewINSS(%v).Liquido: expected %v, actual %v", n.i, n.l, actual.Liquido)
		}
	}
}
