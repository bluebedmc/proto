package proto

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"sync"
)

// Packet is a structured Minecraft packet.
// It can be converted from/to a RawPacket.
type Packet interface {
	FromRaw(p *RawPacket) error
	ToRaw(p *RawPacket) error
}

// RawPacket is an unstructured Minecraft packet that contains id and raw data.
// It can be sent/received over Minecraft Network Protocol.
type RawPacket struct {
	ID   int32
	Data []byte
}

// NewRawPacket creates a new RawPacket
func NewRawPacket() *RawPacket {
	return &RawPacket{}
}

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

// Marshal encodes given types to the raw packet
func (p *RawPacket) Marshal(types ...Type) error {
	var buffer bytes.Buffer

	for _, t := range types {
		_, err := t.WriteTo(&buffer)
		if err != nil {
			return err
		}
	}

	p.Data = buffer.Bytes()

	return nil
}

// Unmarshal parses the raw packet and store the result in given types
func (p *RawPacket) Unmarshal(types ...Type) error {
	reader := bytes.NewReader(p.Data)
	for _, t := range types {
		_, err := t.ReadFrom(reader)
		if err != nil {
			return err
		}
	}

	return nil
}

// Pack packs the raw packet to the writer
func (p *RawPacket) Pack(writer io.Writer, threshold int) error {
	if threshold >= 0 {
		return p.packWithCompression(writer, threshold)
	}
	return p.packWithoutCompression(writer)
}

func (p *RawPacket) packWithoutCompression(writer io.Writer) error {

	buffer := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buffer)
	buffer.Reset()

	_, err := VarInt(p.ID).WriteTo(buffer)
	if err != nil {
		return err
	}

	totalLength := int(buffer.Len()) + len(p.Data)

	_, err = VarInt(totalLength).WriteTo(writer)
	if err != nil {
		return err
	}

	_, err = writer.Write(buffer.Bytes())
	if err != nil {
		return err
	}

	_, err = writer.Write(p.Data)
	if err != nil {
		return err
	}

	return nil
}

func (p *RawPacket) packWithCompression(writer io.Writer, threshold int) error {
	buffer := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buffer)
	buffer.Reset()

	if len(p.Data) < threshold {

		_, err := VarInt(0).WriteTo(buffer)
		if err != nil {
			return err
		}
		_, err = VarInt(p.ID).WriteTo(buffer)
		if err != nil {
			return err
		}
		_, err = buffer.Write(p.Data)
		if err != nil {
			return err
		}
		// Packet Length
		_, err = VarInt(buffer.Len()).WriteTo(writer)
		if err != nil {
			return err
		}
		// Data Length + Packet ID + Data
		_, err = buffer.WriteTo(writer)
		if err != nil {
			return err
		}

	} else {
		zlibWriter := zlib.NewWriter(buffer)

		idLength, err := VarInt(p.ID).WriteTo(zlibWriter)
		if err != nil {
			return err
		}

		dataLength, err := zlibWriter.Write(p.Data)
		if err != nil {
			return err
		}

		err = zlibWriter.Close()
		if err != nil {
			return err
		}

		totalDataLength := bufPool.Get().(*bytes.Buffer)
		defer bufPool.Put(&dataLength)
		totalDataLength.Reset()

		totalDataLengthLength, err := VarInt(int(idLength) + dataLength).WriteTo(totalDataLength)
		if err != nil {
			return err
		}

		// Packet Length
		_, err = VarInt(int(totalDataLengthLength) + buffer.Len()).WriteTo(writer)
		if err != nil {
			return err
		}
		// Data Length
		_, err = totalDataLength.WriteTo(writer)
		if err != nil {
			return err
		}
		// PacketID + Data
		_, err = buffer.WriteTo(writer)
		if err != nil {
			return err
		}
	}

	return nil
}

// Unpack unpacks the raw packet from the reader
func (p *RawPacket) Unpack(reader io.Reader, threshold int) error {
	if threshold >= 0 {
		return p.unpackWithCompression(reader, threshold)
	}
	return p.unpackWithoutCompression(reader)
}

func (p *RawPacket) unpackWithoutCompression(reader io.Reader) error {
	var length VarInt
	_, err := length.ReadFrom(reader)
	if err != nil {
		return err
	}

	var id VarInt
	idLength, err := id.ReadFrom(reader)
	if err != nil {
		return err
	}
	p.ID = int32(id)

	dataLength := int(length) - int(idLength)

	if cap(p.Data) < dataLength {
		p.Data = make([]byte, dataLength)
	} else {
		p.Data = p.Data[:dataLength]
	}

	_, err = io.ReadFull(reader, p.Data)
	if err != nil {
		return err
	}

	return nil
}

func (p *RawPacket) unpackWithCompression(reader io.Reader, threshold int) error {
	var length VarInt
	_, err := length.ReadFrom(reader)
	if err != nil {
		return err
	}

	buffer := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buffer)
	buffer.Reset()

	_, err = io.CopyN(buffer, reader, int64(length))
	if err != nil {
		return err
	}
	reader = bytes.NewReader(buffer.Bytes())

	var dataLength VarInt
	dataLengthLength, err := dataLength.ReadFrom(reader)
	if err != nil {
		return err
	}

	var id VarInt
	if dataLength != 0 {
		if int(dataLength) < threshold {
			return fmt.Errorf("compressed packet error: size of %d is below threshold of %d", dataLength, threshold)
		}

		const maxDataLength = 2097152
		if dataLength > maxDataLength {
			return fmt.Errorf("compressed packet error: size of %d is larger than protocol maximum of %d", dataLength, maxDataLength)
		}

		zlibReader, err := zlib.NewReader(reader)
		if err != nil {
			return err
		}
		defer zlibReader.Close()

		reader = zlibReader

		idLength, err := id.ReadFrom(reader)
		if err != nil {
			return err
		}

		dataLength -= VarInt(idLength)
	} else {
		idLength, err := id.ReadFrom(reader)
		if err != nil {
			return err
		}

		dataLength = VarInt(int64(length) - dataLengthLength - idLength)
	}

	if cap(p.Data) < int(dataLength) {
		p.Data = make([]byte, dataLength)
	} else {
		p.Data = p.Data[:dataLength]
	}

	p.ID = int32(id)

	_, err = io.ReadFull(reader, p.Data)
	if err != nil {
		return err
	}

	return nil
}
