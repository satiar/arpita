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

Q. Find index of an element in a sorted, but rotated list
// Get the pivot (first element in list which is smaller than previous)
// If given element is > arr[0], then binary search in left sub-tree, else in right subtree
func getIndex(element int, elemList []int) int {
	if len(elemList) == 0 {
			//error
	}
	//find pivot
	var pivot int
	for index, val := range elemList {
		if index > 0 && val < elemList[index-1] {
			//found pivot
			pivot = index
			Break
		}
	}
	fmt.Println("pivot is at ", pivot)
	if elemList[0] < element {
		return binarySearch(elemList, 0, pivot-1, element)
	} 
		return binarySearch(elemList, pivot, len(elemList)-1, element)
}


func binarySearch(arr []int,  low,  high,  key int) int {
   if (high < low) {
       return -1
	}
   mid := (low + high)/2
   if (key == arr[mid]) {
       return mid
   }
   if (key > arr[mid]) {
       return binarySearch(arr, (mid + 1), high, key)
  }
   return binarySearch(arr, low, (mid -1), key)
}


Q. Given a matrix in which each row and each column is sorted, write a method to nd an element in it 
	//Start at top right
	// If element less than current move left, else move down
func findInSortedRowNCol(matrix [][]int, element int ) (int, int) {
	return find(0, len(matrix)-1, element, matrix)
}


func find(row, col, element int , matrix [][]int) (int, int) {
	fmt.Println(row, col)
	//boundary cases
	if row <0 || row > len(matrix)-1 || col < 0 || col > len(matrix[0])-1 {
		return -1, -1
	}
	currElement := matrix[row][col]
	fmt.Println(currElement)
	if element == currElement {
		return row, col
	}
	if element < currElement {
		return find(row, col-1, element, matrix)
	}
	if element > currElement {
		return find(row+1, col, element, matrix)
	}
	return -1, -1
}

Q. Find the rank of a given number in a stream of integers.
Store the integers in a BST. Algo then is :
func getRank(node Node, x int) int{
	if node.Data == x, return node.LeftSize()
	if x < node.Data , return getRank(node.Left, x)
	if x >node.Data, return node.LeftSize + 1 + getRank(node.Right, x)
}

Q. Peaks and valley - Given an array of integers, sort them in alteranting peaks(adjacent less than element)
and valley(adjacent more than element)

Algo : normal sort the array
start from index 1, jump 2 elements at a time
swap each a[i]th element with previous element


LINKED LISTS
1. Check if a linked list is a palindrome
Algo: Traverse to mid, reverse the list from mid->end. Compare the values in each sublist
