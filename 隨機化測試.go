package main

import (
	"math/rand"
	"testing"
	"time"
)

//回傳使用亂數rng產生的隨機長度與內容迴文
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)        //最大到24的隨機長度
	runes := make([]rune, n) //slice
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) //最大到'\u0999'的隨機rune
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindrome(t *testing.T) {
	//亂數產生器初始化
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
