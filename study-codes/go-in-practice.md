
## Chapter 1 Getting in to Go
- 1.1 What is Go?
- 1.2 Noteworthy Aspects of Go
- 1.2.1 Multiple Return Values
- 1.2.2 A Modern Standard Library
- 1.2.3 Concurrency with goroutines and channels
- 1.2.4 Go the toolchain, more than a language
- 1.3 Go in the vast language landscape
- 1.3.1 C and Go
- 1.3.2 Go and Java
- 1.3.3 Python, PHP, and Go
- 1.3.4 JavaScript, node.js, and Go
- 1.4 Getting up and running in Go
- 1.4.1 Installing Go
- 1.4.2 Git, Mercurial, and Version Control
- 1.4.3 Workspace
- 1.4.4 Environment Variables
- 1.5 Hello Go

#### chapter1 -> hello -> hello.go
```go
package main

import "fmt"

func getName() string {
	return "World!"
}

func main() {
	name := getName()
	fmt.Println("Hello ", name)
}
```
#### chapter1 -> channel.go
```go
package main

import (
	"fmt"
	"time"
)

func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c
		fmt.Print(num, " ")
	}
}

func main() {
	c := make(chan int)
	a := []int{8, 6, 7, 5, 3, 0, 9, -1}

	go printCount(c)

	for _, v := range a {
		c <- v
	}
	time.Sleep(time.Millisecond * 1)
	fmt.Println("End of main")
}
```
#### chapter1 -> inigo.go
```go
package main

import (
	"fmt"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello, my name is Inigo Montoya")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:4000", nil)
}
```
#### chapter1 -> hello -> hello_test.go
```go
package main

import "testing"

func TestName(t *testing.T) {
	name := getName()

	if name != "World!" {
		t.Error("Respone from getName is unexpected value")
	}
}
```
#### chapter1 -> goroutine.go
```go
package main

import (
	"fmt"
	"time"
)

func count() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 1)
	}
}

func main() {
	go count()
	time.Sleep(time.Millisecond * 2)
	fmt.Println("Hello World")
	time.Sleep(time.Millisecond * 5)
}
```
#### chapter1 -> http_get.go
```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://example.com/")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()
}
```
#### chapter1 -> returns.go
```go
package main

import (
	"fmt"
)

func Names() (string, string) {
	return "Foo", "Bar"
}

func main() {
	n1, n2 := Names()
	fmt.Println(n1, n2)

	n3, _ := Names()
	fmt.Println(n3)
}
```
#### chapter1 -> read_status.go
```go
package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "golang.org:80")
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}
```
#### chapter1 -> returns2.go
```go
package main

import (
	"fmt"
)

func Names() (first string, second string) {
	first = "Foo"
	second = "Bar"
	return
}

func main() {
	n1, n2 := Names()
	fmt.Println(n1, n2)
}
```


## Chapter 2 A Solid Foundation
- 2.1 CLI Applications, the Go Way
- 2.1.1 Command Line Flags
- 2.1.2 Command Line Frameworks
- 2.2 Handling Configuration
- 2.3 Working with Real World Web Servers
- 2.3.1 Starting up and shutting down a server
- 2.3.2 Routing Web Requests

#### chapter2 -> multiple_handlers.go
```go
package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/goodbye/", goodbye)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Goodbye ", name)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The homepage.")
}
```
#### chapter2 -> flag_cli.go
```go
package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "A name to say hello to.")

var spanish bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language.")
	flag.BoolVar(&spanish, "s", false, "Use Spanish language.")
}

func main() {
	flag.Parse()
	if spanish == true {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}
```
#### chapter2 -> regex_handlers.go
```go
package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	rr := newPathResolver()
	rr.Add("GET /hello", hello)
	rr.Add("(GET|HEAD) /goodbye(/?[A-Za-z0-9]*)?", goodbye)
	http.ListenAndServe(":8080", rr)
}

func newPathResolver() *regexResolver {
	return &regexResolver{
		handlers: make(map[string]http.HandlerFunc),
		cache:    make(map[string]*regexp.Regexp),
	}
}

type regexResolver struct {
	handlers map[string]http.HandlerFunc
	cache    map[string]*regexp.Regexp
}

func (r *regexResolver) Add(regex string, handler http.HandlerFunc) {
	r.handlers[regex] = handler
	cache, _ := regexp.Compile(regex)
	r.cache[regex] = cache
}

func (r *regexResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path
	for pattern, handlerFunc := range r.handlers {
		if r.cache[pattern].MatchString(check) == true {
			handlerFunc(res, req)
			return
		}
	}

	http.NotFound(res, req)
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := ""
	if len(parts) > 2 {
		name = parts[2]
	}
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Goodbye ", name)
}
```
#### chapter2 -> manners_shutdown.go
```go
package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/braintree/manners"
)

func main() {
	handler := newHandler()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go listenForShutdown(ch)

	manners.ListenAndServe(":8080", handler)
}

func newHandler() *handler {
	return &handler{}
}

type handler struct{}

func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

func listenForShutdown(ch <-chan os.Signal) {
	<-ch
	manners.Close()
}
```
#### chapter2 -> path_handlers.go
```go
package main

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

func main() {
	pr := newPathResolver()
	pr.Add("GET /hello", hello)
	pr.Add("* /goodbye/*", goodbye)
	http.ListenAndServe(":8080", pr)
}

func newPathResolver() *pathResolver {
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path
	for pattern, handlerFunc := range p.handlers {
		if ok, err := path.Match(pattern, check); ok && err == nil {
			handlerFunc(res, req)
			return
		} else if err != nil {
			fmt.Fprint(res, err)
		}
	}

	http.NotFound(res, req)
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Goodbye ", name)
}
```
#### chapter2 -> ini_config.go
```go
package main

import (
	"fmt"

	"code.google.com/p/gcfg"
)

func main() {
	config := struct {
		Section struct {
			Enabled bool
			Path    string
		}
	}{}

	err := gcfg.ReadFileInto(&config, "conf.ini")
	if err != nil {
		fmt.Println("Failed to parse config file: %s", err)
	}
	fmt.Println(config.Section.Enabled)
	fmt.Println(config.Section.Path)
}
```
#### chapter2 -> json_config.go
```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type configuration struct {
	Enabled bool
	Path    string
}

func main() {
	file, _ := os.Open("conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(conf.Path)
}
```
#### chapter2 -> hello_cli.go
```go
package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "hello_cli"
	app.Usage = "Print hello world"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Value: "World",
			Usage: "Who to say hello to.",
		},
	}
	app.Action = func(c *cli.Context) error {
		name := c.GlobalString("name")
		fmt.Printf("Hello %s!\n", name)

		return nil
	}

	app.Run(os.Args)
}
```
#### chapter2 -> callback_shutdown.go
```go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}

func shutdown(res http.ResponseWriter, req *http.Request) {
	os.Exit(0)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The homepage.")
}
```
#### chapter2 -> yaml_config.go
```go
package main

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
)

func main() {
	config, err := yaml.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Get("path"))
	fmt.Println(config.GetBool("enabled"))
}
```
#### chapter2 -> go_flags_example.go
```go
package main

import (
	"fmt"
	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	Name    string `short:"n" long:"name" default:"World" description:"A name to say hello to."`
	Spanish bool   `short:"s" long:"spanish" description:"Use Spanish Language"`
}

func main() {
	flags.Parse(&opts)

	if opts.Spanish == true {
		fmt.Printf("Hola %s!\n", opts.Name)
	} else {
		fmt.Printf("Hello %s!\n", opts.Name)
	}
}
```
#### chapter2 -> count_cli.go
```go
package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Count up or down."
	app.Commands = []cli.Command{
		{
			Name:      "up",
			ShortName: "u",
			Usage:     "Count Up",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "stop, s",
					Usage: "Value to count up to",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("stop")
				if start <= 0 {
					fmt.Println("Stop cannot be negative.")
				}
				for i := 1; i <= start; i++ {
					fmt.Println(i)
				}
				return nil
			},
		},
		{
			Name:      "down",
			ShortName: "d",
			Usage:     "Count Down",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "start, s",
					Usage: "Start counting down from",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("start")
				if start < 0 {
					fmt.Println("Start cannot be negative.")
				}
				for i := start; i >= 0; i-- {
					fmt.Println(i)
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
```
#### chapter2 -> env_config.go
```go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The homepage.")
}
```


## Chapter 3 Concurrency in Go
- 3.1 Understanding Go&#39;s Concurrency Model
- 3.2 Working with Goroutines
- 3.3 Working with Channels

#### chapter3 -> race -> fixed.go
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	w := newWords()
	for _, f := range os.Args[1:] {
		wg.Add(1)
		go func(file string) {
			if err := tallyWords(file, w); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(f)
	}
	wg.Wait()

	fmt.Println("Words that appear more than once:")
	w.Lock()
	for word, count := range w.found {
		if count > 1 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}
	w.Unlock()
}

type words struct {
	sync.Mutex
	found map[string]int
}

func newWords() *words {
	return &words{found: map[string]int{}}
}

func (w *words) add(word string, n int) {
	w.Lock()
	defer w.Unlock()
	count, ok := w.found[word]
	if !ok {
		w.found[word] = n
		return
	}
	w.found[word] = count + n
}

