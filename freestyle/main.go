package main

import 	(
	"fmt"
	"time"
	"math/rand"
	"math"
	"math/cmplx"
	"runtime"
	"reflect"
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

type ver struct{
	X int 
	Y int 
}

var (
	v1 = ver{5, 5}
	v2 = ver{X: 10, Y:10}
	v3 = ver{}
	p = &ver{1, 2}
)


type MyStruct struct {
	Array1 [5]int
	Array2 [5]int
}

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
	fmt.Println(pow(3,2,10), pow(3,3,20))
	fmt.Println(sqrt(2))

	fmt.Print("running on ")
	switch os := runtime.GOOS; os {
	case "darwin": 
		fmt.Println("OS X")
	case "linux":	
		fmt.Println("Linux")
	default: 
		fmt.Printf("%s\n", os)
	}

	fmt.Println("Saturday is:")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2: 
		fmt.Println("In two days")
	default:
		fmt.Println("Not soon")
		fmt.Println("Today:", today)
	}
	t := time.Now()
	switch {
	case t.Hour() < 12: 
		fmt.Println("Good morning")
	case t.Hour() < 17: 
		fmt.Println("Good afternoon")
	default: 
		fmt.Println("Good evening")
	}

	for i:=0; i<10; i++ {
		fmt.Println(i)
	}

	operand1, operand2 := 100, 121
	ptr := &operand1
	fmt.Println("Val ptr:", *ptr) //100
	fmt.Println("Val operand1:", operand1) //100
	*ptr= 9
	fmt.Println("Val operand1", operand1) //9
	ptr =  &operand2
	*ptr = *ptr/11
	fmt.Println("Val operand2:", operand2) //11
	fmt.Println("Val operand1:", operand1) //9
	*ptr = *ptr*5
	fmt.Println("Val operand2:", operand2) //55	
	ptr2 := &operand1
	*ptr2 = *ptr2+5
	fmt.Println("Val operand1:", operand1) //14
	ptr = &operand1
	fmt.Println("Val operand1:", *ptr) //14
	fmt.Println("Val operand1:", *ptr2) //14	
	ptr = &operand1
	ptr2 = &operand2
	fmt.Println("Val operand1:", *ptr) //14
	fmt.Println("Val operand1:", *ptr2+1) //56	
	*ptr2 = *ptr2+5
	*ptr = operand2 + operand1 // 60 + 14 
	fmt.Println("Val operand1:", operand2) //56	
	fmt.Println("Val operand1:", operand1) //74		
	*ptr2 = *ptr + 5
	fmt.Println("Val operand1:", operand2) //79	

	v := ver{1,2}
	p := &v 
	p.X = 1e9
	p.Y = 2e9
	fmt.Println(v.X, v.Y)
	v.X = 101
	v.Y = 102
	v1.X = 1001
	v1.Y = 1002
	v3.X = 10001
	v3.Y = 10002
	fmt.Println(v.X + v.Y)	
	fmt.Println(v1, v2, v3, *p)	
	var a [2]string
	a[1] = "What"
	a[0] = "What"
	fmt.Println(a[1], a[0])
	primes := [3]int{1,2,3}
	fmt.Println(primes[:3])

	names := [3]string {
		"mel",
		"del",
		"fel",
	}
	fmt.Println(names)
	n1 := names[0:2]
	n2 := names[1:3]
	fmt.Println(n1, n2)
	n2[0] = "pel"
	fmt.Println(n1, n2)
	fmt.Println(names)	
	ar1 := []int{1,2,3,3,2,1}
	ar2 := []bool{true, false, true, false, true, false}
	fmt.Println(ar1)
	fmt.Println(ar2)

	myStruct := MyStruct{
		Array1: [5]int{1, 2, 3, 4, 5},
		Array2: [5]int{6, 7, 8, 9, 0},
	}

	prArray := [5]int{6,7,8,9,0}

	if myStruct.matchesAnyArray(prArray) {
		fmt.Println("array matches")
	} else {
		fmt.Println("no match")
	}
}

func arraysMatch(arr1, arr2 [5]int) bool {
	return reflect.DeepEqual(arr1, arr2)
}

func (s MyStruct) matchesAnyArray(arr [5]int) bool {
	return arraysMatch(s.Array1, arr) || arraysMatch(s.Array2, arr)
}




func pow(x, n, lim float64) float64 {
	if v:= math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g/n", v, lim)
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


