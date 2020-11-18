package string

import "testing"

func TestFreqAlphabets(t *testing.T) {
	s := freqAlphabets("10#11#12")
	expect := "jkab"
	if s != expect {
		t.Errorf("expected is [%s], actual is [%s]", expect, s)
	}

	s = freqAlphabets("1326#")
	expect = "acz"
	if s != expect {
		t.Errorf("expected is [%s], actual is [%s]", expect, s)
	}

	s = freqAlphabets("25#")
	expect = "y"
	if s != expect {
		t.Errorf("expected is [%s], actual is [%s]", expect, s)
	}

	s = freqAlphabets("12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#")
	expect = "abcdefghijklmnopqrstuvwxyz"
	if s != expect {
		t.Errorf("expected is [%s], actual is [%s]", expect, s)
	}
}
