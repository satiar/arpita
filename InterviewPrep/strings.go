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

10. Replace all occurences of whitespace with %20
func replaceWhitespace(str string) string {
	str = strings.Trim(str," ")
	//replace (string, old, new , n(if <0, then replace all occurences))
	//if n =2, replae first 2 occurences etc...
	return strings.Replace(str," ", "%20", 2)	
}

11. Find all permutations of a characters of a string (n!)
Algo: Example, 
Find all perm of abc:
We will find all perms of the given string based on all perms of its substring
To find all perms of abc, we need to find all perms of bc, and then add a to those perms
To find all perms of bc, we need to find all perms of c, and add b to those perms
To find all perms of c, well we know that is only c

Now we can find all perms of bc, insert c at every available space in b
bc --> bc, cb
Now we need to add a to all perms of bc
abc, bac, bca (insert a in bc) -- acb, cab, cba (insert a in cb)

func getPermutations(str string) []string {
	// base case, for one char, all perms are [char]
	if len(str) == 1 {
		return []string{str}
	}

	current := str[0:1] // current char
	remStr := str[1:] // remaining string

	perms := getPermutations(remStr) // get perms for remaining string

	allPerms := make([]string, 0) // array to hold all perms of the string based on perms of substring

	// for every perm in the perms of substring
	for _, perm := range perms {
	        // add current char at every possible position
		for i := 0; i <= len(perm); i++ {
			newPerm := insertAt(i, current, perm)
			allPerms = append(allPerms, newPerm)
		}
	}

	return allPerms
}

// Insert a char in a word
func insertAt(i int, char string, perm string) string {
	start := perm[0:i]
	end := perm[i:len(perm)]
	return start + char + end
}


12. Remove duplicates
Algo: maintain a list of not-duplicated values, by checking for existing values in a maintained Map
func removeDuplicates(a []int) []int {
	result := []int{}
	seen := map[int]int{}
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = val
		}
	}
	return result
}

13. Check if one string(C) is the result of interleaving of two other strings( A & B)  - NOT IMPLEMENTED !!!!!!!
Algo: Dynamic programming
Start with first char of C.
//pseudo-code
for len(c) > 0 {
	//first check length of c == len(a) + len(b)
	//then compare first char of c with first of a, if match, move both to next char
	// else check if matches b, 
	// else return false
	if c[0] == a[0] {
		c = c[1:]
		a = a[1:]
	} else if c[0] == b[0] }
		c = c[1:]
	b = b[1:]
	} else {
		return false
}
	
if c[0] == a[0] {
		recurse with 
	
	
14. Remove a char from a string
Algo : Use the built-in Replace method
func removeChar(s string, char string) string {
	return strings.Replace(str,char, "", 2)	
}
	
//!!!!!!!!!TODOS!!!!!!!!!!!!!!!!!!!!!!
15. Find the longest palindrome
	
16. Sort a list of strings based on length of the strings
type ByLength []string
func (s ByLength) Len() int {
    return len(s)
}
func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func order(s []string)[]string {
	sort.Sort(ByLength(s))
	return s
}
	
17. Check if two strings are permutations of each other
Algo 1 : split to get a list ->sort ->join to get a string ->both strings should be equal
Algo 2 : for each char in first string, do a map[char]++, 
	and for each char in second string, do a map[char]--....if map[char] <0,return 0
