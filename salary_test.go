package salary

import "testing"

var newSalaryTests = []struct {
	i    float32 // input
	d    float32 // discounts
	b    float32 // gross expected
	ain  float32 // inss rate expected
	bin  float32 // inss base expected
	vin  float32 // inss value expected
	air  float32 // ir rate expected
	bir  float32 // ir base expected
	vdir float32 // ir discount expected
	virs float32 // ir without discount expected
	vir  float32 // ir value expected
	l    float32 // net expected
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

func TestNewSalary(t *testing.T) {
	for _, n := range newSalaryTests {
		actual := NewSalary(n.i, n.d)

		if actual.Gross != n.b {
			t.Errorf("NewSalary(%v).Gross: expected %v, actual %v", n.i, n.b, actual.Gross)
		}

		if actual.INSSRate != n.ain {
			t.Errorf("NewSalary(%v).INSSRate: expected %v, actual %v", n.i, n.ain, actual.INSSRate)
		}

		if actual.INSSBase != n.bin {
			t.Errorf("NewSalary(%v).INSSBase: expected %v, actual %v", n.i, n.bin, actual.INSSBase)
		}

		if actual.INSS != n.vin {
			t.Errorf("NewSalary(%v).INSS: expected %v, actual %v", n.i, n.vin, actual.INSS)
		}

		if actual.Discounts != n.d {
			t.Errorf("NewSalary(%v).Discounts: expected %v, actual %v", n.i, n.d, actual.Discounts)
		}

		if actual.IRRate != n.air {
			t.Errorf("NewSalary(%v).IRRate: expected %v, actual %v", n.i, n.air, actual.IRRate)
		}

		if actual.IRBase != n.bir {
			t.Errorf("NewSalary(%v).IRBase: expected %v, actual %v", n.i, n.bir, actual.IRBase)
		}

		if actual.IR != n.vir {
			t.Errorf("NewINSS(%v).IR: expected %v, actual %v", n.i, n.vir, actual.IR)
		}

		if actual.IRWithoutDiscount != n.virs {
			t.Errorf("NewINSS(%v).IRWithoutDiscount: expected %v, actual %v", n.i, n.virs, actual.IRWithoutDiscount)
		}

		if actual.IRDiscount != n.vdir {
			t.Errorf("NewINSS(%v).IRDiscount: expected %v, actual %v", n.i, n.vdir, actual.IRDiscount)
		}

		if actual.Net != n.l {
			t.Errorf("NewINSS(%v).Net: expected %v, actual %v", n.i, n.l, actual.Net)
		}
	}
}
