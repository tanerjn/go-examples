package main

import 	(
	"fmt"
	"time"
	"math/rand"
	"math"
	"math/cmplx"
	"runtime"
	"reflect"
	"strings"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/tree"
	"io"
	"sync"
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
	fmt.Println(len(prArray[0:2]))

	if myStruct.matchesAnyArray(prArray) {
		fmt.Println("array matches")
	} else {
		fmt.Println("no match")
	}

	fmt.Printf("Type: %T Value: %v\n", isFalse, isFalse)
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", complx, complx)
	const Truth = true
	fmt.Printf("Truth type %T", Truth)
	fmt.Printf("\nvalue %v\n", Big)
	fmt.Println(needInt(Big))
	sl1 := make([]int, 5) //len(a) = 5
	sl2 := make([]int, 4, 5) //len(b) = 5, cap(b) = 5
	printSlice("sl1", sl1) 
	printSlice("sl2", sl2)	

	board := [][]string {
		[]string{"","",""},
		[]string{"","",""},
		[]string{"","",""},
	}

	board[0][0] = "X"
	board[0][1] = "0"
	board[0][2] = "0"
	board[1][0] = "X"
	board[1][1] = "X"
	board[1][2] = "0"
	board[2][0] = "X"
	board[2][1] = "0"
	board[2][2] = "X"

	for i:=0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var s[]int
	s = append(s, 1)
	printSliceInt(s)
	s = append(s, 1, 2, 3)
	printSliceInt(s)

	for i, v := range powa{
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for _, value := range powa {
		fmt.Printf("%d\n", value)
	}

	for i, _ := range powa {
		fmt.Printf("%d\n", i)
	}

	pic.Show(Pic)

	m = make(map[string]Coordinates)
	m["Bell Bell"] = Coordinates{
		10.1231, -12.123,
	}

	m["Boll Boll"] = Coordinates{
		12.1231, -12.123,
	}

	m["Boll Boll Boll"] = Coordinates{
		10.1231, -12.123,
	}

	fmt.Println(m["Bell Bell"])
	fmt.Println(m["Boll Boll"])
	fmt.Println(s1["bell"], s1["dull"])	
	fmt.Println(s1)	
	delete(s1, "dull")
	fmt.Println(s1["bell"], s1["dull"])	
	fmt.Println(s1)	
	fmt.Println(s2)	


	counter := 0
	for key, value := range s2 {
		counter++
		if counter == 2 {
			fmt.Printf("Second el.: Key: %s, Value: %d\n", key, value)
			break
		}
	}


	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
		pos(i),
		neg(-2*i),
			)
	}

	rec := Receiver{3, 4}
	fmt.Println("rec1 Scale+Hyp:", rec.Hyp())
	rec2 := Receiver{3, 4}
	rec2.Scale(10)
	fmt.Println("rec2 Scale+Hyp:", rec2.Hyp())
	
	tcp := TCP{}
	udp := UDP{}
	quic := &QUIC{"QUIC"}

	EstablishConnection(tcp)
	EstablishConnection(udp)
	EstablishConnection(quic)

	var customConnector interface{}
	describe(customConnector)
	customConnector = 13	
	describe(customConnector)
	customConnector = "tring"	
	describe(customConnector)

	var iface interface{} = "iface"
	s1 := iface.(string)
	fmt.Println(s1)

	s1, ok := iface.(string)
	fmt.Println(s1, ok)
	//f1 := iface.(float64) will panic atinterface conversion 
	//fmt.Println(f1)
	fl, ok := iface.(float64)
	fmt.Println(fl, ok)

	if err := printError(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	r :=  strings.NewReader("Reader New")
	b := make([]byte, 12)
	for { 
	n, err := r.Read(b)
	fmt.Printf("n = %v err = %v b = %v\n" , n, err, b)
	fmt.Printf("b[:n] = %q\n", b[:n])
	if err == io.EOF { 
		break
	}
	}

	var wgr sync.WaitGroup
	wgr.Add(1) //Increment WaitGroup counter by 1 for the gorouitine
	go say("world", &wgr)
	wgr.Wait()
	say("hello", nil) // call nil instead of wg bec it's not needed

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("first chval:", <-ch)
	fmt.Println("second chval:", <-ch)

	ch2 := make(chan int, 5)
	go chancalc(cap(ch2), ch2)
	for i := range ch2 {
		fmt.Println(i)
	}
	
	ch3 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-ch3)
		}
		quit <- 0
	}()
	chancalc2(ch3, quit)

	// Tree 1	
	tree1 := &Tree{Value: 2}
	Insert(tree1, 1)
	Insert(tree1, 3)
	fmt.Println("Tree 1 structure")
	PrintTree(tree1, 0, "")


	//Tree 2
	tree2 := &Tree{Value: 2}
	Insert(tree2, 1)
	Insert(tree2, 3)
	fmt.Println("Tree 2 structure")
	PrintTree(tree2, 0, "")

	//Tree 3
	tree3 := &Tree{Value:4}
	Insert(tree3, 1)
	Insert(tree3, 8)
	Insert(tree3, 4)
	Insert(tree3, 5)
	fmt.Println("Tree 3 structure")
	PrintTree(tree3, 0, " ")

	//Tree 6
	tree6 := &Tree{Value:15}
	Insert(tree6, 90)
	Insert(tree6, 80)
	Insert(tree6, 70)
	Insert(tree6, 100)
	Insert(tree6, 40)
	Insert(tree6, 145)
	Insert(tree6, 70)
	Insert(tree6, 105)
	Insert(tree6, 4)
	Insert(tree6, 4)
	Insert(tree6, 115)
	Insert(tree6, 8)
	Insert(tree6, 7)
	Insert(tree6, 120)
	Insert(tree6, 125)
	fmt.Println("Tree 6 strutcture")
	PrintTree(tree6, 0, " ")

	fmt.Printf("Tree t1 and Tree t2 are the same: %v\n", Same(tree1, tree2)) // Should print: true
	fmt.Printf("Tree t1 and Tree t3 are the same: %v\n", Same(tree1, tree3)) // Should print: true

	c := SafeCounter{v: make(map[string]int)}
	for i :=0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}


