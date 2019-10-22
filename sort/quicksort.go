package sort

func qsort(a []int) {
	if len(a) == 0 {
		return
	}

	i := 0
	j := len(a) - 1
	base := a[i]
	for i < j {
		fmt.Println(i, j)
		for i < j && a[j] >= base {
			j -= 1
		}
		a[i] = a[j]
		for i < j && a[i] < base {
			i += 1
		}
		a[j] = a[i]
	}
	a[i] = base
	qsort(a[:i])
	qsort(a[i+1:])
}
