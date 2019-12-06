package word

import (
	"math/rand"
	"testing"
	"time"
)

func TestIsPalindrome(t *testing.T) {
	// 基于表的测试
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"🦢a🦢", true},
		{"A man, a plan, a canal: Panama", true},
		{"palindrome", false},
		{"desserts", false},
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

// randomPalindrome 返回一个回文字符串，它的长度和内容都是随机数生成器rng生成的
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // 随机字符串最大长度24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // 随机字符最大是'\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

// TestRandomPalindrome 测试随机生成的回文字符串
func TestRandomPalindrome(t *testing.T) {
	// 初始化一个伪随机数生成器
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
