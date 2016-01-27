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
	fmt.Println("Hello User your Ip is ", getLocalIp())
	fmt.Println("Want to establish a new server or want to be a normal node")
	decide, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	if decide[0] == 'y' {
		go server()
	} else if decide[0] == 'n' {
		go joinPeer()
	}

	for {
		time.Sleep(1000 * time.Second)
	}
}

func server() {
	fmt.Println("server called")
	ln, _ := net.Listen("tcp", ":10000")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		conn.Write([]byte("helllo new connn"))
		fmt.Println(conn.RemoteAddr())

	}
}
func joinPeer() {
	conn, err := net.Dial("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Fatal(err)
	}
	conn.Write([]byte("joined new peer"))
	for {
		buffer := make([]byte, 1024)
		conn.Read(buffer)
		fmt.Println(string(buffer))
	}
	fmt.Println(conn.RemoteAddr())
}

func acceptConn() {

}
func chat(conn net.Conn) {
	for {
		r := bufio.NewReader(os.Stdin)
		o, err := r.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(o))
	}
}
func getLocalIp() string {
	host, err := os.Hostname()
	lip, err := net.LookupHost(host)
	if err != nil {
		log.Fatal(err)
	}
	return lip[0]
}
