package types

import (
	"net"

	uuid "github.com/satori/go.uuid"
)

// Peer Struct
type Peer struct {
	Conn   net.Conn
	Msg    chan []byte
	IP     string
	ID     uuid.UUID
	Active bool
}
