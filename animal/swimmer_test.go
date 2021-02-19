package animal

import "testing"

type TestSwimmer struct{}

func (ts *TestSwimmer) Fins() int {
	return 999
}

func TestDescribeSwimmer(t *testing.T) {
	ts := &TestSwimmer{}

	if ts.Fins() != 999 {
		t.Errorf("expected swimmer to have fin count %d but received %d", 999, ts.Fins())
	}

	expected := "swims with 999 fins"
	if DescribeSwimmer(ts) != expected {
		t.Errorf("expected description \"%s\" but received \"%s\"", expected, DescribeSwimmer(ts))
	}
}
