package optimalization

import "testing"

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	actual := coinChange(coins, 11)
	expected := 3
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	coins = []int{2}
	actual = coinChange(coins, 3)
	expected = -1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	coins = []int{2}
	actual = coinChange(coins, 0)
	expected = 0
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	coins = []int{1}
	actual = coinChange(coins, 1)
	expected = 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	coins = []int{1}
	actual = coinChange(coins, 2)
	expected = 2
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	coins = []int{2, 5, 10, 1}
	actual = coinChange(coins, 27)
	expected = 4
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	coins = []int{186, 419, 83, 408}
	actual = coinChange(coins, 6249)
	expected = 20
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
