/*
Assignment
After submitting our last snippet of code for review, a more experienced gopher told us to use a type switch instead of successive assertions. Let's make that improvement!
Implement the getExpenseReport function using a type switch.

	    If the expense is an email then it should return the email's toAddress and the cost of the email.
	    If the expense is an sms then it should return the sms's toPhoneNumber and its cost.
	    If the expense has any other underlying type, just return an empty string and 0.0 for the cost.

		func printNumericValue(num interface{}) {
		switch v := num.(type) {
		case int:
			fmt.Printf("%T\n", v)
		case string:
			fmt.Printf("%T\n", v)
		default:
			fmt.Printf("%T\n", v)
		}
*/
package main

func getExpenseReport(e expense) (string, float64) {

	//type dento de () es una palabra clave!
	switch kind := e.(type) {
	case email:
		return kind.toAddress, e.cost()
	case sms:
		return kind.toPhoneNumber, e.cost()
	default:
		return "", .0
	}

}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}
