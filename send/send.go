package send

import (
	"crypto/rand"
	"fmt"
	"io"
	"net"
	"file_transfer/error_handling"
)


func SendFile(size int) error {
        file := make([]byte, size)
        _, err := io.ReadFull(rand.Reader, file)
	error_handling.CheckErr(err)

        conn, err := net.Dial("tcp", ":3000")
	error_handling.CheckErr(err)

        n, err := conn.Write(file)
	error_handling.CheckErr(err)
        fmt.Printf("written %d bytes over the network\n", n)
        return nil
}


