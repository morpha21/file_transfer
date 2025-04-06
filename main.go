package main

import (
	"crypto/rand"
	"time"
	"io"
	"fmt"
	"net"
	"log"
)


type FileServer struct {}

func  (fs *FileServer) start () {
	ln, err := net.Listen("tcp", ":3000")
	checkErr(err)

	for {
		conn, err := ln.Accept()
		checkErr(err)
		go fs.readLoop(conn)
	}
}

func (fs *FileServer) readLoop(conn net.Conn){
	buf := make([]byte, 2048)

	for{
		n, err := conn.Read(buf)
		checkErr(err)
		file := buf[:n]
		fmt.Println(file)
		fmt.Printf("received %d bytes over the network\n", n)
	}
}

func sendFile(size int) error {
	file := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, file)
	checkErr(err)

	conn, err := net.Dial("tcp", ":3000")
	checkErr(err)

	n, err := conn.Write(file)
	checkErr(err)
	fmt.Printf("written %d bytes over the network\n", n)
	return nil
}


func main () {
	go func () {
		time.Sleep(4 * time.Second)
		sendFile(1000)
	}()

	server := &FileServer{}
	server.start()
}


func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