func tallyWords(filename string, dict *words) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict.add(word, 1)
	}
	return scanner.Err()
}
```
#### chapter3 -> simplelock -> lock.go
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	lock := make(chan bool, 1)
	for i := 1; i < 7; i++ {
		go worker(i, lock)
	}
	time.Sleep(10 * time.Second)
}

func worker(id int, lock chan bool) {
	fmt.Printf("%d wants the lock\n", id)
	lock <- true
	fmt.Printf("%d has the lock\n", id)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("%d is releasing the lock\n", id)
	<-lock
}
```
#### chapter3 -> waitgroup -> simple_gz.go
```go
package main

import (
	"compress/gzip"
	"io"
	"os"
	//"sync"
)

func main() {
	for _, file := range os.Args {
		compress(file)
	}
	//var wg sync.WaitGroup
}

func compress(filename string) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	gzout.Close()

	return err
}
```
#### chapter3 -> scratch -> main.go
```go
package main

import (
	"time"
)

func main() {
	//time.Sleep(5 * time.Second)

	//sleep := time.After(5 * time.Second)
	//<-sleep
	//pp := make(chan string)
	//go pingpong(pp, true)
	//go pingpong(pp, false)
	ch := make(chan bool)
	go ping(ch)
	go pong(ch)
	ch <- true
	select {}
}

func ping(ch chan bool) {
	for {
		select {
		case <-ch:
			println("ping")
			ch <- true
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func pong(ch chan bool) {
	for {
		select {
		case <-ch:
			println("pong")
			ch <- true
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func pingpong(ch chan string, start bool) {
	if start {
		ch <- "ping"
	}

	for {
		select {
		case str := <-ch:
			println(str)
			if str == "ping" {
				ch <- "pong"
				time.Sleep(200 * time.Millisecond)
				continue
			}
			ch <- "ping"
		}
	}
}
```
#### chapter3 -> waitgroup -> wg.go
```go
package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var i int
	var file string
	for i, file = range os.Args[1:] {
		wg.Add(1)
		go func(filename string) {
			compress(filename)
			wg.Done()
		}(file)
	}
	wg.Wait()

	fmt.Printf("Compressed %d files\n", i+1)
}

func compress(filename string) error {
	fmt.Printf("Compressing %s\n", filename)
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	gzout.Close()

	return err
}
```
#### chapter3 -> echoback -> echoback.go
```go
package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	go echo(os.Stdin, os.Stdout)
	time.Sleep(30 * time.Second)
	fmt.Println("Timed out.")
	os.Exit(0)
}

func echo(in io.Reader, out io.Writer) {
	io.Copy(out, in)
}
```
#### chapter3 -> closing -> bad.go
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)
	until := time.After(5 * time.Second)

	go send(msg)

	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			close(msg)
			time.Sleep(500 * time.Millisecond)
			return
		}
	}
}

func send(ch chan string) {
	for {
		ch <- "hello"
		time.Sleep(500 * time.Millisecond)
	}
}
```
#### chapter3 -> closing -> okay.go
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)
	done := make(chan bool)
	until := time.After(5 * time.Second)

	go send(msg, done)

	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			done <- true
			time.Sleep(500 * time.Millisecond)
			return
		}
	}
}

func send(ch chan<- string, done <-chan bool) {
	for {
		select {
		case <-done:
			println("Done")
			close(ch)
			return
		default:
			ch <- "hello"
			time.Sleep(500 * time.Millisecond)
		}
	}
}
```
#### chapter3 -> race -> race.go
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	w := newWords()
	for _, f := range os.Args[1:] {
		wg.Add(1)
		go func(file string) {
			if err := tallyWords(file, w); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(f)
	}
	wg.Wait()

	fmt.Println("Words that appear more than once:")
	for word, count := range w.found {
		if count > 1 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}
}

type words struct {
	found map[string]int
}

func newWords() *words {
	return &words{found: map[string]int{}}
}

func (w *words) add(word string, n int) {
	count, ok := w.found[word]
	if !ok {
		w.found[word] = n
		return
	}
	w.found[word] = count + n
}

func tallyWords(filename string, dict *words) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict.add(word, 1)
	}
	return scanner.Err()
}
```
#### chapter3 -> closures -> simple.go
```go
package main

import (
	"fmt"
	//"runtime"
)

func main() {
	fmt.Println("Outside a goroutine.")
	go func() {
		fmt.Println("Inside a goroutine")
	}()
	fmt.Println("Outside again.")

	//runtime.Gosched()
}
```
#### chapter3 -> scratch -> closingstuff.go
```go
package main

import "time"

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
	time.Sleep(120 * time.Millisecond)
	ch <- true
	close(ch)
	println("Sent and closed")
}
```
#### chapter3 -> closing -> close-sender.go
```go
package main

import "time"

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
	time.Sleep(120 * time.Millisecond)
	ch <- true
	close(ch)
	println("Sent and closed")
}
```
#### chapter3 -> echoredux -> echoredux.go
```go
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	done := time.After(30 * time.Second)
	echo := make(chan []byte)
	go readStdin(echo)
	for {
		select {
		case buf := <-echo:
			os.Stdout.Write(buf)
		case <-done:
			fmt.Println("Timed out")
			close(echo)
			os.Exit(0)
		}
	}
}

func readStdin(out chan<- []byte) {
	for {
		data := make([]byte, 1024)
		l, _ := os.Stdin.Read(data)
		if l > 0 {
			out <- data
		}
	}
}
```


## Chapter 4 Handling Errors and Panics
- 4.1 Error Handling
- 4.2 The Panic System
- 4.2.1 Differentiating panics from errors
- 4.2.2 Working with Panics
- 4.2.3 Recovering from Panics
- 4.2.4 Panics and Goroutines

#### chapter4 -> http_server.go
```go
package main

import (
	"errors"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
func handler(res http.ResponseWriter, req *http.Request) {
	panic(errors.New("Fake panic!"))
}
```
#### chapter4 -> simple_server_start.go
```go
package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
)

func main() {
	listen()
}

func listen() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Caught a panic and recovered")
		}
	}()
	listener, err := net.Listen("tcp", ":1026")
	if err != nil {
		fmt.Println("Failed to open port on 1026")
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection")
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	reader := bufio.NewReader(conn)

	data, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println("Failed to read from socket.")
		conn.Close()
	}

	response(data, conn)
}

func response(data []byte, conn net.Conn) {
	defer func() {
		conn.Close()
	}()
	conn.Write(data)
	panic(errors.New("Failure in response!"))
}
```
#### chapter4 -> zero_divider.go
```go
package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("Can't divide by zero")

func main() {
	fmt.Println("Divide 1 by 0")
	_, err := precheckDivide(1, 0)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Println("Divide 2 by 0")
	divide(2, 0)
}

func precheckDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return divide(a, b), nil
}

func divide(a, b int) int {
	return a / b
}
```
#### chapter4 -> safely -> safely.go
```go
package safely

import (
	"log"
)

type GoDoer func()

func Go(todo GoDoer) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic in safely.Go: %s", err)
			}
		}()
		todo()
	}()
}
```
#### chapter4 -> two_errors.go
```go
package main

import (
	"errors"
	"fmt"
	"math/rand"
)

var ErrTimeout = errors.New("The request timed out")
var ErrRejected = errors.New("The request was rejected")

var random = rand.New(rand.NewSource(35))

//var random = rand.New(rand.NewSource(78))

func main() {
	response, err := SendRequest("Hello")
	for err == ErrTimeout {
		fmt.Println("Timeout. Retrying.")
		response, err = SendRequest("Hello")
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}

func SendRequest(req string) (string, error) {

	switch random.Int() % 3 {
	case 0:
		return "Success", nil
	case 1:
		return "", ErrRejected
	default:
		return "", ErrTimeout
	}
}
```
#### chapter4 -> safely_example.go
```go
package main

import (
	"./safely"
	"errors"
	"time"
)

func message() {
	println("Inside goroutine")
	panic(errors.New("Oops!"))
}

func main() {
	safely.Go(message)
	println("Outside goroutine")
	time.Sleep(200)
}
```
#### chapter4 -> parser_error.go
```go
package main

import "fmt"

func main() {

	err := &ParseError{
		Message: "Unexpected char ';'",
		Line:    5,
		Char:    38,
	}

	fmt.Println(err.Error())
}

type ParseError struct {
	Message    string
	Line, Char int
}

func (p *ParseError) Error() string {
	format := "%s on Line %d, Char %d"
	return fmt.Sprintf(format, p.Message, p.Line, p.Char)
}
```
#### chapter4 -> useful_recover.go
```go
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	var file io.ReadCloser
	file, err := OpenCSV("data.csv")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	defer file.Close()

	// Do something with file.

}

func OpenCSV(filename string) (file *os.File, err error) {
	defer func() {
		if r := recover(); r != nil {
			file.Close()
			err = r.(error)
		}
	}()

	file, err = os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file\n")
		return file, err
	}

	RemoveEmptyLines(file)

	return file, err
}

func RemoveEmptyLines(f *os.File) {
	panic(errors.New("Failed parse"))
}
```
#### chapter4 -> error_example.go
```go
package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Concat concatenates a bunch of strings.
// Strings are separated by spaces.
// It returns an empty string and an error if no strings were passed in.
func Concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("No strings supplied")
	}

	return strings.Join(parts, " "), nil
}

func main() {

	args := os.Args[1:]

	/*
		if result, err := Concat(args...); err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Printf("Concatenated string: '%s'\n", result)
		}
	*/
	result, _ := Concat(args...)
	fmt.Printf("Concatenated string: '%s'\n", result)

}

func assumeGoodDesign() {
	// Streamlined error handling
	batch := []string{}
	result, _ := Concat(batch...)
	fmt.Printf("Concatenated string: '%s'\n", result)

}
```
#### chapter4 -> simple_server.go
```go
package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
)

func main() {
	listen()
}

func listen() {
	listener, err := net.Listen("tcp", ":1026")
	if err != nil {
		fmt.Println("Failed to open port on 1026")
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection")
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Fatal error: %s", err)
		}
		conn.Close()
	}()
	reader := bufio.NewReader(conn)

	data, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println("Failed to read from socket.")
	}

	response(data, conn)
}

func response(data []byte, conn net.Conn) {
	conn.Write(data)
	panic(errors.New("Pretend I'm a real error"))
}
```
#### chapter4 -> same_error.go
```go
package main

import "errors"

func main() {
	errA := errors.New("Error")
	//errB := errors.New("Error")

	if errA != errA {
		panic("Expected same!")
	}
	println("Done")
}
```
#### chapter4 -> simple_defer.go
```go
package main

import "fmt"

func main() {
	defer goodbye()

	fmt.Println("Hello world.")
}

func goodbye() {
	fmt.Println("Goodbye")
}
```
#### chapter4 -> recover_panic.go
```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	msg := "Everything's fine"
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Trapped panic: %s (%T)\n", err, err)
		}
		fmt.Println(msg)
	}()

	yikes()
}

