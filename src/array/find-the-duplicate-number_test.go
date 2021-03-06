package array

import "testing"

func TestFindDuplicate(t *testing.T) {
	nums := []int{2, 3, 1, 0, 6, 5, 3}
	check(nums)
	actual := findDuplicate(nums)
	expected := 3
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	nums = []int{1, 3, 2, 4, 1}
	actual = findDuplicate(nums)
	expected = 1
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}
}

func TestFindDuplicate2(t *testing.T) {
	nums := []int{2, 3, 1, 0, 6, 5, 3}
	check(nums)
	actual := findDuplicate2(nums)
	expected := 3
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	nums = []int{1, 3, 2, 4, 1}
	actual = findDuplicate2(nums)
	expected = 1
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}
}

func TestFindDuplicate3(t *testing.T) {
	nums := []int{2, 3, 1, 0, 6, 5, 3}
	check(nums)
	actual := findDuplicate3(nums)
	expected := 3
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	nums = []int{1, 3, 2, 4, 1}
	actual = findDuplicate3(nums)
	expected = 1
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}
}

func check(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] >= n {
			panic("illegal input array")
		}
	}
}
