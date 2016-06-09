package tornago

import (
	"testing"
)

var groupTests = []struct {
	i int
	e uint16
}{
	{
		i: -1,
		e: 0,
	},
	{
		i: 0,
		e: 0x1,
	},
	{
		i: 20,
		e: 0xFFFF,
	},
	{
		i: 5,
		e: 0x20,
	},
}

func TestGroup(t *testing.T) {
	for i, test := range groupTests {
		if f := Group(test.i); f != test.e {
			t.Errorf("[%d] Group(%d) = 0x%.4X, want 0x%.4X", i, test.i, f, test.e)
		}
	}
}

func TestMask(t *testing.T) {
	for i, test := range groupTests {
		if f := Mask(test.i); f != test.e {
			t.Errorf("[%d] Mask(%d) = 0x%.4X, want 0x%.4X", i, test.i, f, test.e)
		}
	}
}

func TestConstantGroups(t *testing.T) {
	if GroupNone != 0 {
		t.Error("GroupNone is not 0")
	}

	groups := [...]uint16{GroupNone, Group1, Group2, Group3, Group4, Group5, Group6, Group7,
		Group8, Group9, Group10, Group11, Group12, Group13, Group14, Group15}
	expected := [...]uint16{0, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096,
		8192, 16384, 32768}

	for n, group := range groups {
		if group != expected[n] {
			t.Errorf("[%d] not equal", n)
		}
	}
}
