package ordenamiento

func Ordenamiento(s []int) (slice []int) {
	var min int
	n := len(s)
	print("we")
	for i := 0; i < n; i++ {
		min = calcMinimo(s[i], s)
		s = append(s, min)
	}

	return
}