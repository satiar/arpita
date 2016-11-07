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


MERGESORT  - it is not in-place
Divide array into smaller portions till only 1 element left in each,
Then merge each small portion with its neighbor to form a bigger sorted array.
So at each step of merge, both the left and right sub-array that are being merged are individually sorted
Time : O(n log n )   https://cdn.kastatic.org/ka-cs-algorithms/merge_sort_tree_4.png
Space: O(n) - extra space needed to store resultset

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) /2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	
	return merge(left, right)
}

//merge 2 sorted arrays
func merge( left, right []int) []int {
	result := make([]int, 0, len(left) + len(right)) 
	for len(left)> 0 || len(right) >0 {
		if len(left) == 0 {
			// left is all done, just add all sorted elements of right to the result
			return append(result, right...)
		}
		if len(right) == 0 {
			//right is all done, append all remaining elems of left
			return append(result, left...)
		}
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else  {
			//first element of right is smaller, shud get added to the result
			result = append(result, right[0])
			right = right[1:]
		}
	}
	return result
}



QUICKSORT:
Algo: Choose a random element as the pivot..divide original array into 2 subarrays s.t. all elements 
in left subarray are less than pivot, and all in right as greater than pivot.
Recurse till only 1 element left in the array
Time: O(n log n )
Space : O(1)

//Foll uses go-routines, since each sub-array can work independently to get its pivot and sort

func qsort_pass(arr []int, done chan int) []int{
    if len(arr) < 2 {
        done <- len(arr)
        return arr
    }
    pivot := arr[0]
    i, j := 1, len(arr)-1
    for i != j {
        for arr[i] < pivot && i!=j{
            i++
        }
        for arr[j] >= pivot && i!=j{
            j--
        }
        if arr[i] > arr[j] {
            arr[i], arr[j] = arr[j], arr[i]
        }
    }
    if arr[j] >= pivot {
        j--
    }
    arr[0], arr[j] = arr[j], arr[0]
    done <- 1;

    go qsort_pass(arr[:j], done)
    go qsort_pass(arr[j+1:], done)
    return arr
}

func qsort(arr []int) []int {
    done := make(chan int)
    defer func() {
        close(done)
    }()

    go qsort_pass(arr[:], done)

    rslt := len(arr)
    //Wait for all goroutines to be done.
    //#goroutines == num elements
    for rslt > 0 {
        rslt -= <-done;
    }
    return arr
}


STRINGS/ARRAYS:
1. Check if a string is a palindrome 
Algo: Loop from element 0 to mid, check if alternate chars are the same
func isPalindrome(s string) bool {
	runes := []rune(s)
	numRunes := len(runes) - 1
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[numRunes-i] {
			return false
		}
	}
	return true
}

2. Check if a string can be a palindrome (only 1 char can have an odd no. of occurences, else cannot make a palindrome)
func checkPalindrome(strVal string) bool {
	valuesMap := make(map[string]bool)
	valList := strings.Split(strVal, "")
	// Initiate all alphabets in string to false first
	for _, val := range valList {
		valuesMap[val] = false
	}
	for _, val := range valList {
		value, _ := valuesMap[val]
		// for each occurence, toggle value stored for the alphabet in the map
		// i.e if it was false, make it true, and vice-versa
		// In the end if the value for an alphabet is false means it occurs even no. of times
		if value == false {
			valuesMap[val] = true
		} else {
			valuesMap[val] = false
		}
	}
	count := 0
	for _, val := range valuesMap {
		// If any value in Map is true, means that value occurs odd no. of times, 
		// hence is not a palindrome
		if val == true {
			count++
		}
	}
	//Check if more than one chars had an odd number of occurences
	if count > 1 {
		return false
	}
	return true		
}


LINKED LISTS
1. Check if a linked list is a palindrome
Algo: Traverse to mid, reverse the list from mid->end. Compare the values in each sublist
