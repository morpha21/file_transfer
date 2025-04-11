package main

import (
	"fmt"
	"net"
	"file_transfer/error_handling"
	"os"
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
	// var file []byte
	for{
		n, err := conn.Read(buf)
		error_handling.CheckErr(err)
		file := buf[:n]
		fmt.Println(file)
		fmt.Printf("received %d bytes over the network\n", n)
		err = os.WriteFile("volume2/arquivo", []byte(file), 0644)
		error_handling.CheckErr(err)
	}

}



func main () {
	server := &FileServer{}
	server.start()
}


