- main.go
```go
package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
	"os"
)

func main() {

	var checkgems bool
	var h, service, command string
	flag.StringVar(&h, "host", "", "specify the host such as 10.6.0.134 or contour-bonaire-1b-qa-vm.devstaging.gems.energy")
	flag.StringVar(&service, "s", "", "specify the service")
	flag.BoolVar(&checkgems, "checkgems", false, "specify if only checkgems need to run")
	flag.Parse()

	if net.ParseIP(h) == nil {
		log.Printf("You input a DNS name %s and please make sure it is valid", h)
	}

	if checkgems {
		command = "sudo checkgems"
	} else if len(service) == 0 {
		log.Fatalf("you need specify the servie")
	} else {
		command = fmt.Sprintf(`sudo systemctl restart %s; sleep 1s; sudo systemctl status %s`, service, service)
	}

	log.Printf(`the commad [%s] will send to %s`, command, h)

	signer, err := ssh.ParsePrivateKey([]byte(_PRIVATE_KEY))
	if err != nil {
		log.Fatalf("cannot parse private key: %v", err)
	}

	config := ssh.ClientConfig{
		User: "nessus",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	address := fmt.Sprintf("%s:22", h)
	conn, err := ssh.Dial("tcp", address, &config)

	if err != nil {
		log.Fatalf("cannot connect: %v", err)
	}

	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("cannot set up session: %v", err)

	}

	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Fatalf("unable to setup stdout for session: %v", err)
	}
	go io.Copy(os.Stdout, stdout)

	stderr, err := session.StderrPipe()
	if err != nil {
		log.Fatalf("unable to setup stderr for session: %v", err)
	}
	go io.Copy(os.Stdout, stderr) // stderr in Jenkins may cause fake failure

	err = session.Run(command)

	if err != nil {
		log.Fatalf("cannot send request: %v", err)

	}
	defer conn.Close()

}
```
