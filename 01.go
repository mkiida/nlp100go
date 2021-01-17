package main

// メイン関数
func main() {
	var ans = nlp01("stressed")
	print(ans)
}

// 00. 文字列の逆順
func nlp01(s string) string {
	rs := []rune(s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}

	return string(rs)
}
