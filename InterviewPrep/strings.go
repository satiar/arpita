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
