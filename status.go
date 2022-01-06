package proto

import "fmt"

// --- Response ---

// Response is a packet that contains Server List Ping.
// Clientbound (S -> C)
// Implements proto.Packet interface.
type Response struct {
	JSONResponse String
}

// Response_ID is the Response packet ID.
const Response_ID = 0x00

// ToRaw marshals the Response Packet to the given RawPacket.
func (pi *Response) ToRaw(p *RawPacket) (err error) {
	p.ID = Response_ID
	return p.Marshal(&pi.JSONResponse)
}

// FromRaw unmarshals the Response Packet from the given RawPacket.
func (pi *Response) FromRaw(p *RawPacket) (err error) {
	if p.ID != Response_ID {
		return fmt.Errorf("invalid packet ID for Response: %d", p.ID)
	}
	return p.Unmarshal(&pi.JSONResponse)
}

// --- Pong ---

// Pong is a packet sent as a response to Ping.
// Clientbound (S -> C)
// Implements proto.Packet interface.
type Pong struct {
	Payload Long
}

// Pong_ID is the Pong packet ID.
const Pong_ID = 0x01

// ToRaw marshals the Pong Packet to the given RawPacket.
func (pi *Pong) ToRaw(p *RawPacket) (err error) {
	p.ID = Pong_ID
	return p.Marshal(&pi.Payload)
}

// FromRaw unmarshals the Pong Packet from the given RawPacket.
func (pi *Pong) FromRaw(p *RawPacket) (err error) {
	if p.ID != Pong_ID {
		return fmt.Errorf("invalid packet ID for Pong: %d", p.ID)
	}
	return p.Unmarshal(&pi.Payload)
}

// --- Request ---

// Request is a packet asking for Server List Ping.
// Serverbound (C -> S)
// Implements proto.Packet interface.
type Request struct {
}

// Request_ID is the Request packet ID.
const Request_ID = 0x00

// ToRaw marshals the Request Packet to the given RawPacket.
func (pi *Request) ToRaw(p *RawPacket) (err error) {
	p.ID = Request_ID
	return nil
}

// FromRaw unmarshals the Request Packet from the given RawPacket.
func (pi *Request) FromRaw(p *RawPacket) (err error) {
	if p.ID != Request_ID {
		return fmt.Errorf("invalid packet ID for Request: %d", p.ID)
	}
	return nil
}

// --- Ping ---

// Ping is a packet sent to ask a Pong response.
// Serverbound (C -> S)
// Implements proto.Packet interface.
type Ping struct {
	Payload Long
}

// Ping_ID is the Ping packet ID.
const Ping_ID = 0x01

// ToRaw marshals the Ping Packet to the given RawPacket.
func (pi *Ping) ToRaw(p *RawPacket) (err error) {
	p.ID = Ping_ID
	return p.Marshal(&pi.Payload)
}

// FromRaw unmarshals the Ping Packet from the given RawPacket.
func (pi *Ping) FromRaw(p *RawPacket) (err error) {
	if p.ID != Ping_ID {
		return fmt.Errorf("invalid packet ID for Ping: %d", p.ID)
	}
	return p.Unmarshal(&pi.Payload)
}
