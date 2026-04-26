// What is a channel in Go?

// A channel is a pipe for communication between goroutines.

package main

import "fmt"

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch) // important
}

func consumer(ch chan int) {
	for num := range ch {
		fmt.Println("Received:", num)
	}
}

func main() {
	ch := make(chan int)

	go producer(ch)

	consumer(ch)
}


// WE CALL producer() ONCE, BUT IT SEEMS TO RUN 5 TIMES. WHY?

/*
	Good question — there are two different ideas here: what a channel is, and why code seems to “run multiple times.”

	What is a channel in Go?
	A channel is a pipe for communication between goroutines.

	Think of it like a conveyor belt:

	One goroutine puts data onto the belt (producer)
	Another takes data off (consumer)
	ch <- i     // send into channel
	num := <-ch // receive from channel

	This is Go’s “share memory by communicating” idea, instead of many threads mutating shared variables.

	Your code
	func producer(ch chan int) {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}

	You call producer() only once:

	go producer(ch)

	Correct.

	Then why does it “run” 5 times?

	It isn't being called 5 times.

	This loop inside the function runs 5 times:

	for i := 1; i <= 5; i++ {
		ch <- i
	}

	Single function call.

	Inside that one call:

	send 1
	send 2
	send 3
	send 4
	send 5

	So one call, five sends.

	Same for consumer

	You call consumer once:

	consumer(ch)

	But inside:

	for num := range ch {
		fmt.Println("Received:", num)
	}

	This loop keeps receiving until the channel is closed.

	Equivalent idea:

	receive 1 -> print
	receive 2 -> print
	receive 3 -> print
	...

	Again:

	one function call
	multiple loop iterations
	Why close(ch) matters

	Without:

	close(ch)

	this loop:

	for num := range ch

	would wait forever for more data.

	close(ch) tells consumer:

	"No more values are coming."

	Then the loop exits.

	Producer/Consumer flow
	Producer sends: 1 ---> channel ---> Consumer receives 1
	Producer sends: 2 ---> channel ---> Consumer receives 2
	Producer sends: 3 ---> channel ---> Consumer receives 3
	...
	close(channel)
	Consumer stops
	Another subtle thing happening: synchronization

	This line:

	ch <- i

	(on an unbuffered channel)

	blocks until someone receives it.

	And this:

	num := <-ch

	blocks until someone sends.

	They synchronize each other.

	That’s why channels are not just "queues"; they can coordinate execution.

	If you remove go

	Try:

	producer(ch)
	consumer(ch)

	Deadlock.

	Why?

	Because producer tries:

	ch <- 1

	and waits for a receiver...

	but consumer hasn’t started yet.

	That often makes channels “click” conceptually.

*/