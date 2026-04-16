package switchstatement

import (
	"fmt"
	"time"
)

func main() {
	// Simple Switch Statement
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Invalid day")
	}

	// Multiple condition switch statement
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	
	// Type Switch Statement
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("its an integer", t)
		case string:
			fmt.Println("its a string")
		default:
			fmt.Printf("unknown type %T\n", t)
		}
	}

	whoAmI(42)
	whoAmI("Hello")
	whoAmI(3.14)
}
