- Use Done and Task double channels and multiplexer

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

- Use Waitgroup

```go
package main

import "sync"

func main() {

	var wg sync.WaitGroup
	for i := 1; i < 16; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			switch {
			case i%15 == 0:
				println("fizzbuzz", i)
			case i%3 == 0:
				println("fizz", i)
			case i%5 == 0:
				println("buzz", i)
			}
		}(i)
	}
	wg.Wait()
}
```

- Use broadcaster/closer
```go
package main

func main() {
	var tasks = make(chan int)

	for i := 1; i < 16; i++ {
		go printInt(tasks)
		tasks <- i
	}
	close(tasks)
}

func printInt(tasks chan int) {
	i := <-tasks
	switch  {
	case i%15 == 0:
		println("fizzbuzz", i)
	case i%3 == 0:
		println("fizz", i)
	case i%5 == 0:
		println("buzz", i)
	}
}
```
