package reverse

import "strconv"

// INt returns the decimal reversal of the integer i. 
func Int(i int) int {
	i, _ = strconv.Atoi(String(strconv.Itoa(i)))
	return i
}


