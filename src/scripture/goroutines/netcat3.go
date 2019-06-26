package goroutines

import (
	"io"
	"log"
	"net"
	"os"
)

// Netcat3 is netcat like client
func Netcat3() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			log.Println("done")
			done <- struct{}{}
		}
	}()
	mustCopy3(conn, os.Stdin)
	defer conn.Close()
	<-done
}

func mustCopy3(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
