package main

import 	"fmt"

type Number interface { 
	int64 | float64
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first" : 34,
		"second" : 45, 
	}

	floats := map[string]float64{
		"first" : 34.34,
		"second" : 44.44,
	}
	fmt.Printf("Non-generics sums: %v and %v \n", SumInts(ints), SumFloats(floats))
	fmt.Printf("Generic Sums: %v and %v\n",SumIntsOrFloats(ints),  SumIntsOrFloats(floats))
	fmt.Printf("Non-generics sums(with interface): %v and %v \n", SumNumbers(ints), SumNumbers(floats))
}

//SumInts add the values of m.
func SumInts(m map[string] int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}


func SumFloats(m map[string] float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}


//Sum calculation of Ints or Floats hte values of map m. Supports both int64 and float64
func SumIntsOrFloats[K comparable , V int64 | float64] (m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
	
func SumNumbers[K comparable, V Number](m map[K]V) V{
	var s V
	for _, v := range m {
		s += v
	}
	return s
}