func yikes() {
	panic(errors.New("Something bad happened."))
}
```
#### chapter4 -> two_defers.go
```go
package main

func main() {
	defer func() { println("a") }()
	defer func() { println("b") }()
}
```
#### chapter4 -> closure_scope.go
```go
package main

import "fmt"

func main() {
	var msg string
	defer func() {
		fmt.Println(msg)
	}()
	msg = "Hello world"
}
```
#### chapter4 -> proper_panic.go
```go
package main

import "errors"

func main() {
	panic(errors.New("Something bad happened."))
}
```


## Chapter 5 Debugging and Testing
- 5.1 Locating Bugs
- 5.1.1 Wait, Where Is My Debugger?
- 5.2 Logging
- 5.2.1 Using Go&#8217;s Logger
- 5.2.2 Working with System Loggers
- 5.3 Accessing Stack Traces
- 5.4 Testing
- 5.4.1 Unit Testing
- 5.4.2 Generative Testing
- 5.5 Performance Tests and Benchmarks

#### chapter5 -> tests -> generative -> generative_test.go
```go
package main

import (
	"log"
	"strings"
	"testing"
	"testing/quick"
)

// Pad whitespace-pads a string to a given length.
//
// If the string is longer than that, it truncates.
func Pad(s string, max uint) string {
	log.Printf("Testing Len: %d, Str: %s\n", max, s)
	l := uint(len(s))
	if l > max {
		return s[:max]
	}
	s += strings.Repeat(" ", int(max-l))
	return s
}

func TestPad(t *testing.T) {
	if r := Pad("test", 6); len(r) != 6 {
		t.Errorf("Expected 6, got %d", len(r))
	}
}

func TestPadGenerative(t *testing.T) {
	fn := func(s string, max uint8) bool {
		p := Pad(s, uint(max))
		return len(p) == int(max)
	}

	if err := quick.Check(fn, &quick.Config{MaxCount: 200}); err != nil {
		t.Error(err)
	}
}
```
#### chapter5 -> tests -> canary -> canary_test.go
```go
package canary

import (
	"io"
	"testing"
)

type MyWriter struct{}

func (m *MyWriter) Write([]byte) error {
	return nil
}

func main() {
	m := map[string]interface{}{
		"w": &MyWriter{},
	}
}

func doSomething(m map[string]interface{}) {
	w := m["w"].(io.Writer)
}

func TestWriter(t *testing.T) {
	var _ io.Writer = &MyWriter{}
}
```
#### chapter5 -> logs -> log_client.go
```go
package main

import (
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:1902")
	if err != nil {
		panic("Failed to connect to localhost:1902")
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example ", f)

	//logger := log.New(logfile, "example ", log.LstdFlags|log.Lshortfile)

	logger.Println("This is a regular message.")
	logger.Panicln("This is a panic.")
}
```
#### chapter5 -> tests -> msg -> send_message_test.go
```go
package msg

import (
	"testing"
)

type MockMessage struct {
	email, subject string
	body           []byte
}

func (m *MockMessage) Send(email, subject string, body []byte) error {
	m.email = email
	m.subject = subject
	m.body = body
	return nil
}

func TestAlert(t *testing.T) {
	msgr := new(MockMessage)
	body := []byte("Critical Error")

	Alert(msgr, body)

	if msgr.subject != "Critical Error" {
		t.Errorf("Expected 'Critical Error', Got '%s'", msgr.subject)
	}
	// ...
}
```
#### chapter5 -> logs -> syslog.go
```go
package main

import (
	"log/syslog"
)

func main() {
	logger, err := syslog.New(syslog.LOG_LOCAL3, "narwhal")
	if err != nil {
		panic("Cannot attach to syslog")
	}

	logger.Debug("Debug message.")
	logger.Notice("Notice message.")
	logger.Warning("Warning message.")
	logger.Alert("Alert message.")
}
```
#### chapter5 -> stack -> trace.go
```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	foo()
}

func foo() {
	bar()
}

func bar() {
	buf := make([]byte, 1024)
	runtime.Stack(buf, false)
	fmt.Printf("Trace:\n %s\n", buf)

	caller, file, line, _ := runtime.Caller(0)
	fmt.Printf("Ptr: %d File: %s, Line: %d", caller, file, line)
}
```
#### chapter5 -> tests -> hello -> unit_test.go
```go
package hello

import "testing"

func TestHello(t *testing.T) {
	if Hello() != "hello" {
		t.Errorf("Expected 'hello', but got '%s'", Hello())
	}
}
```
#### chapter5 -> delvetest -> main.go
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	time.Sleep(2 * time.Second)
	fmt.Println("OH HAI!")
}
```
#### chapter5 -> tests -> msg -> send_message.go
```go
package msg

type Message struct {
	// ...
}

func (m *Message) Send(email, subject string, body []byte) error {
	// ...
	return nil
}

type Messager interface {
	Send(email, subject string, body []byte) error
}

func Alert(m Messager, problem []byte) error {
	return m.Send("noc@example.com", "Critical Error", problem)
}
```
#### chapter5 -> logs -> syslog_logger.go
```go
package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	priority := syslog.LOG_LOCAL3 | syslog.LOG_NOTICE
	flags := log.Ldate | log.Lshortfile
	logger, err := syslog.NewLogger(priority, flags)
	if err != nil {
		fmt.Printf("Can't attach to syslog: %s", err)
		return
	}

	logger.Println("This is a test log message.")
}
```
#### chapter5 -> logs -> simple.go
```go
package main

import (
	"log"
)

func main() {
	log.Println("This is a regular message.")
	log.Fatalln("This is a fatal error.")
	log.Println("This is the end of the function.")
}
```
#### chapter5 -> logs -> outfile.go
```go
package main

import (
	"log"
	"os"
)

func main() {
	logfile, _ := os.Create("./log.txt")
	defer logfile.Close()

	f := log.Ldate | log.Lmicroseconds | log.Lshortfile | log.Llongfile
	logger := log.New(logfile, "example ", f)

	//logger := log.New(logfile, "example ", log.LstdFlags|log.Lshortfile)

	logger.Println("This is a regular message.")
	logger.Fatalln("This is a fatal error.")
	logger.Println("This is the end of the function.")
}
```
#### chapter5 -> tests -> bench -> bench_test.go
```go
package main

import (
	"bytes"
	"testing"
	"text/template"
)

func BenchmarkTemplates(b *testing.B) {
	// b.Logf("b.N is %d\n", b.N)
	tpl := "Hello {{.Name}}"
	data := &map[string]string{
		"Name": "World",
	}
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		t, _ := template.New("test").Parse(tpl)
		t.Execute(&buf, data)
		buf.Reset()
	}
}
func BenchmarkCompiledTemplates(b *testing.B) {
	// b.Logf("b.N is %d\n", b.N)
	tpl := "Hello {{.Name}}"
	t, _ := template.New("test").Parse(tpl)
	data := &map[string]string{
		"Name": "World",
	}
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		t.Execute(&buf, data)
		buf.Reset()
	}
}

func BenchmarkParallelTemplates(b *testing.B) {
	tpl := "Hello {{.Name}}"
	t, _ := template.New("test").Parse(tpl)
	data := &map[string]string{
		"Name": "World",
	}
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			t.Execute(&buf, data)
			buf.Reset()
		}
	})
}

func BenchmarkParallelOops(b *testing.B) {
	tpl := "Hello {{.Name}}"
	t, _ := template.New("test").Parse(tpl)
	data := &map[string]string{
		"Name": "World",
	}
	var buf bytes.Buffer
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t.Execute(&buf, data)
			buf.Reset()
		}
	})
}
```
#### chapter5 -> logs -> log_server.go
```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

// Simple log server
func main() {

	listener, err := net.Listen("tcp", ":1902")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Failed to accept: %s", err)
		}
		printmsg(conn)
	}
}

func printmsg(conn net.Conn) {
	var err error
	var msg string
	for {
		msg, err = bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Error reading message: %s\n", err)
			}
			return
		}
		fmt.Println(msg)
	}

}
```
#### chapter5 -> logs -> buffered_logger.go
```go
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	logger := New("localhost:1902", 30*time.Second)

	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 20; i++ {

		go func(logger *TcpLogger, id int) {
			for j := 0; j < 100; j++ {
				logger.Printf("Client %d message %d.", id, j)
				time.Sleep(1 * time.Second)
			}
			wg.Done()
		}(logger, i)
	}

	wg.Wait()
}

type TcpLogger struct {
	*log.Logger
	Addr    string
	Timeout time.Duration
	Queue   *logQueue
}

func New(addr string, timeout time.Duration) *TcpLogger {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		panic("Failed to connect to localhost:1902")
	}

	buff := make(chan []byte, 10)
	queue := &logQueue{conn, buff}

	f := log.Ldate | log.Lshortfile
	logger := log.New(queue, "example ", f)

	l := &TcpLogger{
		logger,
		addr,
		timeout,
		queue,
	}

	go DequeueLogs(l)

	return l
}

func (t *TcpLogger) Reconnect() error {
	limit := 10
	for i := 0; i < limit; i++ {
		conn, err := net.DialTimeout("tcp", t.Addr, t.Timeout)
		if err == nil {
			t.Queue.Destination = conn
			return nil
		}
		time.Sleep(100 * time.Millisecond)
	}

	msg := "Failed to reconnect after %d tries."
	return fmt.Errorf(msg, limit)
}

type logQueue struct {
	Destination io.Writer
	Messages    chan []byte
}

func (l *logQueue) Write(data []byte) (int, error) {

	// Write can never modify data, so we copy it.
	msg := make([]byte, len(data))
	copy(msg, data)

	l.Messages <- msg

	return len(data), nil
}

func DequeueLogs(logger *TcpLogger) {
	for msg := range logger.Queue.Messages {
		if _, e := logger.Queue.Destination.Write(msg); e != nil {
			fmt.Println("Attempting reconnect.")
			if e := logger.Reconnect(); e != nil {
				fmt.Printf("Failed reconnect. Dropping message. Msg: %s", msg)
			} else {
				fmt.Println("Reconnected. Resending messages.")
				logger.Queue.Destination.Write(msg)
			}
		}
	}
}
```
#### chapter5 -> tests -> hello -> unit.go
```go
package hello

func Hello() string {
	return "hello"
}
```
#### chapter5 -> logs -> udp_logger.go
```go
package main

import (
	"log"
	"net"
	"time"
)

func main() {

	timeout := 30 * time.Second
	conn, err := net.DialTimeout("udp", "localhost:1902", timeout)
	if err != nil {
		panic("Failed to connect to localhost:1902")
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example ", f)

	//logger := log.New(logfile, "example ", log.LstdFlags|log.Lshortfile)

	logger.Println("This is a regular message.")
	logger.Panicln("This is a panic.")
}
```


