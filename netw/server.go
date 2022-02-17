package netw

import (
	"fmt"
	"net"
)

func Server(host string, port int) (listener net.Listener, err error) {
	listener, err = net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, err
	}

	return listener, err

	// for {
	// 	conn, err := l.Accept()
	// 	if err != nil {
	// 		fmt.Println("Error accepting: ", err.Error())
	// 		os.Exit(1)
	// 	}

	// 	go handleRequest(conn)
	// }
}

// func handleRequest(conn net.Conn) {
// 	// Make a buffer to hold incoming data.
// 	buf := make([]byte, 1024)

// 	_, err := conn.Read(buf)
// 	if err != nil {
// 		fmt.Println("Error reading:", err.Error())
// 	}

// 	fmt.Printf("%s\n", buf)
// 	conn.Write([]byte("send server"))
// 	conn.Close()
// }
