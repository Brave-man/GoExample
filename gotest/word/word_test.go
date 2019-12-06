package word

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}

	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}

func TestIsPalindrome2(t *testing.T) {
	if IsPalindrome("pandas") {
		t.Error(`IsPalindrome("panda") = true`)
	}
}

func TestIsPalindrome3(t *testing.T) {
	if !IsPalindrome("🦢和🦢") {
		t.Error(`IsPalindrome("🦢和🦢") = false`)
	}
}

func TestIsPalindrome4(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}