## Chapter 6 HTML and Email Template Patterns
- 6.1 Working with HTML templates
- 6.1.1 Standard library HTML package overview
- 6.1.2 Adding Functionality Inside Templates
- 6.1.3 Limiting template parsing
- 6.1.4 When template execution breaks
- 6.1.5 Mixing Templates
- 6.2 Using Templates for Email

#### chapter6 -> inherit_templates -> inherit.go
```go
package main

import (
	"html/template"
	"net/http"
)

var t map[string]*template.Template

func init() {
	t = make(map[string]*template.Template)
	temp := template.Must(template.ParseFiles("base.html", "user.html"))
	t["user.html"] = temp
	temp = template.Must(template.ParseFiles("base.html", "page.html"))
	t["page.html"] = temp
}

type Page struct {
	Title, Content string
}

type User struct {
	Username, Name string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "An Example",
		Content: "Have fun stormin’ da castle.",
	}
	t["page.html"].ExecuteTemplate(w, "base", p)
}

func displayUser(w http.ResponseWriter, r *http.Request) {
	u := &User{
		Username: "swordsmith",
		Name:     "Inigo Montoya",
	}
	t["user.html"].ExecuteTemplate(w, "base", u)
}

func main() {
	http.HandleFunc("/user", displayUser)
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
```
#### chapter6 -> simple_template.go
```go
package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title, Content string
}

func diaplayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "An Example",
		Content: "Have fun stormin’ da castle.",
	}
	t := template.Must(template.ParseFiles("templates/simple.html"))
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", diaplayPage)
	http.ListenAndServe(":8080", nil)
}
```
#### chapter6 -> cache_template.go
```go
package main

import (
	"html/template"
	"net/http"
)

var t = template.Must(template.ParseFiles("./templates/simple.html"))

type Page struct {
	Title, Content string
}

func diaplayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "An Example",
		Content: "Have fun stormin’ da castle.",
	}
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", diaplayPage)
	http.ListenAndServe(":8080", nil)
}
```
#### chapter6 -> email.go
```go
package main

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strconv"
	"text/template"
)

type EmailMessage struct {
	From, Subject, Body string
	To                  []string
}

type EmailCredentials struct {
	Username, Password, Server string
	Port                       int
}

const emailTemplate = `From: {{.From}}
To: {{.To}}
Subject {{.Subject}}

{{.Body}}
`

var t *template.Template

func init() {
	t = template.New("email")
	t.Parse(emailTemplate)
}

func main() {
	message := &EmailMessage{
		From:    "me@example.com",
		To:      []string{"you@example.com"},
		Subject: "A test",
		Body:    "Just saying hi",
	}

	var body bytes.Buffer
	t.Execute(&body, message)

	fmt.Printf("%s", body.Bytes())

	authCreds := &EmailCredentials{
		Username: "myUsername",
		Password: "myPass",
		Server:   "smtp.example.com",
		Port:     25,
	}

	auth := smtp.PlainAuth("",
		authCreds.Username,
		authCreds.Password,
		authCreds.Server,
	)

	smtp.SendMail(authCreds.Server+":"+strconv.Itoa(authCreds.Port),
		auth,
		message.From,
		message.To,
		body.Bytes())
}
```
#### chapter6 -> nested_templates -> nested_templates.go
```go
package main

import (
	"html/template"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("index.html", "head.html"))
}

type Page struct {
	Title, Content string
}

func diaplayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "An Example",
		Content: "Have fun stormin’ da castle.",
	}
	t.ExecuteTemplate(w, "index.html", p)
}

func main() {
	http.HandleFunc("/", diaplayPage)
	http.ListenAndServe(":8080", nil)
}
```
#### chapter6 -> buffered_template.go
```go
package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("./templates/simple.html"))
}

type Page struct {
	Title, Content string
}

func diaplayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "An Example",
		Content: "Have fun stormin’ da castle.",
	}
	var b bytes.Buffer
	err := t.Execute(&b, p)
	if err != nil {
		fmt.Fprint(w, "A error occured.")
		return
	}
	b.WriteTo(w)
}

func main() {
	http.HandleFunc("/", diaplayPage)
	http.ListenAndServe(":8080", nil)
}
```
#### chapter6 -> object_templates -> object_templates.go
```go
package main

import (
	"bytes"
	"html/template"
	"net/http"
)

var t *template.Template
var qc template.HTML

func init() {
	t = template.Must(template.ParseFiles("index.html", "quote.html"))
}

type Page struct {
	Title   string
	Content template.HTML
}

type Quote struct {
	Quote, Person string
}

func main() {
	q := &Quote{
		Quote: `You keep using that word. I do not think
				it means what you think it means.`,
		Person: "Inigo Montoya",
	}
	var b bytes.Buffer
	t.ExecuteTemplate(&b, "quote.html", q)
	qc = template.HTML(b.String())

	http.HandleFunc("/", diaplayPage)
	http.ListenAndServe(":8080", nil)
}

func diaplayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "A Quote",
		Content: qc,
	}
	t.ExecuteTemplate(w, "index.html", p)
}
```
#### chapter6 -> date_command.go
```go
package main

import (
	"html/template"
	"net/http"
	"time"
)

var tpl = `<!DOCTYPE HTML>
<html>
  <head>
    <meta charset="utf-8">
    <title>Date Example</title>
  </head>
  <body>
  	<p>{{.Date | dateFormat "Jan 2, 2006"}}</p>
  </body>
</html>`

var funcMap = template.FuncMap{
	"dateFormat": dateFormat,
}

func dateFormat(layout string, d time.Time) string {
	return d.Format(layout)
}

func serveTemplate(res http.ResponseWriter, req *http.Request) {
	t := template.New("date")
	t.Funcs(funcMap)
	t.Parse(tpl)
	data := struct{ Date time.Time }{
		Date: time.Now(),
	}
	t.Execute(res, data)
}

func main() {
	http.HandleFunc("/", serveTemplate)
	http.ListenAndServe(":8080", nil)
}
```


## Chapter 7 Serving and Receiving, Assets and Forms
- 7.1 Serving static content
- 7.2 Handling Form Posts
- 7.2.1 Introduction to form requests
- 7.2.2 Working with files and miltipart submissions
- 7.2.3 Working with raw multipart data

