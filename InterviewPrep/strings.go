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

3. Check if 2 strings are anagrams of each other
Algo : split bothe strings into array, then sort both arrays. Then join the individual arrays, and see if the result of joining 
is the same
func checkAnagram(first, second string) bool {
	one := strings.Split(first, "")
	two := strings.Split(second, "")
	sort.Strings(one)
	sort.Strings(two)
	//Make a string out of the list
	if strings.Join(one,"") == strings.Join(two, "") {
		return true
	}
	return false
}

4. Find the first unique character in a string
func findUnique(s string) string {
	m := make(map[rune]uint, len(s))
	for _, r := range s {
		m[r]++
	}
	for _, r := range s {
		if m[r] == 1 {
			return string(r)
		}
	}
	return ""
}

5. Reverse a string
Algo: Create a new list of runes, and assign input[i] to result[len-i+1]
func stringReverse(s string) string {
	
    o := make([]rune, len(s));
    i := len(o);
    for _, c := range s {
    	i--;
    	o[i] = c;
    }
    return string(o);


}

6. Check if a string contains only digits
Algo: use strconv.Atoi to see if get any errors. If it is a number, there will no error during conversion
func checkIfNumber(s string)bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}

7. Get duplicates in a string
Algo: Add values to a map if occurs a second time and also append to a buffer
func findDup(s string) string {
	//
	var buff bytes.Buffer
	valueMap := make(map[rune]bool)
	for _, val := range s {
		_, exists := valueMap[val]
		if exists {
			// This is a duplicate, since it already exists in the Map
			buff.WriteString(string(val))
		} else {
			// Add this value now, so we can track any future dups
			valueMap[val] = true
		}
	}
	return buff.String()
}

8. Find the number of vowels in a string
func findVowels(s string) int {
	runes := []rune(strings.ToLower(s))
	count := 0
	for _, val := range runes {
		switch string(val) {
			case "a", "e", "i", "o", "u":
				count++
				break;
			default:
		}
	}
	return count
}

9. Find number of occurences of a given char in a string
func findCount(s string, char string) int {
	runes := []rune(strings.ToLower(s))
	count := 0
	for _, val := range runes {
		if string(val) == char {
			count++
		}
	}
	return count
}
OR
//Use inbuilt method
func findCount(s string, char string) int {
	return strings.Count(strings.ToLower(s), char)
}

10. Replace all occurences of whitespace with %20func replaceWhitespace(str string) string {
	str = strings.Trim(str," ")
	//replace (string, old, new , n(if <0, then replace all occurences))
	//if n =2, replae first 2 occurences etc...
	return strings.Replace(str," ", "%20", 2)	
}
