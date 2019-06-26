package goroutines

import (
	"io"
	"log"
	"net"
	"os"
)

// Netcat1 is netcat like client
func Netcat2() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy2(os.Stdout, conn)
	mustCopy2(conn, os.Stdin)
}

func mustCopy2(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
