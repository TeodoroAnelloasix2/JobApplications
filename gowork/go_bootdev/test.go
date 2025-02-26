/*
Calcular Saldo
Necesitamos poder calcular el costo de lotes completos de mensajes de una sola vez.
Tarea
Usando las variables proporcionadas, escribe declaraciones condicionales para calcular y actualizar los valores.

	Primero, asigna finalCost al valor de bulkMessageCost.
	Si el usuario es premium, aplica discountRate a finalCost.
	    Por ejemplo, si discountRate es 0.10, el precio con descuento por mensaje será el 90% del precio original.
	Luego, verifica si el usuario tiene suficiente dinero en su accountBalance:
	    Si es así:
	        Procesa el pago restando finalCost de accountBalance.
	        Imprime purchaseSuccessMessage.
	    Si no tiene saldo suficiente, simplemente imprime insufficientFundMessage.
*/
package main

import "fmt"

func main() {
	var insufficientFundMessage string = "Purchase failed. Insufficient funds."
	var purchaseSuccessMessage string = "Purchase successful."
	var accountBalance float64 = 100.0
	var bulkMessageCost float64 = 75.0
	var isPremiumUser bool = true
	var discountRate float64 = 0.10
	var finalCost float64 = bulkMessageCost

	// don't edit above this line

	// ?

	// don't edit below this line
	if isPremiumUser {
		finalCost = bulkMessageCost - (bulkMessageCost * discountRate)

	}

	if accountBalance >= finalCost {
		accountBalance = accountBalance - finalCost
		fmt.Println(purchaseSuccessMessage)

	} else {
		fmt.Println(insufficientFundMessage)
	}
	fmt.Println("Account balance:", accountBalance)
}

/*
//package main

//import "fmt"

func printReports(intro, body, outro string) {
	// ?
	printCostReport(func(a string) int {return len(a) * 2}, intro )
	printCostReport(func(a string) int {return len(a)*3 }, body )
	printCostReport(func(a string) int {return len(a)*4}, outro)


}

// don't touch below this line

func main() {
	printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(Message: "%s" Cost: %v cents, message, cost)
	fmt.Println()

}*/