#### chapter7 -> file_increment_save.go
```go
package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func fileForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("file_plus.html")
		t.Execute(w, nil)
	} else {
		mr, err := r.MultipartReader()
		values := make(map[string][]string)

		if err != nil {
			panic("Failed to read multipart message")
		}

		maxValueBytes := int64(10 << 20)
		for {

			part, err := mr.NextPart()
			if err == io.EOF {

				break
			}

			name := part.FormName()
			if name == "" {
				continue
			}

			filename := part.FileName()
			var b bytes.Buffer
			if filename == "" {
				n, err := io.CopyN(&b, part, maxValueBytes)
				if err != nil && err != io.EOF {
					fmt.Fprint(w, "Error processing form")
					return
				}
				maxValueBytes -= n
				if maxValueBytes == 0 {
					fmt.Fprint(w, "multipart message too large")
					return
				}
				values[name] = append(values[name], b.String())
				continue
			}

			dst, err := os.Create("/tmp/dstfile." + filename)
			defer dst.Close()
			if err != nil {
				return
			}
			for {
				buffer := make([]byte, 100000)
				cBytes, err := part.Read(buffer)
				if err == io.EOF {
					break
				}
				dst.Write(buffer[0:cBytes])
			}
		}
		fmt.Println("Upload done")
		fmt.Println(values)

		fmt.Fprint(w, "Upload complete")
	}
}

func main() {
	http.HandleFunc("/", fileForm)
	http.ListenAndServe(":8080", nil)
}
```
#### chapter7 -> one_level_serving.go
```go
package main

import (
	"fmt"
	"net/http"
	"path"
)

func main() {
	pr := newPathResolver()
	pr.Add("GET /hello", hello)
	dir := http.Dir("./files")
	handler := http.StripPrefix("/static/", http.FileServer(dir))
	pr.Add("GET /static/*", handler.ServeHTTP)
	http.ListenAndServe(":8080", pr)
}

func newPathResolver() *pathResolver {
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path
	for pattern, handlerFunc := range p.handlers {
		if ok, err := path.Match(pattern, check); ok && err == nil {
			handlerFunc(res, req)
			return
		} else if err != nil {
			fmt.Fprint(res, err)
		}
	}

	http.NotFound(res, req)
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}
```
#### chapter7 -> file_not_found.go
```go
package main

import (
	"fmt"
	fs "github.com/Masterminds/go-fileserver"
	"net/http"
)

func main() {

	fs.NotFoundHandler = func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintln(w, "The requested page could not be found.")
	}

	dir := http.Dir("./files")
	http.ListenAndServe(":8080", fs.FileServer(dir))
}
```
#### chapter7 -> serve_subdir_simple.go
```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	dir := http.Dir("./files/")
	handler := http.StripPrefix("/static/", http.FileServer(dir))
	http.Handle("/static/", handler)

	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The homepage.")
}
```
#### chapter7 -> cache_serving.go
```go
package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

type cacheFile struct {
	content io.ReadSeeker
	modTime time.Time
}

var cache map[string]*cacheFile
var mutex = new(sync.RWMutex)

func main() {
	cache = make(map[string]*cacheFile)
	http.HandleFunc("/", serveFiles)
	http.ListenAndServe(":8080", nil)
}

func serveFiles(res http.ResponseWriter, req *http.Request) {
	mutex.RLock()
	v, found := cache[req.URL.Path]
	mutex.RUnlock()

	if !found {
		mutex.Lock()
		defer mutex.Unlock()
		fileName := "./files" + req.URL.Path
		f, err := os.Open(fileName)
		defer f.Close()

		if err != nil {
			http.NotFound(res, req)
			return
		}

		var b bytes.Buffer
		_, err = io.Copy(&b, f)
		if err != nil {
			http.NotFound(res, req)
			return
		}
		r := bytes.NewReader(b.Bytes())

		info, _ := f.Stat()
		v = &cacheFile{
			content: r,
			modTime: info.ModTime(),
		}
		cache[req.URL.Path] = v
	}

	http.ServeContent(res, req, req.URL.Path, v.modTime, v.content)
}
```
#### chapter7 -> embedded_files -> embedded_files.go
```go
package main

import (
	"github.com/GeertJohan/go.rice"
	"net/http"
)

func main() {
	box := rice.MustFindBox("../files/")
	httpbox := box.HTTPBox()
	http.ListenAndServe(":8080", http.FileServer(httpbox))
}
```
#### chapter7 -> form_file.go
```go
package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func fileForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("file.html")
		t.Execute(w, nil)
	} else {
		f, h, err := r.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		filename := "/tmp/" + h.Filename
		out, _ := os.Create(filename)
		defer out.Close()

		io.Copy(out, f)
		fmt.Fprint(w, "Upload complete")
	}
}

func main() {
	http.HandleFunc("/", fileForm)
	http.ListenAndServe(":8080", nil)
}
```
#### chapter7 -> form_file_multiple.go
```go
package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func fileForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("file_multiple.html")
		t.Execute(w, nil)
	} else {
		err := r.ParseMultipartForm(16 << 20)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}

		data := r.MultipartForm
		files := data.File["files"]
		for _, fh := range files {
			f, err := fh.Open()
			defer f.Close()
			if err != nil {
				fmt.Fprint(w, err)
				return
			}

			out, err := os.Create("/tmp/" + fh.Filename)
			defer out.Close()
			if err != nil {
				fmt.Fprint(w, err)
				return
			}

			_, err = io.Copy(out, f)

			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
		}

		fmt.Fprint(w, "Upload complete")
	}
}

func main() {
	http.HandleFunc("/", fileForm)
	http.ListenAndServe(":8080", nil)
}
```
#### chapter7 -> file_mime_type.go
```go
package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func fileForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("file.html")
		t.Execute(w, nil)
	} else {
		f, h, err := r.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		filename := "/tmp/" + h.Filename
		out, _ := os.Create(filename)
		defer out.Close()
		fmt.Println(h.Header["Content-Type"][0])

		buff := make([]byte, 512)
		_, err = f.Read(buff)
		filetype := http.DetectContentType(buff)
		fmt.Println(filetype)

		io.Copy(out, f)
		fmt.Fprint(w, "Upload complete")
	}
}

func main() {
	http.HandleFunc("/", fileForm)
	http.ListenAndServe(":8080", nil)
}
```
#### chapter7 -> file_serving.go
```go
package main

import (
	"net/http"
)

func main() {
	dir := http.Dir("./files")
	http.ListenAndServe(":8080", http.FileServer(dir))
}
```
#### chapter7 -> servefile.go
```go
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", readme)
	http.ListenAndServe(":8080", nil)
}

func readme(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./files/readme.txt")
}
```
#### chapter7 -> file_increment_save_notes.go
```go
package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func fileForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("file.html")
		t.Execute(w, nil)
		fmt.Println(os.TempDir())
	} else {

		mr, err := r.MultipartReader()
		if err != nil {
			panic("Failed to read multipart message: ")
		}

		length := r.ContentLength
		for {

			part, err := mr.NextPart()
			if err == io.EOF {

				break
			}
			var read int64
			var p float32

			filename := part.FileName()
			dst, err := os.Create("/tmp/dstfile." + filename)
			if err != nil {
				return
			}
			for {
				buffer := make([]byte, 100000)
				cBytes, err := part.Read(buffer)
				if err == io.EOF {
					break
				}
				read = read + int64(cBytes)
				//fmt.Printf("read: %v \n",read )
				p = float32(read) / float32(length) * 100
				fmt.Printf("progress: %v \n", p)
				dst.Write(buffer[0:cBytes])
			}
		}

		// f, h, err := r.FormFile("file")
		// if err != nil {
		// 	panic(err)
		// }
		// defer f.Close()
		// filename := "/tmp/" + h.Filename
		// out, _ := os.Create(filename)
		// defer out.Close()
		// fmt.Println(h.Header["Content-Type"][0])

		// buff := make([]byte, 512)
		// _, err = f.Read(buff)
		// filetype := http.DetectContentType(buff)
		// fmt.Println(filetype)

		// io.Copy(out, f)
		fmt.Println("Upload done")

		fmt.Fprint(w, "Upload complete")
	}
}

func main() {
	http.HandleFunc("/", fileForm)
	http.ListenAndServe(":8080", nil)
}
```
#### chapter7 -> serve_alternate_location.go
```go
package main

import (
	"flag"
	"html/template"
	"net/http"
)

var t *template.Template
var l = flag.String("location", "http://localhost:8080", "A location.")

var tpl = `<!DOCTYPE HTML>
<html>
  <head>
    <meta charset="utf-8">
    <title>A Demo</title>
    <link rel="stylesheet" href="{{.Location}}/styles.css">
  </head>
  <body>
  	<p>A demo.</p>
  </body>
</html>`

func init() {
	t = template.Must(template.New("date").Parse(tpl))
}

func servePage(res http.ResponseWriter, req *http.Request) {
	data := struct{ Location *string }{
		Location: l,
	}
	t.Execute(res, data)
}

func main() {
	flag.Parse()
	http.HandleFunc("/", servePage)
	http.ListenAndServe(":8080", nil)
}
```


## Chapter 8 Working with Web Services
- 8.1 Using REST APIs
- 8.1.1 Using The HTTP Client
- 8.1.2 When Faults Happen
- 8.2 Passing And Handling Errors Over HTTP
- 8.2.1 Generating Custom Errors
- 8.2.2 Reading And Using Custom Errors
- 8.3 Parsing And Mapping JSON
- 8.4 Versioning REST APIs

