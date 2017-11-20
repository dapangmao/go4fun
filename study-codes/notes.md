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


- Mod for a [sample](https://github.com/dapangmao/go4fun/blob/master/study-codes/go-in-practice.md#chapter3---scratch---closingstuffgo) 

```go
func main() {
	ch := make(chan bool)
	timeout := time.After(600 * time.Millisecond)
	go send(ch)
	for {
		select {
		case <-ch:
			println("Got message.")
		case <-timeout:
			println("Time out")
			return
		default:
			println("*yawn*")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func send(ch chan bool) {
	for {
		time.Sleep(120 * time.Millisecond)
		ch <- true
		println("Sent a message")
	}
}
```

- A sync worker model like Celery
```go
package main

import (
	"fmt"
	"net/http"
	"time"
	"strings"
	"log"
)

func worker(name string, results chan<- string) {
	for i := 1; i <= 10; i++ {
		var current string
		if i < 10 {
			current = fmt.Sprintf("The progress of %s is at %v percnet", name, i*10)
			time.Sleep(time.Second)
		} else {
			current = "Work done"
		}
		results <- current
		if i == 10 {
			return
		}
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request, data map[string][]string) {
	queryValues := r.URL.Query()
	name := queryValues.Get("name")

	if val, ok := data[name]; ok {
		w.Write([]byte(strings.Join(val, "\n")))
		return
	}
	data[name] = []string{}
	res := make(chan string)
	go worker(name, res)
lp:
	for {
		select {
		case current := <-res:
			data[name] = append(data[name], current)
			if current == "Work done" {
				break lp
			}
		case <-time.After(time.Second * 10):
			log.Fatal("Something wrong after 10 seconds")
			delete(data, "name")
			return
		default:
			println("Waiting for the results to come")
			time.Sleep(time.Millisecond * 200)
		}
	}
	w.Write([]byte(strings.Join(data[name], "\n")))
}

func main() {
	data := make(map[string][]string)
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		HomeHandler(w, r, data)
	})
	http.ListenAndServe(":8088", nil)
}
```

- Use regex to extract strings 
```go
	re := regexp.MustCompile("\\$\\{(.*?)\\}")
	match := re.FindAllStringSubmatch("git commit -m '${abc}' ${asdfsad } safs", -1)
	fmt.Println(match)
	// [[${abc} abc] [${asdfsad } asdfsad ]]
```


- A crawler
```go
package main

import (
	"regexp"
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"time"
)

func main() {
	var m = sync.Map{}
	tasks := make(chan string, 10000)
	tasks <- "https://xkcd.com"
	timer := time.After(time.Minute)
	lp:
	for {
		select {
		case <- timer:
			_, ok := <- tasks
			if !ok {
				println("\nThe task queue is empty\n")
			} else {
				println("\nstill some tasks are running\n")
			}
			break lp
		case <- time.Tick(time.Second):
			current, ok := <- tasks
			if ok {
				fmt.Printf("%s has been fired up\n", current)
				go fetcher(current, tasks, &m)
			}
		}
	}
	m.Range(func(k, v interface{}) bool {
		println(k.(string), "------>", v.(string))
		return true
	})
}


func fetcher(current string, tasks chan string, results *sync.Map) {
	urls, size := bodyParser(current)
	fmt.Printf("%s has been parsed.\n", current)
	validCount := 0
	for _, x := range urls {
		_, ok := results.Load(x)
		if ok {continue}
		tasks <- x
		results.Store(x, "")
		validCount++
	}
	results.Store(current, fmt.Sprintf("%d|%d|%d", size, len(urls), validCount))
}

func bodyParser(target string) ([]string, int) {
	response, err := http.Get(target)
	if err != nil {
		fmt.Printf("%s", err)
		return []string{}, 0
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		return []string{}, 0
	}
	str := string(contents)
	response.Body.Close()
	var re = regexp.MustCompile(`href="(.*?)"`)
	links := re.FindAllStringSubmatch(str, -1)
	res := []string{}
	for _, x := range links {
		current := x[1]
		if strings.Contains(current, "xkcd.com") && strings.Contains(current, "http") {
			res = append(res, current)
		}
	}
	return res, len(str)
}
```
