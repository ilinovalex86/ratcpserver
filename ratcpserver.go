package ratcpserver

import (
	"fmt"
	"net"
)

func server(serverIpAddress string, connector func(conn net.Conn), serverName string) {
	const funcNameLog = "server(): "
	ln, err := net.Listen("tcp", serverIpAddress)
	if err != nil {
		ToLog(tcpServer, fmt.Sprintf("%s %s err: %s", serverName, funcNameLog, serverIpAddress), true)
	}
	fmt.Printf("start %s on: %s\n", serverName, serverIpAddress)
	for {
		conn, err := ln.Accept()
		if err != nil {
			ToLog(tcpServer, funcNameLog+"ln.Accept()", true)
		}
		fmt.Printf("Connect to %s: %s\n", serverName, conn.RemoteAddr())
		ToLog(tcpServer, fmt.Sprintf("%s %s connect: %s", serverName, funcNameLog, conn.RemoteAddr()), false)
		go connector(conn)
	}
}