#### chapter8 -> versioned_url_api.go
```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type testMessage struct {
	Message string `json:"message"`
}

func displayTest(w http.ResponseWriter, r *http.Request) {
	data := testMessage{"A test message."}
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(b))
}

func main() {
	http.HandleFunc("/api/v1/test", displayTest)
	http.ListenAndServe(":8080", nil)
}
```
#### chapter8 -> http_timeout_handling.go
```go
package main

import (
	// "io"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, err := os.Create("ff.dmg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	location := "https://download-installer.cdn.mozilla.net/pub/firefox/releases/40.0.3/mac/en-US/Firefox%2040.0.3.dmg"
	err = download(location, file, 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	fi, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Got it with %v bytes downloaded", fi.Size())
}

func download(location string, file *os.File, retries int64) error {
	req, err := http.NewRequest("GET", location, nil)
	if err != nil {
		return err
	}
	fi, err := file.Stat()
	if err != nil {
		return err
	}
	current := fi.Size()
	if current > 0 {
		start := strconv.FormatInt(current, 10)
		req.Header.Set("Range", "bytes="+start+"-")
	}

	cc := &http.Client{Timeout: 5 * time.Minute}
	res, err := cc.Do(req)
	if err != nil && hasTimedOut(err) {
		if retries > 0 {
			return download(location, file, retries-1)
		}
		return err
	} else if err != nil {
		return err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		errFmt := "Unsuccess HTTP request. Status: %s"
		return fmt.Errorf(errFmt, res.Status)
	}

	if res.Header.Get("Accept-Ranges") != "bytes" {
		retries = 0
	}

	_, err = io.Copy(file, res.Body)
	if err != nil && hasTimedOut(err) {
		if retries > 0 {
			return download(location, file, retries-1)
		}
		return err
	} else if err != nil {
		return err
	}

	return nil
}

func hasTimedOut(err error) bool {
	switch err := err.(type) {
	case *url.Error:
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			return true
		}
	case net.Error:
		if err.Timeout() {
			return true
		}
	case *net.OpError:
		if err.Timeout() {
			return true
		}
	}

	errTxt := "use of closed network connection"
	if err != nil && strings.Contains(err.Error(), errTxt) {
		return true
	}

	return false
}
```
#### chapter8 -> get_custom_error.go
```go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Error struct {
	HTTPCode int    `json:"-"`
	Code     int    `json:"code,omitempty"`
	Message  string `json:"message"`
}

func (e Error) Error() string {
	fs := "HTTP: %d, Code: %d, Message: %s"
	return fmt.Sprintf(fs, e.HTTPCode, e.Code, e.Message)
}

func get(u string) (*http.Response, error) {
	res, err := http.Get(u)
	if err != nil {
		return res, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		if res.Header.Get("Content-Type") != "application/json" {
			sm := "Unknown error. HTTP status: %s"
			return res, fmt.Errorf(sm, res.Status)
		}
		b, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		var data struct {
			Err Error `json:"error"`
		}
		err = json.Unmarshal(b, &data)
		if err != nil {
			sm := "Unable to parse json: %s. HTTP status: %s"
			return res, fmt.Errorf(sm, err, res.Status)
		}
		data.Err.HTTPCode = res.StatusCode

		return res, data.Err
	}

	return res, nil
}

func main() {
	res, err := get("http://localhost:8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", b)
}
```
#### chapter8 -> http_error.go
```go
package main

import "net/http"

func displayError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "An Error Occurred", http.StatusForbidden)
}

func main() {
	http.HandleFunc("/", displayError)
	http.ListenAndServe(":8080", nil)
}
```
#### chapter8 -> custom_request_delete.go
```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	req, _ := http.NewRequest("DELETE", "http://example.com/foo/bar", nil)
	res, _ := http.DefaultClient.Do(req)
	fmt.Printf("%s", res.Status)
}
```
#### chapter8 -> parse_arbitrary_json.go
```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var ks = []byte(`{
"firstName": "Jean",
"lastName": "Bartik",
"age": 86,
"education": [
	{
		"institution": "Northwest Missouri State Teachers College",
		"degree": "Bachelor of Science in Mathematics"
	},
	{
		"institution": "University of Pennsylvania",
		"degree": "Masters in English"
	}
],
"spouse": "William Bartik",
"children": [
	"Timothy John Bartik",
	"Jane Helen Bartik",
	"Mary Ruth Bartik"
]
}`)

func main() {
	var f interface{}
	err := json.Unmarshal(ks, &f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(f)

	m := f.(map[string]interface{})
	fmt.Println(m["firstName"])

	fmt.Print("interface{} ")
	printJSON(f)
}

func printJSON(v interface{}) {
	switch vv := v.(type) {
	case string:
		fmt.Println("is string", vv)
	case float64:
		fmt.Println("is float64", vv)
	case []interface{}:
		fmt.Println("is an array:")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJSON(u)
		}
	case map[string]interface{}:
		fmt.Println("is an object:")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJSON(u)
		}
	default:
		fmt.Println("Unknown type")
	}
}
```
#### chapter8 -> get_semantic_version_api.go
```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	ct := "application/vnd.mytodos.json; version=2.0"
	req, _ := http.NewRequest("GET", "http://localhost:8080/test", nil)
	req.Header.Set("Accept", ct)
	res, _ := http.DefaultClient.Do(req)
	if res.Header.Get("Content-Type") != ct {
		fmt.Println("Unexpected content type returned")
		return
	}
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", b)
}
```
#### chapter8 -> simple_get.go
```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("http://goinpracticebook.com")
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", b)
}
```
#### chapter8 -> semantic_version_api.go
```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/test", displayTest)
	http.ListenAndServe(":8080", nil)
}

func displayTest(w http.ResponseWriter, r *http.Request) {
	t := r.Header.Get("Accept")
	var err error
	var b []byte
	var ct string
	switch t {
	case "application/vnd.mytodos.json; version=2.0":
		data := testMessageV2{"Version 2"}
		b, err = json.Marshal(data)
		ct = "application/vnd.mytodos.json; version=2.0"
	case "application/vnd.mytodos.json; version=1.0":
		fallthrough
	default:
		data := testMessageV1{"Version 1"}
		b, err = json.Marshal(data)
		ct = "application/vnd.mytodos.json; version=1.0"
	}

	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	w.Header().Set("Content-Type", ct)
	fmt.Fprint(w, string(b))
}

type testMessageV1 struct {
	Message string `json:"message"`
}

type testMessageV2 struct {
	Info string `json:"info"`
}
```
#### chapter8 -> http_custom_error.go
```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	HTTPCode int    `json:"-"`
	Code     int    `json:"code,omitempty"`
	Message  string `json:"message"`
}

func JSONError(w http.ResponseWriter, e Error) {
	data := struct {
		Err Error `json:"error"`
	}{e}
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.HTTPCode)
	fmt.Fprint(w, string(b))
}

func displayError(w http.ResponseWriter, r *http.Request) {
	e := Error{
		HTTPCode: http.StatusForbidden,
		Code:     123,
		Message:  "An Error Occurred",
	}
	JSONError(w, e)
}

func main() {
	http.HandleFunc("/", displayError)
	http.ListenAndServe(":8080", nil)
}
```
#### chapter8 -> simple_json_parser.go
```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
}

var JSON = `{
  "name": "Miracle Max"
}`

func main() {
	var p Person
	err := json.Unmarshal([]byte(JSON), &p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(p)
}
```
#### chapter8 -> custom_client_timeout.go
```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	cc := &http.Client{Timeout: time.Second}
	res, err := cc.Get("http://goinpracticebook.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", b)
}
```


## Chapter 9 Using the Cloud
- 9.1 What Is Cloud Computing?
- 9.1.1 The Different Types of Cloud
- 9.1.2 Containers and Cloud Native Applications
- 9.2 Managing Cloud Services
- 9.2.1 Avoid Cloud Provider Lock&#151;in
- 9.2.2 Dealing with Divergent Errors
- 9.3 Running On Cloud Servers
- 9.3.1 Runtime Detection
- 9.3.2 Building For The Cloud
- 9.3.3 Runtime Monitoring

#### chapter9 -> hostlookup.go
```go
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return
	}

	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, a := range addrs {
		fmt.Println(a)
	}
}
```
#### chapter9 -> api_interface_with_errors.go
```go
package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	ErrFileNotFound   = errors.New("File not found")
	ErrCannotLoadFile = errors.New("Unable to load file")
	ErrCannotSaveFile = errors.New("Unable to save file")
)

type File interface {
	Load(string) (io.ReadCloser, error)
	Save(string, io.ReadSeeker) error
}

type LocalFile struct {
	Base string
}

func (l LocalFile) Load(path string) (io.ReadCloser, error) {
	p := filepath.Join(l.Base, path)
	var oerr error
	o, err := os.Open(p)
	if err != nil && os.IsNotExist(err) {
		log.Printf("Unable to find %s", path)
		oerr = ErrFileNotFound
	} else if err != nil {
		log.Printf("Error loading file %s, err: %s", path, err)
		oerr = ErrCannotLoadFile
	}
	return o, oerr
}

func (l LocalFile) Save(path string, body io.ReadSeeker) error {
	p := filepath.Join(l.Base, path)
	d := filepath.Dir(p)
	err := os.MkdirAll(d, os.ModeDir|os.ModePerm)
	if err != nil {
		log.Printf("Cannot make directory '%s', err: %s", d, err)
		return ErrCannotSaveFile
	}

	f, err := os.Create(p)
	if err != nil {
		log.Printf("Cannot create file '%s', err: %s", p, err)
		return ErrCannotSaveFile
	}
	defer f.Close()

	_, err = io.Copy(f, body)
	if err != nil {
		log.Printf("Cannot save file '%s', err: %s", path, err)
		return ErrCannotSaveFile
	}
	return nil
}

func fileStore() (File, error) {
	return &LocalFile{Base: "/Users/mfarina/store"}, nil
}

func main() {
	content := `Lorem ipsum dolor sit amet, consectetur` +
		`adipiscing elit. Donec a diam lectus.Sed sit` +
		`amet ipsum mauris. Maecenascongue ligula ac` +
		`quam viverra nec consectetur ante hendrerit.`
	body := bytes.NewReader([]byte(content))

	store, err := fileStore()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Storing content...")
	err = store.Save("foo/bar", body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Retrieving content...")
	c, err := store.Load("foo/bar")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	o, err := ioutil.ReadAll(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(o))
}
```
#### chapter9 -> api_interface.go
```go
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type File interface {
	Load(string) (io.ReadCloser, error)
	Save(string, io.ReadSeeker) error
}

type LocalFile struct {
	Base string
}

func (l LocalFile) Load(path string) (io.ReadCloser, error) {
	p := filepath.Join(l.Base, path)
	return os.Open(p)
}

func (l LocalFile) Save(path string, body io.ReadSeeker) error {
	p := filepath.Join(l.Base, path)
	d := filepath.Dir(p)
	err := os.MkdirAll(d, os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, body)
	return err
}

func fileStore() (File, error) {
	return &LocalFile{Base: "/Users/mfarina/store"}, nil
}

func main() {
	content := `Lorem ipsum dolor sit amet, consectetur` +
		`adipiscing elit. Donec a diam lectus.Sed sit` +
		`amet ipsum mauris. Maecenascongue ligula ac` +
		`quam viverra nec consectetur ante hendrerit.`
	body := bytes.NewReader([]byte(content))

	store, err := fileStore()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Storing content...")
	err = store.Save("foo/bar", body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Retrieving content...")
	c, err := store.Load("foo/bar")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	o, err := ioutil.ReadAll(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(o))
}
```
#### chapter9 -> runtime.go
```go
package main

import (
	"log"
	"runtime"
	"time"
)

func monitorRuntime() {
	log.Println("Number of CPUs:", runtime.NumCPU())
	m := &runtime.MemStats{}
	for {
		r := runtime.NumGoroutine()
		log.Println("Number of goroutines", r)

		runtime.ReadMemStats(m)
		log.Println("Allocated memory", m.Alloc)
		time.Sleep(10 * time.Second)
	}
}

func main() {
	go monitorRuntime()

	i := 0
	for i < 40 {
		go func() {
			time.Sleep(15 * time.Second)
		}()
		i = i + 1
		time.Sleep(1 * time.Second)
	}
}
```
#### chapter9 -> detect_dep.go
```go
package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	err := checkDep("fortune")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Time to get your fortunte")
}

func checkDep(name string) error {
	if _, err := exec.LookPath(name); err != nil {
		es := "Could not find '%s' in PATH: %s"
		return fmt.Errorf(es, name, err)
	}

	return nil
}
```


## Chapter 10 Communication Between Cloud Services
- 10.1 Microservices and High Availability
- 10.2 Communicating Between Services
- 10.2.1 Making REST Faster
- 10.2.2 Moving Beyond REST

