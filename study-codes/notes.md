- Producer/consumer model

```go
package main

import (
	"time"
	"math/rand"
	"fmt"
)

func main() {
	ch := make(chan string)
	timeout := time.After(600 * time.Millisecond)
	go producer(ch)
	for {
		select {
		case current := <-ch:
			fmt.Printf("Got message %s \n", current)
		case <-timeout:
			println("Time out")
			return
		default:
			println("The consumer is waiting")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func producer(ch chan string) {
	for {
		current := RandStringRunes(3)
		ch <- current
		fmt.Printf("Sent one task of %s \n", current)
		time.Sleep(50 * time.Millisecond)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
```

- WaitGroup

```go
package main

import "sync"

func main() {
	var wg sync.WaitGroup

	for _, word := range []string{"a", "b", "c"} {
		wg.Add(1)
		go func(w string) {
			defer wg.Done()
			println(w)
		}(word)
	}
	wg.Wait()
}
```

- sync/Mutex
```go
func main() {
	var lock = sync.Mutex{}
	var count int
	var wg sync.WaitGroup

	for i:=0; i<5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			count++
			println("Now the counter is ", count)
			lock.Unlock()
		}()
	}

	for i:=0; i<5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			count--
			println("Now the counter is ", count)
			lock.Unlock()
		}()
	}
	wg.Wait()
}
```

- channel as a simutaneous blocker 

```go
begin := make(chan interface{})
var wg sync.WaitGroup
for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(i int) {
        defer wg.Done()
        <-begin  
        fmt.Printf("%v has begun\n", i)
    }(i)
}

fmt.Println("Unblocking goroutines...")
close(begin)  
wg.Wait()
```
