package main

import "fmt"

func incrementer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}

}

func main() {
	nextInt := incrementer()
	fmt.Println(nextInt()) // i vale 1
	fmt.Println(nextInt()) //i vale 1 + 1 (uno es el valor de arriba )
	fmt.Println(nextInt()) // i vale 2 + 1 (dos viene de arriba)
	fmt.Println(nextInt()) // i vale 3 + 1

	//Sin la estructura definita i no alcanza el valor 4
}
