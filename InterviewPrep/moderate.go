Get max of 2 numbers, without using if-else or logical operators
func getMax(a, b int) int {
		C := a-b			//c is either negative or positive
		K := (c >> 31) & 0x1	//move c (a 32-bit int) 31 places to the right, to get the most significant bit in the lowest place. MSB is 1 for negative nos. Then And it with 1 to get only the lowest bit (which was originally the most significant bit) 
		Return a-k*c		// if k = 0 , then max is k which is true since a-b was +ve, else we get b as the max
}
