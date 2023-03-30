package aas

func Sum(v []int) int {
	r := 0
	for _, i := range v {
		r += i
	}

	return r
}

func SumAll(v ...[]int) []int {
	var r []int

	for _, i := range v {
		r = append(r, Sum(i))
	}

	return r
}

func SumAllTails(v ...[]int) []int {
	var s []int

	for _, i := range v {
		if len(i) > 0 {
			s = append(s, Sum(i[1:]))
		} else {
			s = append(s, 0)
		}
	}

	return s

}
