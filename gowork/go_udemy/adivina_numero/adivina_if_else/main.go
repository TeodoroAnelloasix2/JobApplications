package main

import (
	"fmt"
	"time"
)

const (
	egne rune = 'ñ'
)

func main() {
	t := time.Now()
	hora := t.Hour()

	fmt.Println(t.Hour())

	if hora < 12 {
		fmt.Println("Es de ma" + string(egne) + "ana")
	} else if hora > 12 && hora < 18 {
		fmt.Println("Es de tarde")
	} else {
		fmt.Println("Es de noche")
	}

}
