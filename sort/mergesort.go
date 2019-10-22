package sort

func mergesort(a []int) {
	msort(a, 0, len(a)-1)
}

func msort(a []int, l, r int) {
	if l == r {
		return
	}
	mid := (l + r) >> 1
	msort(a, l, mid)
	msort(a, mid+1, r)
	mrg(a, l, mid, r)
}

func mrg(a []int, l, m, r int) {
	var temp []int

	left := l
	p := m + 1 //注意这里要加一，区间为[l,m]以及（m,r]，不能搞错

	for l <= m && p <= r {
		if a[l] > a[p] {
			temp = append(temp, a[p])
			p += 1
		} else {
			temp = append(temp, a[l])
			l += 1
		}
	}

	for l <= m {
		temp = append(temp, a[l])
		l += 1
	}

	for p <= r {
		temp = append(temp, a[p])
		p += 1
	}

	for _, v := range temp {
		a[left] = v
		left += 1
	}
}
