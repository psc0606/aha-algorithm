package optimalization

import (
	"testing"
)

func TestMaxProfitWithCoolDown(t *testing.T) {
	prices := []int{1, 2, 3, 0, 2}
	actual := maxProfitWithCoolDown(prices)
	expected := 3
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	prices = []int{1, 2, 3, 3, 2}
	actual = maxProfitWithCoolDown(prices)
	expected = 2
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	prices = []int{1, 2, 3, 3, 4}
	actual = maxProfitWithCoolDown(prices)
	expected = 3
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	prices = []int{1, 2, 3, 4, 5}
	actual = maxProfitWithCoolDown(prices)
	expected = 4
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	prices = []int{1}
	actual = maxProfitWithCoolDown(prices)
	expected = 0
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	prices = []int{1, 2}
	actual = maxProfitWithCoolDown(prices)
	expected = 1
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	prices = []int{1, 2, 3}
	actual = maxProfitWithCoolDown(prices)
	expected = 2
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	prices = []int{3, 2, 1}
	actual = maxProfitWithCoolDown(prices)
	expected = 0
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	prices = []int{3, 2, 1, 6}
	actual = maxProfitWithCoolDown(prices)
	expected = 5
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	prices = []int{2, 1, 4}
	actual = maxProfitWithCoolDown(prices)
	expected = 3
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}
}

func TestMaxProfitWithCoolDownBetter(t *testing.T) {
	prices := []int{1, 2, 3, 0, 2}
	actual := maxProfitWithCoolDownBetter(prices)
	expected := 3
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	prices = []int{1, 2, 3, 3, 2}
	actual = maxProfitWithCoolDownBetter(prices)
	expected = 2
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	prices = []int{1, 2, 3, 3, 4}
	actual = maxProfitWithCoolDownBetter(prices)
	expected = 3
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	prices = []int{1, 2, 3, 4, 5}
	actual = maxProfitWithCoolDownBetter(prices)
	expected = 4
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	prices = []int{1}
	actual = maxProfitWithCoolDownBetter(prices)
	expected = 0
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	prices = []int{1, 2}
	actual = maxProfitWithCoolDownBetter(prices)
	expected = 1
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	prices = []int{1, 2, 3}
	actual = maxProfitWithCoolDownBetter(prices)
	expected = 2
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	prices = []int{3, 2, 1}
	actual = maxProfitWithCoolDownBetter(prices)
	expected = 0
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	prices = []int{3, 2, 1, 6}
	actual = maxProfitWithCoolDownBetter(prices)
	expected = 5
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	prices = []int{2, 1, 4}
	actual = maxProfitWithCoolDownBetter(prices)
	expected = 3
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}
}
