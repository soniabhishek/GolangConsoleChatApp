package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	getLocalIp()
	go takeIn()
	ln, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ln.Addr())
	for {
		time.Sleep(1000 * time.Second)
	}
}
func takeIn() {
	for {
		r := bufio.NewReader(os.Stdin)
		o, err := r.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(o))
	}
}
func getLocalIp() {
	host, err := os.Hostname()
	lip, err := net.LookupHost(host)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lip)
}
