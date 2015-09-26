package main

import (
	"io"
	"log"
	"net"
	"os"
)

const listenAddr = "localhost:4000"

func main() {
	var ()

	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		defer c.Close()
		consoleWriter := io.Writer(os.Stdout)
		d := io.MultiWriter(consoleWriter, c)
		if err != nil {
			log.Fatal(err)

		}
		io.Copy(d, c)
	}
}
