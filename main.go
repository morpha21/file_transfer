package main

import (
	"time"
	"fmt"
	"net"
	"file_transfer/send"
	"file_transfer/error_handling"
)


type FileServer struct {}

func  (fs *FileServer) start () {
	ln, err := net.Listen("tcp", ":3000")
	error_handling.CheckErr(err)

	for {
		conn, err := ln.Accept()
		error_handling.CheckErr(err)
		go fs.readLoop(conn)
	}
}

func (fs *FileServer) readLoop(conn net.Conn){
	buf := make([]byte, 2048)

	for{
		n, err := conn.Read(buf)
		error_handling.CheckErr(err)
		file := buf[:n]
		fmt.Println(file)
		fmt.Printf("received %d bytes over the network\n", n)
	}
}



func main () {
	go func () {
		time.Sleep(4 * time.Second)
		send.SendFile(1000)
	}()

	server := &FileServer{}
	server.start()
}