#### chapter10 -> user -> user_generated.go
```go
// ************************************************************
// DO NOT EDIT.
// THIS FILE IS AUTO-GENERATED BY codecgen.
// ************************************************************

package user

import (
	"errors"
	"fmt"
	codec1978 "github.com/ugorji/go/codec"
	"reflect"
	"runtime"
)

const (
	// ----- content types ----
	codecSelferC_UTF85734 = 1
	codecSelferC_RAW5734  = 0
	// ----- value types used ----
	codecSelferValueTypeArray5734 = 10
	codecSelferValueTypeMap5734   = 9
	// ----- containerStateValues ----
	codecSelfer_containerMapKey5734    = 2
	codecSelfer_containerMapValue5734  = 3
	codecSelfer_containerMapEnd5734    = 4
	codecSelfer_containerArrayElem5734 = 6
	codecSelfer_containerArrayEnd5734  = 7
)

var (
	codecSelferBitsize5734                         = uint8(reflect.TypeOf(uint(0)).Bits())
	codecSelferOnlyMapOrArrayEncodeToStructErr5734 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer5734 struct{}

func init() {
	if codec1978.GenVersion != 5 {
		_, file, _, _ := runtime.Caller(0)
		err := fmt.Errorf("codecgen version mismatch: current: %v, need %v. Re-generate file: %v",
			5, codec1978.GenVersion, file)
		panic(err)
	}
	if false { // reference the types, but skip this branch at build/run time
	}
}

func (x *User) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer5734
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [2]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[1] = x.Email != ""
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(2)
			} else {
				yynn2 = 1
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.EncodeMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem5734)
				yym4 := z.EncBinary()
				_ = yym4
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF85734, string(x.Name))
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey5734)
				r.EncodeString(codecSelferC_UTF85734, string("name"))
				z.EncSendContainerState(codecSelfer_containerMapValue5734)
				yym5 := z.EncBinary()
				_ = yym5
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF85734, string(x.Name))
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem5734)
				if yyq2[1] {
					yym7 := z.EncBinary()
					_ = yym7
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF85734, string(x.Email))
					}
				} else {
					r.EncodeString(codecSelferC_UTF85734, "")
				}
			} else {
				if yyq2[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey5734)
					r.EncodeString(codecSelferC_UTF85734, string("Email"))
					z.EncSendContainerState(codecSelfer_containerMapValue5734)
					yym8 := z.EncBinary()
					_ = yym8
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF85734, string(x.Email))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd5734)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd5734)
			}
		}
	}
}

func (x *User) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer5734
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym9 := z.DecBinary()
	_ = yym9
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct10 := r.ContainerType()
		if yyct10 == codecSelferValueTypeMap5734 {
			yyl10 := r.ReadMapStart()
			if yyl10 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd5734)
			} else {
				x.codecDecodeSelfFromMap(yyl10, d)
			}
		} else if yyct10 == codecSelferValueTypeArray5734 {
			yyl10 := r.ReadArrayStart()
			if yyl10 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd5734)
			} else {
				x.codecDecodeSelfFromArray(yyl10, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr5734)
		}
	}
}

func (x *User) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer5734
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys11Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys11Slc
	var yyhl11 bool = l >= 0
	for yyj11 := 0; ; yyj11++ {
		if yyhl11 {
			if yyj11 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey5734)
		yys11Slc = r.DecodeBytes(yys11Slc, true, true)
		yys11 := string(yys11Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue5734)
		switch yys11 {
		case "name":
			if r.TryDecodeAsNil() {
				x.Name = ""
			} else {
				x.Name = string(r.DecodeString())
			}
		case "Email":
			if r.TryDecodeAsNil() {
				x.Email = ""
			} else {
				x.Email = string(r.DecodeString())
			}
		default:
			z.DecStructFieldNotFound(-1, yys11)
		} // end switch yys11
	} // end for yyj11
	z.DecSendContainerState(codecSelfer_containerMapEnd5734)
}

func (x *User) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer5734
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj14 int
	var yyb14 bool
	var yyhl14 bool = l >= 0
	yyj14++
	if yyhl14 {
		yyb14 = yyj14 > l
	} else {
		yyb14 = r.CheckBreak()
	}
	if yyb14 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd5734)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem5734)
	if r.TryDecodeAsNil() {
		x.Name = ""
	} else {
		x.Name = string(r.DecodeString())
	}
	yyj14++
	if yyhl14 {
		yyb14 = yyj14 > l
	} else {
		yyb14 = r.CheckBreak()
	}
	if yyb14 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd5734)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem5734)
	if r.TryDecodeAsNil() {
		x.Email = ""
	} else {
		x.Email = string(r.DecodeString())
	}
	for {
		yyj14++
		if yyhl14 {
			yyb14 = yyj14 > l
		} else {
			yyb14 = r.CheckBreak()
		}
		if yyb14 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem5734)
		z.DecStructFieldNotFound(yyj14-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd5734)
}
```
#### chapter10 -> closing_connections.go
```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	r, err := http.Get("http://mattfarina.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	o, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	r.Body.Close()
	fmt.Println(string(o))

	r2, err := http.Get("http://mattfarina.com/about")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	o, err = ioutil.ReadAll(r2.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	r2.Body.Close()
	fmt.Println(string(o))
}
```
#### chapter10 -> hellopb -> hello.pb.go
```go
// Code generated by protoc-gen-go.
// source: hello.proto
// DO NOT EDIT!

/*
Package chapter10 is a generated protocol buffer package.

It is generated from these files:
	hello.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
*/
package chapter10

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HelloResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*HelloRequest)(nil), "chapter10.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "chapter10.HelloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Hello service

type HelloClient interface {
	Say(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) Say(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := grpc.Invoke(ctx, "/chapter10.Hello/Say", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hello service

type HelloServer interface {
	Say(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(HelloServer).Say(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chapter10.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Say",
			Handler:    _Hello_Say_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0xce, 0x48, 0x2c, 0x28, 0x49, 0x2d,
	0x32, 0x34, 0x50, 0x52, 0xe2, 0xe2, 0xf1, 0x00, 0xc9, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97,
	0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81,
	0xd9, 0x4a, 0x9a, 0x5c, 0xbc, 0x50, 0x35, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x12, 0x5c,
	0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0x30, 0x75, 0x30, 0xae, 0x91, 0x33, 0x17, 0x2b, 0x58,
	0xa9, 0x90, 0x15, 0x17, 0x73, 0x70, 0x62, 0xa5, 0x90, 0xb8, 0x1e, 0xdc, 0x2a, 0x3d, 0x64, 0x7b,
	0xa4, 0x24, 0x30, 0x25, 0x20, 0x86, 0x2b, 0x31, 0x24, 0xb1, 0x81, 0x5d, 0x69, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x07, 0xf3, 0x53, 0x0a, 0xb4, 0x00, 0x00, 0x00,
}
```
#### chapter10 -> grpc_server.go
```go
package main

import (
	"log"
	"net"

	pb "github.com/Masterminds/go-in-practice/chapter10/hellopb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Say(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	msg := "Hello " + in.Name + "!"
	return &pb.HelloResponse{Message: msg}, nil
}

func main() {
	l, err := net.Listen("tcp", ":55555")
	if err != nil {
		log.Fatalf("failed to listen for tcp: %s", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})
	s.Serve(l)
}
```
#### chapter10 -> userpb -> user.pb.go
```go
// Code generated by protoc-gen-go.
// source: user.proto
// DO NOT EDIT!

/*
Package chapter10 is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	User
*/
package chapter10

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type User struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Id               *int32  `protobuf:"varint,2,req,name=id" json:"id,omitempty"`
	Email            *string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *User) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *User) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "chapter10.User")
}

var fileDescriptor0 = []byte{
	// 94 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0xce, 0x48, 0x2c, 0x28, 0x49, 0x2d, 0x32,
	0x34, 0x50, 0xd2, 0xe7, 0x62, 0x09, 0x05, 0x4a, 0x08, 0xf1, 0x70, 0xb1, 0xe4, 0x25, 0xe6, 0xa6,
	0x4a, 0x30, 0x2a, 0x30, 0x69, 0x70, 0x0a, 0x71, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x01, 0xd9,
	0xac, 0x42, 0xbc, 0x5c, 0xac, 0xa9, 0xb9, 0x89, 0x99, 0x39, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x9c,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x61, 0x0d, 0xf1, 0xf6, 0x48, 0x00, 0x00, 0x00,
}
```
#### chapter10 -> json.go
```go
package main

import (
	"fmt"
	"os"

	"github.com/Masterminds/go-in-practice/chapter10/user"
	"github.com/ugorji/go/codec"
)

func main() {
	jh := new(codec.JsonHandle)
	u := &user.User{
		Name:  "Inigo Montoya",
		Email: "inigo@montoya.example.com",
	}

	var out []byte
	err := codec.NewEncoderBytes(&out, jh).Encode(&u)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))

	var u2 user.User
	err = codec.NewDecoderBytes(out, jh).Decode(&u2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(u2)
}
```
#### chapter10 -> user -> user.go
```go
//go:generate codecgen -o user_generated.go user.go

package user

type User struct {
	Name  string `codec:"name"`
	Email string `codec:",omitempty"`
}
```
#### chapter10 -> grpc_client.go
```go
package main

import (
	"fmt"
	"os"

	pb "github.com/Masterminds/go-in-practice/chapter10/hellopb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	address := "localhost:55555"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("unable to connect:", err)
		os.Exit(1)
	}
	defer conn.Close()
	c := pb.NewHelloClient(conn)

	name := "Inigo Montoya"
	hr := &pb.HelloRequest{Name: name}
	r, err := c.Say(context.Background(), hr)
	if err != nil {
		fmt.Println("could not say:", err)
		os.Exit(1)
	}
	fmt.Println(r.Message)
}
```
#### chapter10 -> multiple_connections.go
```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	r, err := http.Get("http://mattfarina.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer r.Body.Close()
	o, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(o))

	r2, err := http.Get("http://mattfarina.com/about")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer r2.Body.Close()
	o, err = ioutil.ReadAll(r2.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(o))
}
```
#### chapter10 -> protobuf_client.go
```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	pb "github.com/Masterminds/go-in-practice/chapter10/userpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	res, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body: %v", err)
		os.Exit(1)
	}

	var u pb.User
	err = proto.Unmarshal(b, &u)
	if err != nil {
		fmt.Println("Error decoding response body: %v", err)
		os.Exit(1)
	}

	fmt.Println(u.GetName())
	fmt.Println(u.GetId())
	fmt.Println(u.GetEmail())
}
```
#### chapter10 -> protobuf_server.go
```go
package main

import (
	"net/http"

	pb "github.com/Masterminds/go-in-practice/chapter10/userpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(res http.ResponseWriter, req *http.Request) {
	u := &pb.User{
		Name:  proto.String("Inigo Montoya"),
		Id:    proto.Int32(1234),
		Email: proto.String("inigo@montoya.example.com"),
	}

	body, err := proto.Marshal(u)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/x-protobuf")
	res.Write(body)
}
```


