package network

import (
	"log"

	"github.com/gogo/protobuf/proto"
	protos "github.com/sjljrvis/gArch/protos"
	types "github.com/sjljrvis/gArch/types"
)

func write(peer *types.Peer) {
	for {
		msg := <-peer.Msg
		_msg := &protos.Arc{}
		err := proto.Unmarshal(msg, _msg)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}

		switch _msg.GetType() {

		case "handshake":
			peer.Conn.Write(msg)

		default:
			log.Println("Default message")

		}
	}

}
