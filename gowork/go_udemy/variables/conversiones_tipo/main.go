package main

import (
	"fmt"
	"log"
	"strconv"
)

var (
	entero16 int16  = 50
	entero32 int32  = 100
	op1      string = "100"
	num      int8   = 42
)

func main() {
	fmt.Println(entero16 + int16(entero32))
	op1_convertido, err := strconv.Atoi(op1)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(op1_convertido + op1_convertido)

	num_to_string := strconv.Itoa(int(num))

	fmt.Println(num_to_string + num_to_string)

}
