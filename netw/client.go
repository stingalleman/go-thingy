package netw

import (
	"net"
)

func Connect(host string) (conn net.Conn, err error) {
	return net.Dial("tcp", host)
}