## Chapter 11 Reflection and Code Generation
- 11.1 Three features of reflection
- 11.2 Structs, tags, and annotations
- 11.2.1 Annotating structs
- 11.2.2 Using tag annotations
- 11.3 Generating Go code with Go code

#### chapter11 -> generator -> queue.go
```go
package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

var tpl = `package {{.Package}}

type {{.MyType}}Queue struct {
	q []{{.MyType}}
}

func New{{.MyType}}Queue() *{{.MyType}}Queue {
	return &{{.MyType}}Queue{
		q: []{{.MyType}}{},
	}
}

func (o *{{.MyType}}Queue) Insert(v {{.MyType}}) {
	o.q = append(o.q, v)
}

func (o *{{.MyType}}Queue) Remove() {{.MyType}} {
	if len(o.q) == 0 {
		panic("Oops.")
	}
	first := o.q[0]
	o.q = o.q[1:]
	return first
}
`

func main() {
	tt := template.Must(template.New("queue").Parse(tpl))
	for i := 1; i < len(os.Args); i++ {
		dest := strings.ToLower(os.Args[i]) + "_queue.go"
		file, err := os.Create(dest)
		if err != nil {
			fmt.Printf("Could not create %s: %s (skip)\n", dest, err)
			continue
		}

		vals := map[string]string{
			"MyType":  os.Args[i],
			"Package": os.Getenv("GOPACKAGE"),
		}
		tt.Execute(file, vals)

		file.Close()
	}
}
```
#### chapter11 -> values -> sumkind.go
```go
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type MyInt int64

func main() {
	var a uint8 = 2
	var b int = 37
	var c string = "3.2"
	var d MyInt = 1
	res := sum(a, b, c, d)
	fmt.Printf("Result: %f\n", res)
}

func sum(v ...interface{}) float64 {
	var res float64 = 0
	for _, val := range v {
		ref := reflect.ValueOf(val)
		switch ref.Kind() {
		case reflect.Int, reflect.Int64:
			res += float64(ref.Int())
		case reflect.Uint8:
			res += float64(ref.Uint())
		case reflect.String:
			a, err := strconv.ParseFloat(ref.String(), 64)
			if err != nil {
				panic(err)
			}
			res += a
		default:
			fmt.Printf("Unsupported type %T. Ignoring.\n", val)
		}
	}
	return res
}
```
#### chapter11 -> annotations -> json.go
```go
package main

import (
	"encoding/json"
	"fmt"
)

type Name struct {
	First string `json:"firstName" xml:"FirstName"`
	Last  string `json:"lastName,omitempty"`
	Other string `not,even.a=tag`
}

func main() {
	n := &Name{"Inigo", "Montoya", ""}
	data, _ := json.Marshal(n)
	fmt.Printf("%s\n", data)
}
```
#### chapter11 -> annotations -> jsonxml -> jsonxml.go
```go
package main

import (
	"encoding/json"
	"encoding/xml"
	"os"
)

type Person struct {
	FirstName string `json:"first" xml:"firstName,attr"`
	LastName  string `json:"last" xml:"lastName"`
}

func main() {
	p := &Person{FirstName: "Inigo", LastName: "Montoya"}
	j, _ := json.MarshalIndent(p, "", "  ")
	os.Stdout.Write(j)
	println()

	x, _ := xml.MarshalIndent(p, "", "  ")
	os.Stdout.Write(x)
}
```
#### chapter11 -> interfaces -> typeassert.go
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBuffer([]byte("Hello"))
	if isStringer(b) {
		fmt.Printf("%T is a stringer\n", b)
	}
	i := 123
	if isStringer(i) {
		fmt.Printf("%T is a stringer\n", i)
	}
}

func isStringer(v interface{}) bool {
	_, ok := v.(fmt.Stringer)
	return ok
}
```
#### chapter11 -> generator -> myint_queue.go
```go
package main

type MyIntQueue struct {
	q []MyInt
}

func NewMyIntQueue() *MyIntQueue {
	return &MyIntQueue{
		q: []MyInt{},
	}
}

func (o *MyIntQueue) Insert(v MyInt) {
	o.q = append(o.q, v)
}

func (o *MyIntQueue) Remove() MyInt {
	if len(o.q) == 0 {
		panic("Oops.")
	}
	first := o.q[0]
	o.q = o.q[1:]
	return first
}
```
#### chapter11 -> structs -> structwalker.go
```go
package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("Walking a simple integer")
	var one MyInt = 1
	walk(one, 0)

	fmt.Println("Walking a simple struct")
	two := struct{ Name string }{"foo"}
	walk(two, 0)

	p := &Person{
		Name:    &Name{"Count", "Tyrone", "Rugen"},
		Address: &Address{"Humperdink Castle", "Florian"},
	}
	fmt.Println("Walking a struct with struct fields")
	walk(p, 0)
}

type MyInt int

type Person struct {
	Name    *Name
	Address *Address
}

type Name struct {
	Title, First, Last string
}

type Address struct {
	Street, Region string
}

func walk(u interface{}, depth int) {
	val := reflect.Indirect(reflect.ValueOf(u))
	t := val.Type()
	tabs := strings.Repeat("\t", depth+1)
	fmt.Printf("%sValue is type %q (%s)\n", tabs, t, val.Kind())
	if val.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldVal := reflect.Indirect(val.Field(i))

			tabs := strings.Repeat("\t", depth+2)
			fmt.Printf("%sField %q is type %q (%s)\n",
				tabs, field.Name, field.Type, fieldVal.Kind())

			if fieldVal.Kind() == reflect.Struct {
				walk(fieldVal.Interface(), depth+1)
			}
		}
	}
}
```
#### chapter11 -> annotations -> load.go
```go
package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Processes struct {
	Total    int     `ini:"total"`
	Running  int     `ini:"running"`
	Sleeping int     `ini:"sleeping"`
	Threads  int     `ini:"threads"`
	Load     float64 `ini:"load"`
}

func main() {
	fmt.Println("Write a struct to output:")
	proc := &Processes{
		Total:    23,
		Running:  3,
		Sleeping: 20,
		Threads:  34,
		Load:     1.8,
	}
	data, err := Marshal(proc)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	fmt.Println("Read the data back into a struct")
	proc2 := &Processes{}
	if err := Unmarshal(data, proc2); err != nil {
		panic(err)
	}
	fmt.Printf("Struct: %#v", proc2)
}

func fieldName(field reflect.StructField) string {
	if t := field.Tag.Get("ini"); t != "" {
		return t
	}
	return field.Name
}

func Marshal(v interface{}) ([]byte, error) {
	var b bytes.Buffer
	val := reflect.Indirect(reflect.ValueOf(v))
	if val.Kind() != reflect.Struct {
		return []byte{}, errors.New("unmarshal can only take structs")
	}

	t := val.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name := fieldName(f)
		raw := val.Field(i).Interface()
		fmt.Fprintf(&b, "%s=%v\n", name, raw)
	}
	return b.Bytes(), nil
}

func Unmarshal(data []byte, v interface{}) error {

	val := reflect.Indirect(reflect.ValueOf(v))
	t := val.Type()

	b := bytes.NewBuffer(data)
	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		line := scanner.Text()
		pair := strings.SplitN(line, "=", 2)
		if len(pair) < 2 {
			// Skip any malformed lines.
			continue
		}
		setField(pair[0], pair[1], t, val)
	}
	return nil
}

func setField(name, value string, t reflect.Type, v reflect.Value) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if name == fieldName(field) {
			var dest reflect.Value
			switch field.Type.Kind() {
			default:
				fmt.Printf("Kind %s not supported.\n", field.Type.Kind())
				continue
			case reflect.Int:
				ival, err := strconv.Atoi(value)
				if err != nil {
					fmt.Printf("Could not convert %q to int: %s\n", value, err)
					continue
				}
				dest = reflect.ValueOf(ival)
			case reflect.Float64:
				fval, err := strconv.ParseFloat(value, 64)
				if err != nil {
					fmt.Printf("Could not convert %q to float64: %s\n", value, err)
					continue
				}
				dest = reflect.ValueOf(fval)
			case reflect.String:
				dest = reflect.ValueOf(value)
			case reflect.Bool:
				bval, err := strconv.ParseBool(value)
				if err != nil {
					fmt.Printf("Could not convert %q to bool: %s\n", value, err)
					continue
				}
				dest = reflect.ValueOf(bval)
			}
			v.Field(i).Set(dest)
		}
	}
}
```
#### chapter11 -> generator -> simple.go
```go
//go:generate echo hello
package main

func main() {
	println("Goodbyte")
}
```
#### chapter11 -> interfaces -> implements.go
```go
package main

import (
	"fmt"
	"io"
	"reflect"
)

type Name struct {
	First, Last string
}

func (n *Name) String() string {
	return n.First + " " + n.Last
}

func main() {
	n := &Name{First: "Inigo", Last: "Montoya"}

	stringer := (*fmt.Stringer)(nil)
	implements(n, stringer)

	writer := (*io.Writer)(nil)
	implements(n, writer)

	s := &[]string{}
	implements(s, writer)
}

func implements(concrete interface{}, target interface{}) bool {
	iface := reflect.TypeOf(target).Elem()

	v := reflect.ValueOf(concrete)
	t := v.Type()

	if t.Implements(iface) {
		fmt.Printf("%T is a %s\n", concrete, iface.Name())
		return true
	}
	fmt.Printf("%T is not a %s\n", concrete, iface.Name())
	return false
}
```
#### chapter11 -> generator -> myint.go
```go
//go:generate ./queue MyInt
package main

import "fmt"

type MyInt int

func main() {
	var one, two, three MyInt = 1, 2, 3
	q := NewMyIntQueue()
	q.Insert(one)
	q.Insert(two)
	q.Insert(three)

	fmt.Printf("First value: %d\n", q.Remove())
}
```
