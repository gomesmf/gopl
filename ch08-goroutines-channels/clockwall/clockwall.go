// Clockwall command acts as a client of several clock servers at once,
// reading the times from each one and displaying the results in a table.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func mustCopy(headerClockConnMap map[string]net.Conn) {
	line := []string{}
	for header, conn := range headerClockConnMap {
		fmt.Println(header, conn)
		buf := bytes.Buffer{}
		io.Copy(&buf, conn)
		line = append(line, fmt.Sprintf("%s:%s", header, buf.String()))
	}
	fmt.Println(line)
}

func main() {
	flag.Parse()
	fmt.Println(flag.Args())

	headerClockAddrMap := make(map[string]string)

	var headerSlice []string

	for _, headerAddrStr := range flag.Args() {
		headerAddrStrSlice := strings.Split(headerAddrStr, "=")
		header := strings.TrimSpace(headerAddrStrSlice[0])
		addr := strings.TrimSpace(headerAddrStrSlice[1])
		headerSlice = append(headerSlice, header)
		headerClockAddrMap[header] = addr
	}

	fmt.Println(headerClockAddrMap)

	fmt.Println(strings.Join(headerSlice, "\t\t"))

	headerClockConnMap := make(map[string]net.Conn)
	for header, addr := range headerClockAddrMap {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		headerClockConnMap[header] = conn
	}
	mustCopy(headerClockConnMap)
	time.Sleep(10 * time.Second)
}
