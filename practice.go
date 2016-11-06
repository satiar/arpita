SELECTION SORT:
Time complexity : O(n2)
Space : O(1)
func selectionSort(arr []int) []int{
	for i := 0; i < len(arr) ; i++ {
		min := i
		for j := i + 1; j < len(arr) ; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr = swap(arr, i, min)
	}
	fmt.Println(arr)
	return arr
}

func swap(arr []int, i, j int) []int{
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
	return arr
}


