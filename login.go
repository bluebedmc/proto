package proto

import "fmt"

// --- LoginDisconnect ---

// LoginDisconnect is a packet that tells the user they have been disconnected.
// Clientbound (S -> C)
// Implements proto.Packet interface.
type LoginDisconnect struct {
	Reason Chat
}

// LoginDisconnect_ID is the LoginDisconnect packet ID.
const LoginDisconnect_ID = 0x00

// ToRaw marshals the LoginDisconnect Packet to the given RawPacket.
func (pi *LoginDisconnect) ToRaw(p *RawPacket) (err error) {
	p.ID = LoginDisconnect_ID
	return p.Marshal(&pi.Reason)
}

// FromRaw unmarshals the LoginDisconnect Packet from the given RawPacket.
func (pi *LoginDisconnect) FromRaw(p *RawPacket) (err error) {
	if p.ID != LoginDisconnect_ID {
		return fmt.Errorf("invalid packet ID for LoginDisconnect: %d", p.ID)
	}
	return p.Unmarshal(&pi.Reason)
}

// --- EncryptionRequest ---

// EncryptionRequest is a packet that initiate encryption process.
// Clientbound (S -> C)
// Implements proto.Packet interface.
type EncryptionRequest struct {
	ServerID          String
	PublicKeyLength   VarInt
	PublicKey         ByteArray
	VerifyTokenLength VarInt
	VerifyToken       ByteArray
}

// EncrpytionRequest_ID is the EncryptionRequest packet ID.
const EncryptionRequest_ID = 0x01

// ToRaw marshals the EncryptionRequest Packet to the given RawPacket.
func (pi *EncryptionRequest) ToRaw(p *RawPacket) (err error) {
	p.ID = EncryptionRequest_ID
	return p.Marshal(&pi.ServerID, &pi.PublicKeyLength, &pi.PublicKey, &pi.VerifyTokenLength, &pi.VerifyToken)
}

// FromRaw unmarshals the EncryptionRequest Packet from the given RawPacket.
func (pi *EncryptionRequest) FromRaw(p *RawPacket) (err error) {
	if p.ID != EncryptionRequest_ID {
		return fmt.Errorf("invalid packet ID for EncryptionRequest: %d", p.ID)
	}
	return p.Unmarshal(&pi.ServerID, &pi.PublicKeyLength, &pi.PublicKey, &pi.VerifyTokenLength, &pi.VerifyToken)
}

// --- LoginSuccess ---

// LoginSuccess is a packet that tells the client they have successfully logged in.
// Clientbound (S -> C)
// Implements proto.Packet interface.
type LoginSuccess struct {
	UUID     UUID
	Username String
}

// LoginSuccess_ID is the LoginSuccess packet ID.
const LoginSuccess_ID = 0x02

// ToRaw marshals the LoginSuccess Packet to the given RawPacket.
func (pi *LoginSuccess) ToRaw(p *RawPacket) (err error) {
	p.ID = LoginSuccess_ID
	return p.Marshal(&pi.UUID, &pi.Username)
}

// FromRaw unmarshals the LoginSuccess Packet from the given RawPacket.
func (pi *LoginSuccess) FromRaw(p *RawPacket) (err error) {
	if p.ID != LoginSuccess_ID {
		return fmt.Errorf("invalid packet ID for LoginSuccess: %d", p.ID)
	}
	return p.Unmarshal(&pi.UUID, &pi.Username)

}

// --- SetCompression ---

// SetCompression is a packet that tells the client to use compression.
// Clientbound (S -> C)
// Implements proto.Packet interface.
type SetCompression struct {
	Threshold VarInt
}

// SetCompression_ID is the SetCompression packet ID.
const SetCompression_ID = 0x03

// ToRaw marshals the SetCompression Packet to the given RawPacket.
func (pi *SetCompression) ToRaw(p *RawPacket) (err error) {
	p.ID = SetCompression_ID
	return p.Marshal(&pi.Threshold)
}

// FromRaw unmarshals the SetCompression Packet from the given RawPacket.
func (pi *SetCompression) FromRaw(p *RawPacket) (err error) {
	if p.ID != SetCompression_ID {
		return fmt.Errorf("invalid packet ID for SetCompression: %d", p.ID)
	}
	return p.Unmarshal(&pi.Threshold)
}

// --- LoginPluginRequest ---

