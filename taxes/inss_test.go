package taxes

import "testing"

var newINSSTests = []struct {
	i float32 // input
	a float32 // rate expected
	b float32 // base expected
	v float32 // value expected
}{
	{0, 0.08, 0, 0},
	{1000, 0.08, 1000, 80},
	{2000, 0.09, 2000, 180},
	{3000, 0.11, 3000, 330},
	{6000, 0.11, 5531.31, 608.4441},
}

func TestNewINSS(t *testing.T) {
	for _, n := range newINSSTests {
		actual := NewINSS(n.i)

		if actual.Rate != n.a {
			t.Errorf("NewINSS(%v).Rate: expected %v, actual %v", n.i, n.a, actual.Rate)
		}

		if actual.Base != n.b {
			t.Errorf("NewINSS(%v).Base: expected %v, actual %v", n.i, n.b, actual.Base)
		}

		if actual.Value != n.v {
			t.Errorf("NewINSS(%v).Value: expected %v, actual %v", n.i, n.v, actual.Value)
		}
	}
}
