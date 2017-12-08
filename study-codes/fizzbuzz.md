- Use sync.WaitGroup

```go

package main

func main() {
	var tasks = make(chan int)
	var done = make(chan bool)

	for i := 1; i < 16; i++ {
		go printInt(tasks, done)
		tasks <- i
	}
	for i := 1; i < 16; i++ {
		done <- true
	}
}

func printInt(tasks <-chan int, done <-chan bool) {
	for {
		select {
		case <-done:
			return
		case i := <-tasks:
			switch {
			case i%15 == 0:
				println("fizzbuzz", i)
			case i%3 == 0:
				println("fizz", i)
			case i%5 == 0:
				println("buzz", i)
			}
		}
	}
}


```
