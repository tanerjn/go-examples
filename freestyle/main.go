package main

import 	(
	"fmt"
	"time"
	"math/rand"
	"math"
	"math/cmplx"
)

var a, b  int = 1, 2


var (
	isTrue   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	inn int
	f float64
	bl bool
	s string
)


func main(){
	var i, j, k = true, false, "no!"
	fmt.Println("flan, , 世界")
	fmt.Println("Time : ", time.Now())
	
	// Get the density of each number
	density := testRandDensity()
	// Print the density of each number
	for num, count := range density {
		fmt.Printf("Number %d: %d hits\n", num, count)
	}
	n, err := fmt.Printf("SQ: %f:\n", math.Sqrt(121))
	if err != nil{
		fmt.Printf("Error getting square: %v\n", n)
		return
	}
	fmt.Println(math.Pi)
	fmt.Println("Sum:", add(math.Sqrt(121), math.Sqrt(225)))
	fmt.Println(swap("good", "all"))
	fmt.Println(split(57))
	fmt.Println(a, b, i, j, k)
	fmt.Printf("Type: %T, Value: %v\n", isTrue, isTrue)
	fmt.Printf("Type: %T, Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T, Value: %v\n", z,z)
	fmt.Printf("%v %v %v %q\n", inn, f, bl, s)

	var ip, jp int = 4, 2
	var f float64 = math.Sqrt(float64(ip*ip + jp*jp))
	var g float64 = math.Sqrt(float64(24))
	var u uint = uint(f)
	var u2 uint = uint(g)
	fmt.Println(u)
	fmt.Println(u2)
	
	j2 := 30.1i + 0.12
	fmt.Printf("j2 is of type %T\n", j2)
	const pi=3.14
	fmt.Println("Pi:", pi, "Num")
	const right = true
	fmt.Println("Right:", right)
	fmt.Println(sqrt(-1))
	

}

func pow(x, n, lim float64) float64 {
	if v:= math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func sqrt(x float64) string { 
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func add(x, y float64) float64 {
	return x + y
}

func swap(x, y string) (string, string) {
	return y,x
}

func split(a int)(x,y int) {
	x = a * 4 / 9
	y = a -x
	return
}




func testRandDensity() map[int]int {
	// Seed the random number generator to ensure different results on each run
	rand.Seed(time.Now().UnixNano())
	// Initialize a map to store the frequency of each number from 0 to 100
	hitCount := make(map[int]int)
	// Run the random number generator 100 times
	for i := 0; i < 1000000; i++ {
		// Generate a random number between 0 and 100
		num := rand.Intn(3) // 101 because Intn(n) generates a number between 0 and n-1
		// Increment the count for the generated number
		hitCount[num]++
	}
	return hitCount
}


