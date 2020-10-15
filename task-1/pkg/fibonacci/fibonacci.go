package fibonacci

func Calculate(n int) int {
	arr := []int{0, 1}

	if n < 3 {
		return arr[n-1]
	}

	for i := 2; i < n; i++ {
		arr = append(arr, arr[i-1]+arr[i-2])
	}

	return arr[n-1]
}
