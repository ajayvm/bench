package main

import (
	"fmt"
)

type pt struct {
	X int
	Y int
}

// Calculate returns x + 2.
func Calculate(x int) (result int) {

	return x + 2
}

func main() {
	a := 3
	var p *int
	p = &a
	origin := pt{1, 1}
	fmt.Println("hello world 1", a, p, origin, origin.Y)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)

	var m = make(map[string]int)

	m["Ajay"] = 5
	m["Antu"] = 10

	fmt.Println(m)

	/*
		for i := 0; i < 10; i++ {
			fmt.Print(i)
		}
		fmt.Println()
		fmt.Println(runtime.GOARCH)
		fmt.Println(runtime.GOOS)

		bs := []byte("Hello")
		fmt.Println("byte string is is the byte string format of the string ", bs)
		fmt.Printf("%v \n", str[0])

		var x [5]int
		x[2] = 34
		fmt.Println(x)

		y := make([]int, 5)
		y[2] = 4
		fmt.Println(y, time.Now())
	*/
}
