package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
func main() {

}

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

func main() {
	fmt.Println(math.Pi)
}
*/

func add(x int, y int) int {
	return x + y
}

func main() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100)
	y := rand.Intn(100)
	fmt.Println("My number is", x)
	fmt.Println("My number is", y)
	fmt.Println(add(x, y))
}
