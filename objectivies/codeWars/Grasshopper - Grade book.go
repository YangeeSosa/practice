package kata

func GetGrade(a, b, c int) rune {
	avg := (a + b + c) / 3
	switch {
	case 90 <= avg:
		return 'A'
	case 80 <= avg:
		return 'B'
	case 70 <= avg:
		return 'C'
	case 60 <= avg:
		return 'D'
	default:
		return 'F'
	}
}
