package gmachine_test

import (
	"testing"

	"gmachine"
)

func TestNew(t *testing.T) {
	t.Parallel()

	g := gmachine.New()
	wantMemSize := gmachine.DefaultMemSize
	gotMemSize := len(g.Memory)
	if wantMemSize != gotMemSize {
		t.Errorf("want %d words of memory, got %d", wantMemSize, gotMemSize)
	}

	var wantP uint64 = 0
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}

	var wantMemValue uint64 = 0
	gotMemValue := g.Memory[gmachine.DefaultMemSize-1]
	if wantMemValue != gotMemValue {
		t.Errorf("want last memory location to contain %d, got %d", wantMemValue, gotMemValue)
	}

	var wantA uint64 = 0
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
}

func TestHALT(t *testing.T) {
	t.Parallel()

	g := gmachine.New()
	g.RunProgram([]uint64{
		gmachine.OpHALT,
	})

	var wantP uint64 = 1
	if wantP != g.P {
		t.Errorf("want P value %d, got %d", wantP, g.P)
	}
}

func TestNOOP(t *testing.T) {
	t.Parallel()

	g := gmachine.New()
	g.RunProgram([]uint64{
		gmachine.OpNOOP,
		gmachine.OpHALT,
	})

	var wantP uint64 = 2
	if wantP != g.P {
		t.Errorf("want P value %d, got %d", wantP, g.P)
	}
}

func TestINCA(t *testing.T) {
	t.Parallel()

	g := gmachine.New()
	g.RunProgram([]uint64{
		gmachine.OpINCA,
	})

	var wantA uint64 = 1
	if wantA != g.A {
		t.Errorf("want A value %d, got %d", wantA, g.A)
	}
}

func TestDECA(t *testing.T) {
	t.Parallel()

	g := gmachine.New()
	g.RunProgram([]uint64{
		gmachine.OpSETA,
		2,
		gmachine.OpDECA,
	})

	var wantA uint64 = 1
	if wantA != g.A {
		t.Errorf("want A value %d, got %d", wantA, g.A)
	}
}

func TestSubtract2(t *testing.T) {
	t.Parallel()

	g := gmachine.New()

	testcases := []struct {
		input []uint64
		wantA uint64
	}{
		{[]uint64{gmachine.OpSETA, 3, gmachine.OpDECA, gmachine.OpDECA}, 1},
		{[]uint64{gmachine.OpSETA, 4, gmachine.OpDECA, gmachine.OpDECA}, 2},
		{[]uint64{gmachine.OpSETA, 5, gmachine.OpDECA, gmachine.OpDECA}, 3},
	}

	for _, tc := range testcases {
		g.RunProgram(tc.input)
		if tc.wantA != g.A {
			t.Errorf("want A value %d, got %d", tc.wantA, g.A)
		}
	}
}

func TestSETA(t *testing.T) {
	t.Parallel()

	g := gmachine.New()
	g.RunProgram([]uint64{
		gmachine.OpSETA,
		5,
	})

	var wantA, wantP uint64 = 5, 3
	if wantA != g.A {
		t.Errorf("want A value %d, got %d", wantA, g.A)
	}

	if wantP != g.P {
		t.Errorf("want P value %d, got %d", wantP, g.P)
	}
}
