package sort

func bubbleSort(a []int) {
	var length = len(a)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if a[i] > a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
}