// LoginPluginRequest is a packet used to implement a custom handshaking flow together with LoginPluginResponse.
// Clientbound (S -> C)
// Implements proto.Packet interface.
type LoginPluginRequest struct {
	MessageID VarInt
	Channel   Identifier
	Data      ByteArray
}

// LoginPluginResponse_ID is the LoginPluginResponse packet ID.
const LoginPluginRequest_ID = 0x04

// ToRaw marshals the LoginPluginRequest Packet to the given RawPacket.
func (pi *LoginPluginRequest) ToRaw(p *RawPacket) (err error) {
	p.ID = LoginPluginRequest_ID
	return p.Marshal(&pi.MessageID, &pi.Channel, &pi.Data)
}

// FromRaw unmarshals the LoginPluginRequest Packet from the given RawPacket.
func (pi *LoginPluginRequest) FromRaw(p *RawPacket) (err error) {
	if p.ID != LoginPluginRequest_ID {
		return fmt.Errorf("invalid packet ID for LoginPluginRequest: %d", p.ID)
	}
	return p.Unmarshal(&pi.MessageID, &pi.Channel, &pi.Data)
}

// --- LoginStart ---

// LoginStart is a packet sent by client to initiate the login process.
// Serverbound (C -> S)
// Implements proto.Packet interface.
type LoginStart struct {
	Name String
}

// LoginStart_ID is the LoginStart packet ID.
const LoginStart_ID = 0x00

// ToRaw marshals the LoginStart Packet to the given RawPacket.
func (pi *LoginStart) ToRaw(p *RawPacket) (err error) {
	p.ID = LoginStart_ID
	return p.Marshal(&pi.Name)
}

// FromRaw unmarshals the LoginStart Packet from the given RawPacket.
func (pi *LoginStart) FromRaw(p *RawPacket) (err error) {
	if p.ID != LoginStart_ID {
		return fmt.Errorf("invalid packet ID for LoginStart: %d", p.ID)
	}
	return p.Unmarshal(&pi.Name)
}

// --- EncryptionResponse ---

// EncryptionResponse is a packet sent by server to confirm the encryption process.
// Serverbound (C -> S)
// Implements proto.Packet interface.
type EncryptionResponse struct {
	SharedSecretLength VarInt
	SharedSecret       ByteArray
	VerifyTokenLength  VarInt
	VerifyToken        ByteArray
}

// EncrpytionResponse_ID is the EncryptionResponse packet ID.
const EncryptionResponse_ID = 0x01

// ToRaw marshals the EncryptionResponse Packet to the given RawPacket.
func (pi *EncryptionResponse) ToRaw(p *RawPacket) (err error) {
	p.ID = EncryptionResponse_ID
	return p.Marshal(&pi.SharedSecretLength, &pi.SharedSecret, &pi.VerifyTokenLength, &pi.VerifyToken)
}

// FromRaw unmarshals the EncryptionResponse Packet from the given RawPacket.
func (pi *EncryptionResponse) FromRaw(p *RawPacket) (err error) {
	if p.ID != EncryptionResponse_ID {
		return fmt.Errorf("invalid packet ID for EncryptionResponse: %d", p.ID)
	}
	return p.Unmarshal(&pi.SharedSecretLength, &pi.SharedSecret, &pi.VerifyTokenLength, &pi.VerifyToken)
}

// --- LoginPluginResponse ---

// LoginPluginResponse is a packet used to implement a custom handshaking flow together with LoginPluginResponse.
// Serverbound (C -> S)
// Implements proto.Packet interface.
type LoginPluginResponse struct {
	MessageID  VarInt
	Successful Boolean
	Data       ByteArray
}

// LoginPluginResponse_ID is the LoginPluginResponse packet ID.
const LoginPluginResponse_ID = 0x04

// ToRaw marshals the LoginPluginResponse Packet to the given RawPacket.
func (pi *LoginPluginResponse) ToRaw(p *RawPacket) (err error) {
	p.ID = LoginPluginResponse_ID
	return p.Marshal(&pi.MessageID, &pi.Successful, &pi.Data)
}

// FromRaw unmarshals the LoginPluginResponse Packet from the given RawPacket.
func (pi *LoginPluginResponse) FromRaw(p *RawPacket) (err error) {
	if p.ID != LoginPluginResponse_ID {
		return fmt.Errorf("invalid packet ID for LoginPluginResponse: %d", p.ID)
	}
	return p.Unmarshal(&pi.MessageID, &pi.Successful, &pi.Data)
}
