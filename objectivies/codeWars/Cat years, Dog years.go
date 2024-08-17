package kata

func CalculateYears(years int) (result [3]int) {
	var humanYears, catYears, dogYears int
	if years == 1 {
		humanYears = 1
		catYears = 15
		dogYears = 15
	} else if years == 2 {
		humanYears = 2
		catYears = 24
		dogYears = 24
	} else {
		humanYears = years
		catYears = 24 + 4*(years-2)
		dogYears = 24 + 5*(years-2)
	}
	return [3]int{humanYears, catYears, dogYears}
}
