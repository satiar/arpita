SELECTION SORT:
Algo: Loop through elements, and swap ith element with the minimum in that pass
e.g. First pass, loop through all elements and keep a track of the min element
At end of first pass, swap 0th element with this min element
Repeat

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


BUBBLE SORT
Algo: Loop through all elements, and compare each ith element with i+1 th element
If a[i] >a[i+1], swap, else move on
At end of each pass, the "highest element" would have "bubbled up

Time complexity :O(n2)
Space complexity : O(n2)
/*Bubble up the largest number after each pass*/
func bubbleSort(arr []int) []int{
	for i := 0; i < len(arr)  ; i++ {
		//With each pass, the right elements are getting into their final spot, 
		// e.g. with first pass, highest element is in rightmost position
		// with second pass, second highest element is in the next rightmost position and so on
		// so we do not need to iterate through all the elements, but only up to n-i elements starting at the left
		for j := 1; j < len(arr)-i ; j++ {
			if arr[j-1] > arr[j] {
				arr = swap(arr, j, j-1)
			}
		}
	}
	return arr
}


INSERTION SORT
Algo: For each element, compare it to all elements to its left, and keep swapping till a[j] < a[j-1]
Time: O(n2), But better than selection and bubble sort in terms of #comparisons
Space

func insertionSort(arr []int) []int {
	/*For each element in the array, compare it with all elements to its left
	//Swap if it is less than an element on its left
	// With each pass , the left half is getting sorted */
	for i := 1 ; i < len(arr) ;i++ {
		for j := i ; j > 0  && arr[j] < arr[j-1] ; j-- {
			//Keep swapping with the element it is less than
			swap(arr, j , j-1)
		}
	}
	return arr
}


