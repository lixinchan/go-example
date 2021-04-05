package goroutines

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, short string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(short))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", short)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(short))
	time.Sleep(delay)
}

func handleEchoConn(conn net.Conn) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		echo(conn, input.Text(), 1*time.Second)
	}
	defer conn.Close()
}
