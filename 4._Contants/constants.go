package contants

import "fmt"

const fullname string = "Harcharan Singh"

func main() {
	const name string = "Harcharan"
	const lastName = "Singh"

	// Not allowed to change the value of a constant after it has been declared
	// name = "Simran" // This will cause a compile-time error

	fmt.Println(name)
	fmt.Println(lastName)
	fmt.Println(fullname)
	
	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(port, host)
}