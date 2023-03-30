package iteration

func Repeat(r string, t int) string {
	var rv string

	for i := 0; i < t; i++ {
		rv += r
	}

	return rv
}
