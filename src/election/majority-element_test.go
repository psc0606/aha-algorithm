package election

import "testing"

func TestMajorityElement(t *testing.T) {
	nums := []int{3, 2, 3}
	res := majorityElement(nums)
	expected := 3
	if res != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, res)
	}

	nums = []int{2, 2, 1, 1, 1, 2, 2}
	res = majorityElement(nums)
	expected = 2
	if res != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, res)
	}

	nums = []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	res = majorityElement(nums)
	expected = 2
	if res != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, res)
	}
}
