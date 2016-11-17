1. A child is running up the stairs. He can take 1 step at a time, or 2, or 3. How many ways are there for her to get the top.
Algo: To reach step n, we can reach step n-1, then take 1 step OR reach step n-2 and take 2 steps OR reach step n-3 and take 3 hops.
So total no. of ways = Sum ( stepsFor(n-1), stepsFor(n-2) , stepsFor(n-3))

Also, since we will compute say stepsFor(2) multiple times, lets store each such value, and re-use it once computed.

func stepsFor(n int, valMap map[int]int) int{
  if n < 0  {
    return 0
  }
  if n == 0 {
    return 1
  }
  val, ok := valMap[n]
  if !ok {
   result := stepsFor(n-1, valMap) + stepsFor(n-2, valMap) + stepsFor(n-3, valMap)
   valMap[n] = result
   return result
  } else {
    //Value has been computed already, just return that value
    return val
  }
}

2. Find if there is a magic index in an array of 'sorted, distinct' numbers
Algo: if A[5] = 3, then magic index cannot be in the left sub-array, bcoz A[4] should be 4, but that is not possible, since 
it is a sorted array, so the value has to be less than 3, so magiv index will be in right side. Find recursively.
It is like binary search, except the "comparison criteria" is different

func Magic(arr []int) int {
  return magic(arr, 0, len(arr)-1)
}
func magic(arr []int, start, end int)  int {
  if end <start {
    return -1
  }
  mid := (start + end )/2
  if arr[mid] < mid {
      //search right
      return magic(arr, mid+1, end)
  }
  return magic(arr, start, mid)
}
 
3. Power set: Get all subsets of a set (similar to permutations of a string problem)
Algo: For a set with n elements, there can be 2^ n subsets
if P2 = X, P3 (i.e. set with 3 elements) is p2 + all subsets of P2 with new element in P3 added
func getSubsets(arr []int, index int) [][]int {
  if len(arr) == index {
    // new everything
    allSubsets := [][]int{}
    allSubsets = append(allSubsets, []int{})
  }
  //
  allSubsets  = getSubsets(arr, index+1)
  //New item in say, P3
  item := arr[index]
  //Add item to all exising subsets
  for _, subset := range allSubsets {
    //Copy subset, and add item to this list
    newSubset := subset
    newSubset = append(newSUbset, item)
    moreSubsets = append(moreSubsets, newSubset)
  }
  //Add this list of new subsets to "allSubsets"
  
4. Towers of Hanoi S(tart)E(nd)M(iddle), SME, Print, MES
func towers(disk, start, end, middle int) {
    //Check base case
    if start >end {
      //error
    }
    towers(disk-1, start, middle, end )
    fmt.Println("Move disk %d from %d to %d", disk, start, end)
    towers(disk-1, middle, end, start)
    
}
 
5. Find number of valid parenthese combinations.
  Algo: As long as left remaining, add left and recurse.
  func printParen(left, right int, result string) {
    if left == right {
      fmt.Println(result)
    }
    if left <right {
      //invalid state
      return
    }
    if left > 0 {
      //still have left parens, recurse
      printParen(left-1, right, result + "(")
    }
    if right > 0 {
      //have right paren, recurse
      printParen(left, right-1, result + ")")
    }
  }

6. Paint fill. Given a screen, fill a pixel with say Green, and bleed outward till you hit a different color. 
  i.e. fill all red boxes with green start at red box(row, col), till you hit a non-red box

  func fill(screen [][]int, row, col int, origColor, newColor Color) {
    //check edge cases, row and col have not reached either end
    
    if screen[row][col] == origColor {
      //continue filling with new color
      fill(screen, row-1, col, origColor, newColor)
      fill(screen, row, col-1, origColor, newColor)
      fill(screen, row+1, col, origColor, newColor)
      fill(screen, row, col+1, origColor, newColor)
    }
  }

7. Get all possible denoms for a given amount
  func makeChange(amount int, denoms []int, index int, valArr [][]int) int{
    //if value present in array, get it
    if valArr[amount][index] > 0 {
      return valArr[amount][index]
    }
    ways := 0
    denom := denoms[index]
    for i := 0 ; denom  <=amount ;i++ {
      amountRemain = amount - denom *i
      ways += makeChange(amountRemaining, denoms, index+1, valArr)
    }
    valArr[amount][index] = ways
    return ways
  }
  
8. Box stacking
For each box, get the max height with it at the bottom. Save the computed height in a map for re-use
  
  func CreateStack(boxes []Box ) int {
    //Sort by height
    sort.Sort(boxes)
    //intitalise a map to store the already-calculated max height for a given box
    heightMap := make(map[int]int)
    for i := 0 ; i <len(boxes) ; i++ {
      height := createStack(boxes, i , heightMap)
      maxHeight = math.Max(maxHeight, height)
    }
    return maxHeight
  }
  func createStack(boxes []Box, index int, heightMap map[int]int) int {
    //check if maxHeight already computed for Box at this index
    // if yes, return that stored value
    //if not, loop through all boxes "above" this box, and find max height
    bottom := boxes[index]
    for i := index + 1 ; i < len(boxes) ; i++ {
      currBox := boxes[i]
      if canBeAbove(currBox, bottom) {
        height := createStack(boxes, i , heightMap)
        maxHeight = math.Max(maxHeight, height)
      }
      maxHeight += bottom.height
      heightMap[index] = maxHeight
      return maxHeight
    }
  }
  
 
  
  Func checkComplete(arr []int) bool {
	valueMap := make(map[int]string)
	Arr2 := []int{}
	//Traverse the array
  currIndex := 0
  For index := 0 ;index < len(arr)  ; index++   {
  currIndex += arr[currIndex]
  arr2 = append(arr2, arr[currIndex))
  }
  sort.Int(arr1)
  sort.Int(arr2)
  //Check if the array being created after visiting all the elements, 
  return arr1 == arr2
}



