package main

func BestCowLine(s string) string {
	var result []byte
	n := len(s)
	for i := 0; i < n; i++ {
		if s[0] < s[len(s)-1] {
			result = append(result, s[0])
			s = s[1:]
		} else {
			result = append(result, s[len(s)-1])
			s = s[:len(s)-1]
		}
	}
	return string(result)
}
