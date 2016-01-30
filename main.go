package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	ln net.Listener
)

type Peer struct {
	conn net.Conn
	ip   string
}

func osSignal() {
	channel := make(chan os.Signal)
	signal.Notify(channel, syscall.SIGINT)
	<-channel
	fmt.Println("leaving server")
	err := ln.Close()
	if err != nil {
		log.Println("ln.Close():", err)
	}
	os.Exit(0)
}
func main() {
	go osSignal()
	fmt.Println("Hello User your Ip is ", getLocalIp())
	fmt.Println("Want to establish a new server or want to be a normal node")
	fmt.Println("Press y to establish a server or press n to get connected to a node")
	fmt.Println("---------------------------------------------------------------------------------------------------------")
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
	listPeers := list.New()
	fmt.Println("server called")
	ln, _ = net.Listen("tcp", ":10000")
	for {
		conn, err := ln.Accept()
		p := Peer{}
		p.conn = conn
		p.ip = conn.RemoteAddr().String()
		listPeers.PushBack(p)
		if err != nil {
			log.Fatal(err)
		}
		conn.Write([]byte("helllo new connn"))
		go readConn(conn, listPeers)
		fmt.Println(conn.RemoteAddr())

	}
}
func readConn(conn net.Conn, listPeers *list.List) {
	for {
		buffer := make([]byte, 1024)
		conn.Read(buffer)
		fmt.Print(string(buffer))
		for iter := listPeers.Front(); iter != nil; iter = iter.Next() {
			sender := []byte(iter.Value.(Peer).conn.RemoteAddr().String() + ": " + string(buffer))
			if iter.Value.(Peer).conn != conn {
				iter.Value.(Peer).conn.Write(sender)
			}
		}
	}
}

func joinPeer() {
	conn, err := net.Dial("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Fatal(err)
	}
	conn.Write([]byte("joined new peer"))
	go chat(conn)
	for {
		buffer := make([]byte, 1024)
		conn.Read(buffer)
		fmt.Println(string(buffer))
	}
	fmt.Println(conn.RemoteAddr())
}
func chat(conn net.Conn) {
	for {
		r := bufio.NewReader(os.Stdin)
		o, err := r.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}
		conn.Write(o)
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
