package sort

func BubbleSort(a []int, n int) {
	for i := 0; i < n; i++ {
		flag := false
		for j := i; j < n-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

func InsertSort(a []int, n int) {
	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		val := a[i]
		for j := i - 1; j >= 0; j-- {
			if a[j] > val {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = val
	}
}

func selectionSort(a []int, n int) {
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}

}

func MergeSort(a []int, n int) {
	if n <= 1 {
		return
	}
	mergeSort(a, 0, n-1)
}

func mergeSort(a []int, start, end int) {
	if start >= end {
		return
	}

	mid := (end - start) / 2
	mergeSort(a, start, mid)
	mergeSort(a, mid+1, end)
	merge(a, start, mid, end)
}

func merge(a []int, start, mid, end int) {
	tmp := make([]int, end-start+1)
	i := start
	j := mid + 1
	k := 0
	for i <= mid && j <= end {
		if a[i] < a[j] {
			tmp[k] = a[i]
			i++
		} else {
			tmp[k] = a[j]
			j++
		}
		k++
	}

	for ; i <= mid; i++ {
		tmp[k] = a[i]
		k++
	}

	for ; j <= end; j++ {
		tmp[k] = a[j]
		k++
	}

	copy(a[start:end+1], tmp)
}

func QuickSort(a []int, n int) {
	quickSort(a, 0, n-1)
}

func quickSort(a []int, p, r int) {
	if p >= r {
		return
	}

	i := partition(a, p, r)
	quickSort(a, p, i-1)
	quickSort(a, i+1, r)
}

func partition(a []int, p, r int) {
	pivot := a[r]
	i := p
	for j := p; j < r; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], pivot = pivot, a[i]
}
