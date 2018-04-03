



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





