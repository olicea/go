package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	// calculate a sum
	var x int = 5
	var y int = 6
	fmt.Println("sum ", x+y)

	//conditional print
	if x > 1 {
		print("x is greater than 1")
	}

	//define an array
	var arr [5]int
	fmt.Println("arr", arr)

	var arr2 = [2]string{"hello", "world"}
	fmt.Println("arr2", arr2)

	//slices
	a := []int{1, 2, 3, 4, 5}
	fmt.Println("slice a", a)
	a = append(a, 6)
	fmt.Println("slice a (After append)", a)

	//map
	m := make(map[string]int)
	m["key1"] = 1
	m["key2"] = 2
	fmt.Println("map m", m)

	delete(m, "key1")
	fmt.Println("map m after delete", m)

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println("loop i", i)
	}

	// while loop
	j := 0
	for j < 5 {
		fmt.Println("loop j", j)
		j++
	}

	//loop over array
	for index, value := range a {
		fmt.Println("index", index, "value", value)
	}

	// loop over map
	for key, value := range m {
		fmt.Println("key", key, "value", value)
	}

	//calling func
	fmt.Println("sum", sum(1, 2))

	//calling func with error
	result, err := sqrt(-4)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("result", result)
	}

	p := person{name: "Octavio", age: 2}
	fmt.Println("person", p)
	fmt.Println("person age", p.age)

	num := 0
	fmt.Println("num", num)
	increment(&num)
	fmt.Println("num", num)

}

//defining a struct
type person struct {
	name string
	age  int
}

// define a function
func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	// if the argument is invalid return error
	if x < 0 {
		return 0, errors.New("x is negative")
	}
	// otherwise calculate and return
	return math.Sqrt(x), nil
}

//defining a function with return values
func increment(x *int) {
	*x++
}
