package types

import "net"

// Peer Struct
type Peer struct {
	Conn   net.Conn
	Msg    chan []byte
	IP     string
	Active bool
}
