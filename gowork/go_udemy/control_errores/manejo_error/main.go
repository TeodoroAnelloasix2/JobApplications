package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	str := "123a"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(num)

	res, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(res)
}
func divide(dividendo, divisor int) (int, error) {

	if divisor == 0 {
		return 0, errors.New("No se puede dividir por cero") //tambien sirve fmt.Errorf
	}
	return dividendo / divisor, nil
}
