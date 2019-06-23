package network

import (
	"fmt"
	"log"

	"github.com/gogo/protobuf/proto"
	protos "github.com/sjljrvis/gArch/protos"
	types "github.com/sjljrvis/gArch/types"
)

func read(peer *types.Peer) {
	defer peer.Conn.Close()
	data := make([]byte, 1024)
	for {
		len, err := peer.Conn.Read(data)

		if err != nil {
			log.Println("Peer disconnected 2->", peer.IP)
			peer.Conn.Close()
			return
		}

		msg := &protos.Arc{}
		err = proto.Unmarshal(data[0:len], msg)

		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}

		switch msg.GetType() {
		case "handshake":
			peer.IP = (string(msg.GetData()))
			log.Println("Connected to Peer :->", string(msg.GetData()))
			data, err := peerStore.Set(peer.IP, peer)
			fmt.Sprintln(data, err)
		default:
			log.Println("Default message")

		}
	}
}
