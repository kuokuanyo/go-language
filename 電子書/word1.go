//文字遊戲的工具
package word

//回報s是否正反讀均相同
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
