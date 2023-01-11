package utils

func romanToInt(s string) int {
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var sum int
	for i, v := range s {
		if value, ok := m[string(v)]; ok {
			var next uint8
			if len(s)-1 != i {
				next = s[i+1]
			} else {
				next = s[i]
			}

			nextVal := m[string(next)]
			if value < nextVal {
				sum -= value
			} else {
				sum += value
			}

		}
	}

	return sum
}
