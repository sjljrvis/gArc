package p2p

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	webrtc "github.com/keroserene/go-webrtc"
)

var peerConnection *webrtc.PeerConnection
var dataChannel *webrtc.DataChannel
var err error

func generateOfferJSON() {
	log.Println("Generating offer...")
	offer, err := peerConnection.CreateOffer() // blocking
	if err != nil {
		log.Println(err)
		return
	}
	peerConnection.SetLocalDescription(offer)
}

func generateAnswer() {
	fmt.Println("Generating answer...")
	answer, err := peerConnection.CreateAnswer() // blocking
	if err != nil {
		fmt.Println(err)
		return
	}
	peerConnection.SetLocalDescription(answer)
}

func receiveDescription(sdp *webrtc.SessionDescription) {
	err = peerConnection.SetRemoteDescription(sdp)
	if nil != err {
		fmt.Println("ERROR", err)
		return
	}
	fmt.Println("SDP " + sdp.Type + " successfully received.")
	if "offer" == sdp.Type {
		go generateAnswer()
	}
}

func registerPeer(msg string) {
	var parsed map[string]interface{}
	fmt.Println("=============================")
	fmt.Println(peerConnection)
	fmt.Println("=============================")

	err = json.Unmarshal([]byte(msg), &parsed)
	if nil != err {
		return
	}

	if nil != parsed["sdp"] {
		sdp := webrtc.DeserializeSessionDescription(msg)
		if nil == sdp {
			fmt.Println("Invalid SDP.")
			return
		}
		receiveDescription(sdp)
	}

	if nil != parsed["candidate"] {
		ice := webrtc.DeserializeIceCandidate(msg)
		if nil == ice {
			fmt.Println("Invalid ICE candidate.")
			return
		}
		peerConnection.AddIceCandidate(*ice)
		fmt.Println("ICE candidate successfully received.")
	}
}

func prepareDataChannel(channel *webrtc.DataChannel) {
	channel.OnOpen = func() {
		fmt.Println("Data Channel Opened!")
	}
	channel.OnClose = func() {
		fmt.Println("Data Channel closed.")
	}
	channel.OnMessage = func(msg []byte) {
		log.Println(">>>>>>", string(msg))
	}
}

// Start * peer node at host machine
func Start() {
	webrtc.SetLoggingVerbosity(1)
	reader := bufio.NewReader(os.Stdin)
	log.Println("Starting Peer Routing")
	config := webrtc.NewConfiguration(webrtc.OptionIceServer("stun:stun.l.google.com:19302"))
	peerConnection, err = webrtc.NewPeerConnection(config)
	if nil != err {
		log.Println("Failed to create PeerConnection.")
		return
	}

	peerConnection.OnNegotiationNeeded = func() {
		go generateOfferJSON()
	}
	peerConnection.OnIceComplete = func() {
		fmt.Println("Finished gathering ICE candidates.")
		sdp := peerConnection.LocalDescription().Serialize()
		log.Println(sdp)
		// Publish this JSON to peer discovery server
	}
	peerConnection.OnDataChannel = func(channel *webrtc.DataChannel) {
		fmt.Println("Datachannel established by remote... ", channel.Label())
		dataChannel = channel
		prepareDataChannel(channel)
	}

	dataChannel, err = peerConnection.CreateDataChannel("test")
	if nil != err {
		fmt.Println("Unexpected failure creating Channel.")
		return
	}

	for {
		text, _ := reader.ReadString('\n')
		registerPeer(text)
	}
}
