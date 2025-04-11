package send

import (
	"fmt"
	"net"
	"file_transfer/error_handling"
	"os"
)


func SendFile(filePath string) error {
	file, err := os.ReadFile(filePath)
	error_handling.CheckErr(err)

	conn, err := net.Dial("tcp", ":3000")
	error_handling.CheckErr(err)

	n, err := conn.Write(file)
	error_handling.CheckErr(err)
	fmt.Printf("written %d bytes over the network\n", n)
	return nil
}


