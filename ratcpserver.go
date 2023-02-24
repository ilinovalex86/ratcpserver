package ratcpserver

import (
	"fmt"
	"net"
)

func server(serverIpAddress string, connector func(conn net.Conn), serverName string) {
	ln, err := net.Listen("tcp", serverIpAddress)
	check(err)
	fmt.Printf("start %s on: %s\n", serverName, serverIpAddress)
	for {
		conn, err := ln.Accept()
		check(err)
		fmt.Printf("Connect to %s: %s\n", serverName, conn.RemoteAddr())
		go connector(conn)
	}
}
