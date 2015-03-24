package main

import (
	"fmt"
	"os"
	"net"
	"time"
)

func main() {
	port := ":1336"
	udpAddress, err := net.ResolveUDPAddr("udp4", port)

	if err != nil {
		fmt.Println("error resolving UDP address on ", port)
		fmt.Println(err)
		return
	}

	conn, err := net.ListenUDP("udp", udpAddress)

	if err != nil {
		fmt.Println("error listening on UDP port ", port)
		fmt.Println(err)
		return
	}

	defer conn.Close()
	var buf []byte = make([]byte, 1318080)

	f, _ := os.Create("/opt/sse/responses.log")
	defer f.Close()

	for {
		time.Sleep(100 * time.Millisecond)
		n, address, err := conn.ReadFromUDP(buf)

		if err != nil {
			fmt.Println("error reading data from connection")
			fmt.Println(err)
			return
		}

		if address != nil {
			fmt.Println("got message from ", address, " with n = ", n)

			if n > 0 {
				f.Write([]byte(buf[0:n]))
                f.Write([]byte("\n"))
				f.Sync()
				fmt.Println("from address", address, "got message:", string(buf[0:n]), n)
			}
		}
	}

}