type SafeCounter struct { 
	mu sync.Mutex
	v map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// one goroutine can access the map c.v at a time
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only onone goroutine can access the map
	defer c.mu.Unlock()
	return c.v[key]
}

// Insert a new value into the tree
func Insert(root *Tree, value int) *Tree {
    if root == nil {
        return &Tree{Value: value}
    }
    if value < root.Value {
        root.Left = Insert(root.Left, value)
    } else {
        root.Right = Insert(root.Right, value)
    }
    return root
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
   var inorder func(*tree.Tree)
    inorder = func(t *tree.Tree) {
        if t == nil {
            return
        }
        inorder(t.Left)
		fmt.Printf("t left: %v", t.Left)
        ch <- t.Value
		fmt.Printf("t.Value: %v", t.Value)
        inorder(t.Right)
	fmt.Println(t.Right)
		fmt.Printf("t.right: %v", t.Right)
    }
    go func() {
        inorder(t)
        close(ch)
    }()

}

func PrintTree(root *Tree, level int, prefix string) {
    if root == nil {
        return
    }
    // Print the current node
    fmt.Printf("%s%s%d\n", strings.Repeat("  ", level), prefix, root.Value)
    // Print the right subtree
    if root.Right != nil {
        PrintTree(root.Right, level+1, "R-- ")
    } else {
        fmt.Printf("%s%s\n", strings.Repeat("  ", level+1), "R-- nil")
    }
    // Print the left subtree
    if root.Left != nil {
        PrintTree(root.Left, level+1, "L-- ")
    } else {
        fmt.Printf("%s%s\n", strings.Repeat("  ", level+1), "L-- nil")
    }
}

// Same checks if two trees t1 and t2 are identical
func Same(t1, t2 *Tree) bool {
    var sameTrees func(*Tree, *Tree) bool
    sameTrees = func(t1, t2 *Tree) bool {
        if t1 == nil && t2 == nil {
            return true
        }
        if t1 == nil || t2 == nil {
            return false
        }
        return t1.Value == t2.Value && sameTrees(t1.Left, t2.Left) && sameTrees(t1.Right, t2.Right)
    }
    return sameTrees(t1, t2)
}


type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func chancalc(n int, c chan int) {
	x, y := 1, 10
	for i:= 0; i<n ; i++ {
	c <- x
	x, y = y, x*y
	}
	close(c) //close a channel so no more values sent
}

func chancalc2(c, quit chan int) { 
	x, y := 0, 1
	for {
		select {   //select will block one of cases run then it executes the case 
		case c <-x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default: 
			fmt.Println("default")
		}
	}
}
			
func say(s string, wg *sync.WaitGroup) {
	// if WaitGroup is not nil, then defer the Done
	if wg != nil {
	defer wg.Done() //Signal when goroutine is done
	}

	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func printError() error {
	return &MyError {
		time.Now(),
		"ERR",
	}
}

type Connector interface { 
	Connect() string
}

type TCP struct{}
type UDP struct{}
type QUIC struct{
	s string
}

func (t TCP) Connect() string{
	return "TCP:ConnEst"
}

func (u UDP) Connect() string{
	return "UDP:ConnEst"
}

func (q *QUIC) Connect() string{
	if q == nil {
	fmt.Println("<nil>")
	return("Method has nil input")
	}
	return(q.s+":ConnEst")
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func EstablishConnection(c Connector){
	fmt.Println(c.Connect())
}


func (v *Receiver) Scale(f float64) {
	v.x = v.x * f 
	v.y = v.y * f
}

type Receiver struct {
	x, y float64
}

func (v Receiver) Hyp() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}


func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

type Coordinates struct {
	Lat, Long float64
}

var m map[string]Coordinates
var s1 = map[string]Coordinates{
	"bell": {12.12, 12.12},
	"dull": {12.13, 12.13},
}

var s2 = map[string]Coordinates{
	"bell": {12.12, 12.12},
	"dull dull": {12.13, 12.13},
	"pull": {10.13, 10.13},
}


func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)    // Create a slice with 'dy' rows
	for y := 0; y < dy; y++ {
		img[y] = make([]uint8, dx) // Each row has 'dx' columns
		for x := 0; x < dx; x++ {
			img[y][x] = uint8(x ^ y) // Replace this with any function to generate the image
		}
	}
	return img
}

var powa = []int{1,2,4,8,16,32,64, 128} 

func printSliceInt(s []int) {
	fmt.Printf("Slice Int: %d len=%d cap=%d %v\n", s, len(s), cap(s), s)
	}

func printSlice(s string, x []int) {
	fmt.Printf("Slice String: %s len=%d cap=%d %v\n", s, len(x), cap(x), x)
	}


const (
	Big = 1 << 8
	Small = Big >> 98
)

var  (
	isFalse bool = false
	maxInt  uint64 = 1<<64 -1
	complx complex128 = cmplx.Sqrt(-5 + 12i)
)

func needInt(x int) int { return x*10 + 1}

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

func Sqrt(x float64) (float64, error) {
	return 0, nil
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


