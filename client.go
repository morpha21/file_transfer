package main

import (
	"time"
	"file_transfer/send"
)




func main () {
	time.Sleep(4 * time.Second)
	send.SendFile(1000)
}


