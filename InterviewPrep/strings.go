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
		} else {
			// Add this value now, so we can track any future dups
			valueMap[val] = true
			buff.WriteString(string(val))
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
	
18. One-away. Check if one string is 1-away from another. One insert, one delete, or one replacement away
Algo: if len same, check for replacement.
	Loop through length of strings. Get char at ith position for both. If different and the bool foundDiff == true (which means already
	found a difference, then return false)
	if len1 == len2+1, check for insert
	Loop through each element of bigger string, and see if other string "contains" it. Increase count. Return true 
	when count == len(smaller string)
	
19. Compress a string 
E.g. aaabbcc is returned as a3b2c2
func compressAString(strVal string) string {
	strList := strings.Split(strVal, "")
	count := 0
	var buffer bytes.Buffer
	prev := strList[0]
	for _, val :=range strList {
		fmt.Println("val is", val)
		if val == prev {
			count++
		} else {
			fmt.Println(val, "val is not same as", prev)
			buffer.WriteString(prev+strconv.Itoa(count))
			prev = val
			count =1
		}
	}
	//Write out the last character's count
	buffer.WriteString(strList[len(strList)-1]+ strconv.Itoa(count))
	return buffer.String()
}
	
20. Rotate a matrix
func rotateMatrix(strVal [][]int) [3][3]int {
	//val is a 2d matrix
	result := [3][3]int{}


	for x, innerList := range strVal {
		for y, val := range innerList {
			result[y][x] = val
		}
	}
	return result
}
//in-place???
	
21. Return a matrix such that if any element is 0, all elements in that row and column are set to 0
func replaceRowColumnZeros(strVal [][]int) [][]int {
	result := strVal
	rowToMakeZero := []int{}
	columnToMakeZero := []int{}
	for x, innerList := range strVal {
		for y, val := range innerList {
			if val == 0 {
				rowToMakeZero = append(rowToMakeZero, x)
				columnToMakeZero = append(columnToMakeZero, y)
			}
		}
	}
	for _, val := range rowToMakeZero {
		for i:=0; i <3; i++ {
			result[val][i] = 0
		}
	}
	for _, val := range columnToMakeZero {
		for i:=0; i <3; i++ {
			result[i][val] = 0
		}
	}
	
	return result
	
}

22. Check if characters in a string are unique
Algo: create and array of size 128, and intialize it with zeros.
Loop through each char in the string, and set the array[char] = 1, if it is not 1, else return false (since if value is one for 
the given char, means this char has been traversed before, so not unique)
func checkUnique(value string) bool {
	exists := [2048]bool{}
	for index :=0 ;index <len(value); index++ {
		theChar := value[index]
		if exists[theChar] == false {
			exists[theChar] = true
		} else {
			return false
		}
	}
	return true
}

23. Given a sorted array of strings which is interspersed with empty strings, write a method to find the location of a given string
						
Example: find “ball” in [“at”, “”, “”, “”, “ball”, “”, “”, “car”, “”, “”, “dad”, “”, “”] will return 4 
	Example: find “ballcar” in [“at”, “”, “”, “”, “”, “ball”, “car”, “”, “”, “dad”, “”, “”] will return -1 
					
func findInListWithEmptyStrings(list []string, start, end int, element string) int {
	//confirm there is some value in this list
	for start <= end  {
		//clean up empty strings at the end, to save on time and space
		for start <= end && list[end] == "" {
			end--
		}
		if end < start {
			//no real string , only empty strings in this..
			return -1
		}
		mid := (start+end)/2
		if element == list[mid] {
			return mid
		}
		if element < list[mid] {
			//search in left half
			end = mid -1
		} else  {
			//search in right half
			start = mid + 1
		}
	}
	return -1
}
	
24. Check if s2 is a rotation of s1. eg is waterbottle is in erbootlewat ?
Algo : check edge cases, then check if s2 is in s1s1  erbootlewaterbootlewat

25. Given a string containing Roman numerals, print the decimal equivalent of it
	
package main

import (
    "fmt"
)
// I = 1
// V = 5
// X = 10
// L = 50
// C = 100

//Return the decimal equivalent of a roman numeral
// input = XV 10 +5 =15
// XXV =  
// IX = 9 , XI = 11
// VI  = 5+1
// IVX - invalid 

func main() {
    fmt.Println(getDecimal("XV"))
}




//input = XI = 11
//input = IX = 9
// input = XIV = 14

func getDecimal (s string) int {
    eMap := make(map[string]int)
    eMap["I"] = 1
    eMap["V"] = 5
    eMap["X"] = 10
    eMap["L"] = 50
    eMap["C"] = 100
    runes := []rune(s)
    result, ok := eMap[string(runes[len(runes)-1])]
    if ok {
        for i := len(runes) -2; i >= 0 ; i-- {
            //check if i am less than right char then subtract me, else add me
            curr, _ := eMap[string(runes[i])]
            prev, _ := eMap[string(runes[i+1])]
            if curr < prev {
                result -= curr
            } else {
                result += curr
            }
        }
    } else {
        return 0
        //fmt.Errorf("Invalid characters in provided input")
    }
    return result
        
}
26. Print FizzBuzz if no. divisible by 5 & 3, fizz if only 3, buzz if only 5. Print the number otherwise. For all n <100.
func FizzBuzz(amount int) string {  
    results := ""
    for i := 1; i <= amount; i++ {
        result := ""
        if i%3 == 0 { result += "Fizz" }
        if i%5 == 0 { result += "Buzz" }
        if result != "" {
            results += result + "\n"
            continue
        }
        results += fmt.Sprintf("%d\n", i)
    }
    return results
} 

