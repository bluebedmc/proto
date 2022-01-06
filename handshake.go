package proto

import (
	"fmt"
)

// --- Handshake ---

// Handshake is a packet that causes the server to switch into the target state.
// Serverbound (C -> S)
// Implements proto.Packet interface.
type Handshake struct {
	ProtocolVersion VarInt
	ServerAddress   String
	ServerPort      UnsignedShort
	NextState       VarInt
}

// Handshake_ID is the Handshake packet ID.
const Handshake_ID = 0x00

// ToRaw marshals the Handshake Packet to the given RawPacket.
func (h *Handshake) ToRaw(p *RawPacket) (err error) {
	p.ID = Handshake_ID
	return p.Marshal(&h.ProtocolVersion, &h.ServerAddress, &h.ServerPort, &h.NextState)
}

// FromRaw unmarshals the Handshake Packet from the given RawPacket.
func (h *Handshake) FromRaw(p *RawPacket) (err error) {
	if p.ID != Handshake_ID {
		return fmt.Errorf("invalid packet ID for Handshake: %d", p.ID)
	}
	return p.Unmarshal(&h.ProtocolVersion, &h.ServerAddress, &h.ServerPort, &h.NextState)
}

// --- LegacyServerPingList ---

// LegacyServerListPing is a packet asking for Server List Ping.
// While not technically part of the current protocol, modern servers should handle it correctly.
// Serverbound (C -> S)
// Implements proto.Packet interface.
type LegacyServerListPing struct {
	Payload UnsignedByte
}

// LegacyServerListPing_ID is the LegacyServerListPing packet ID.
const LegacyServerListPing_ID = 0xFE

// ToRaw marshals the LegacyServerListPing Packet to the given RawPacket.
func (l *LegacyServerListPing) ToRaw(p *RawPacket) (err error) {
	p.ID = LegacyServerListPing_ID
	return p.Marshal(&l.Payload)
}

// FromRaw unmarshals the LegacyServerListPing Packet from the given RawPacket.
func (l *LegacyServerListPing) FromRaw(p *RawPacket) (err error) {
	if p.ID != LegacyServerListPing_ID {
		return fmt.Errorf("invalid packet ID for LegacyServerListPing: %d", p.ID)
	}
	return p.Unmarshal(&l.Payload)
}
