package main

import (
	"fmt"
	"runtime"
	"time"
)

const (
	egne rune = 'ñ'
)

func main() {

	sis_op := runtime.GOOS //switch op:=runtime.GOOS; op {case ......}

	switch sis_op {
	case "windows":
		fmt.Println("Go corre en --> Windows.....(que mal!)")
	case "linux":
		fmt.Println("Go corre en --> linux.....(ME llenas de orgullo muchacho !)")
	case "darwin":
		fmt.Println("Seguro que tu cartera siempre està llena de billetes")
	default:
		fmt.Println("Go corre en --> Otro")
	}

	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("Es de ma" + string(egne) + "ana")
	case t.Hour() < 17:
		fmt.Println("Es de tarde")
	default:
		fmt.Println("Es de noche")
	}

}
