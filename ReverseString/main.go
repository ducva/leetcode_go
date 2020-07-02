package ReverseString

func ReverseString(s []byte) {
	end := len(s) - 1
	if end <= 0 {
		return
	}
	start := 0
	for ;start < end; {
		temp := s[end]
		s[end] = s[start]
		s[start] = temp
		start ++
		end --
	}
}
