package animal

import "testing"

type TestWalker struct{}

func (tw *TestWalker) Legs() int {
	return 999
}

func TestDescribeWalker(t *testing.T) {
	tw := &TestWalker{}

	if tw.Legs() != 999 {
		t.Errorf("expected walker to have leg count %d but received %d", 999, tw.Legs())
	}

	expected := "walks on 999 legs"
	if DescribeWalker(tw) != expected {
		t.Errorf("expected description \"%s\" but received %s", expected, DescribeWalker(tw))
	}
}
