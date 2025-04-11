package main

import (
	"time"
	"file_transfer/send"
)




func main () {
	filePath := "volume1/source_file" 
	time.Sleep(4 * time.Second)
	send.SendFile(filePath)
}


