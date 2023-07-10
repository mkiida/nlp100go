package main

// メイン関数
func main() {

	var ans1 = nlp00("stressed")
	println(ans1)

	var ans2 = nlp01("パタトクカシーー")
	println(ans2)
}

// 00. 文字列の逆順
func nlp00(s string) string {
	rs := []rune(s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}

	return string(rs)
}

// 01. 「パタトクカシーー」
func nlp01(s string) string {
	rs := []rune(s)

	return string(rs)
}
