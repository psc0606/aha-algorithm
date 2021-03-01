package optimalization

import (
	"testing"
)

func TestCanJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	actual := canJump(nums)
	expected := true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	nums = []int{3, 2, 1, 0, 4}
	actual = canJump(nums)
	expected = false
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	nums = []int{4, 2, 1, 0, 4}
	actual = canJump(nums)
	expected = true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	nums = []int{3, 3, 1, 0, 4}
	actual = canJump(nums)
	expected = true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	nums = []int{1}
	actual = canJump(nums)
	expected = true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	nums = []int{0}
	actual = canJump(nums)
	expected = true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	nums = []int{0, 1}
	actual = canJump(nums)
	expected = false
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}

func TestCanJumpWithOptimize(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	actual := canJumpWithOptimize(nums)
	expected := true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	nums = []int{3, 2, 1, 0, 4}
	actual = canJumpWithOptimize(nums)
	expected = false
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	nums = []int{4, 2, 1, 0, 4}
	actual = canJumpWithOptimize(nums)
	expected = true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	nums = []int{3, 3, 1, 0, 4}
	actual = canJumpWithOptimize(nums)
	expected = true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	nums = []int{1}
	actual = canJumpWithOptimize(nums)
	expected = true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	nums = []int{0}
	actual = canJumpWithOptimize(nums)
	expected = true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	nums = []int{0, 1}
	actual = canJumpWithOptimize(nums)
	expected = false
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
