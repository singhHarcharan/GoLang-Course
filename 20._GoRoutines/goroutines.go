package main
import "fmt"

func runTask(id int) {
	fmt.Printf("Running task %d\n", id)
}

func main() {
	for i:=0; i<5; i++ {
		go runTask(i)
	}

}