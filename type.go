package proto

import (
	"bytes"
	"errors"
	"io"
	"math"

	"github.com/google/uuid"
)

// Type represents a Minecraft protocol data type.
// Implements io.ReaderFrom and io.WriterTo.
type Type interface {
	io.ReaderFrom
	io.WriterTo
}

// --- Boolean ---

// Boolean is either false or true.
// Implements proto.Type interface (Minecraft protocol data type).
type Boolean bool

// ReadFrom reads Boolean data from r until an error occurs.
// The return value n is the number of bytes read.
// Any error encountered during the read is also returned.
func (b *Boolean) ReadFrom(r io.Reader) (n int64, err error) {
	v, err := readByte(r)
	if err != nil {
		return 1, err
	}

	*b = v != 0
	return 1, nil
}

// WriteTo writes Boolean data to w until an error occurs.
// The return value n is the number of bytes written.
// Any error encountered during the write is also returned.
func (b Boolean) WriteTo(w io.Writer) (int64, error) {
	var v byte
	if b {
		v = 0x01
	} else {
		v = 0x00
	}
	nn, err := w.Write([]byte{v})
	return int64(nn), err
}

// --- Byte ---

// Byte is a signed 8-bit integer, two's complement.
// Implements proto.Type interface (Minecraft protocol data type).
type Byte int8

// ReadFrom reads Byte data from r until an error occurs.
// The return value n is the number of bytes read.
// Any error encountered during the read is also returned.
func (b *Byte) ReadFrom(r io.Reader) (n int64, err error) {
	v, err := readByte(r)
	if err != nil {
		return 0, err
	}
	*b = Byte(v)
	return 1, nil
}

// WriteTo writes Byte data to w until an error occurs.
// The return value n is the number of bytes written.
// Any error encountered during the write is also returned.
func (b Byte) WriteTo(w io.Writer) (n int64, err error) {
	nn, err := w.Write([]byte{byte(b)})
	return int64(nn), err
}

// --- UsignedByte ---

// UnsignedByte is an unsigned 8-bit integer.
// Implements proto.Type interface (Minecraft protocol data type).
type UnsignedByte uint8

// ReadFrom reads UnsignedByte data from r until an error occurs.
// The return value n is the number of bytes read.
// Any error encountered during the read is also returned.
func (u *UnsignedByte) ReadFrom(r io.Reader) (n int64, err error) {
	v, err := readByte(r)
	if err != nil {
		return 0, err
	}
	*u = UnsignedByte(v)
	return 1, nil
}

// WriteTo writes UnsignedByte data to w until an error occurs.
// The return value n is the number of bytes written.
// Any error encountered during the write is also returned.
func (u UnsignedByte) WriteTo(w io.Writer) (n int64, err error) {
	nn, err := w.Write([]byte{byte(u)})
	return int64(nn), err
}

// --- Short ---

// Short is a signed 16-bit integer, two's complement.
// Implements proto.Type interface (Minecraft protocol data type).
type Short int16

// ReadFrom reads Short data from r until an error occurs.
// The return value n is the number of bytes read.
// Any error encountered during the read is also returned.
func (s *Short) ReadFrom(r io.Reader) (n int64, err error) {
	var bs [2]byte
	if nn, err := io.ReadFull(r, bs[:]); err != nil {
		return int64(nn), err
	} else {
		n += int64(nn)
	}

	*s = Short(int16(bs[0])<<8 | int16(bs[1]))
	return
}

// WriteTo writes Short data to w until an error occurs.
// The return value n is the number of bytes written.
// Any error encountered during the write is also returned.
func (s Short) WriteTo(w io.Writer) (int64, error) {
	n := uint16(s)
	nn, err := w.Write([]byte{byte(n >> 8), byte(n)})
	return int64(nn), err
}

// --- UnsignedShort ---

// UnsignedShort is an unsigned 16-bit integer.
// Implements proto.Type interface (Minecraft protocol data type).
type UnsignedShort uint16

// ReadFrom reads UnsignedShort data from r until an error occurs.
// The return value n is the number of bytes read.
// Any error encountered during the read is also returned.
func (us *UnsignedShort) ReadFrom(r io.Reader) (n int64, err error) {
	var bs [2]byte
	if nn, err := io.ReadFull(r, bs[:]); err != nil {
		return int64(nn), err
	} else {
		n += int64(nn)
	}

	*us = UnsignedShort(int16(bs[0])<<8 | int16(bs[1]))
	return
}

// WriteTo writes UnsignedShort data to w until an error occurs.
// The return value n is the number of bytes written.
// Any error encountered during the write is also returned.
func (us UnsignedShort) WriteTo(w io.Writer) (int64, error) {
	n := uint16(us)
	nn, err := w.Write([]byte{byte(n >> 8), byte(n)})
	return int64(nn), err
}

// --- Int ---

// Int is a signed 32-bit integer, two's complement.
// Implements proto.Type interface (Minecraft protocol data type).
type Int int32

// ReadFrom reads Int data from r until an error occurs.
// The return value n is the number of bytes read.
// Any error encountered during the read is also returned.
func (i *Int) ReadFrom(r io.Reader) (n int64, err error) {
	var bs [4]byte
	if nn, err := io.ReadFull(r, bs[:]); err != nil {
		return int64(nn), err
	} else {
		n += int64(nn)
	}

	*i = Int(int32(bs[0])<<24 | int32(bs[1])<<16 | int32(bs[2])<<8 | int32(bs[3]))
	return
}

// Write Int data to w
func (i Int) WriteTo(w io.Writer) (int64, error) {
	n := uint32(i)
	nn, err := w.Write([]byte{byte(n >> 24), byte(n >> 16), byte(n >> 8), byte(n)})
	return int64(nn), err
}

// --- Long ---

// Long is a signed 64-bit integer, two's complement.
// Implements proto.Type interface (Minecraft protocol data type).
type Long int64

// ReadFrom reads Long data from r until an error occurs.
// The return value n is the number of bytes read.
// Any error encountered during the read is also returned.
func (l *Long) ReadFrom(r io.Reader) (n int64, err error) {
	var bs [8]byte
	if nn, err := io.ReadFull(r, bs[:]); err != nil {
		return int64(nn), err
	} else {
		n += int64(nn)
	}

	*l = Long(int64(bs[0])<<56 | int64(bs[1])<<48 | int64(bs[2])<<40 | int64(bs[3])<<32 |
		int64(bs[4])<<24 | int64(bs[5])<<16 | int64(bs[6])<<8 | int64(bs[7]))
	return
}

// WriteTo writes Long data to w until an error occurs.
// The return value n is the number of bytes written.
// Any error encountered during the write is also returned.
func (l Long) WriteTo(w io.Writer) (int64, error) {
	n := uint64(l)
	nn, err := w.Write([]byte{
		byte(n >> 56), byte(n >> 48), byte(n >> 40), byte(n >> 32),
		byte(n >> 24), byte(n >> 16), byte(n >> 8), byte(n),
	})
	return int64(nn), err
}

// --- Float ---

// Float is a single-precision 32-bit IEEE 754 floating point number.
// Implements proto.Type interface (Minecraft protocol data type).
type Float float32

// ReadFrom reads Float data from r until an error occurs.
// The return value n is the number of bytes read.
// Any error encountered during the read is also returned.
func (f *Float) ReadFrom(r io.Reader) (n int64, err error) {
	var v Int

	n, err = v.ReadFrom(r)
	if err != nil {
		return
	}

	*f = Float(math.Float32frombits(uint32(v)))
	return
}

// WriteTo writes Float data to w until an error occurs.
// The return value n is the number of bytes written.
// Any error encountered during the write is also returned.
func (f Float) WriteTo(w io.Writer) (n int64, err error) {
	return Int(math.Float32bits(float32(f))).WriteTo(w)
}

// --- Double ---

// Double is a double-precision 64-bit IEEE 754 floating point number.
// Implements proto.Type interface (Minecraft protocol data type).
type Double float64

// ReadFrom reads Double data from r until an error occurs.
// The return value n is the number of bytes read.
// Any error encountered during the read is also returned.
func (d *Double) ReadFrom(r io.Reader) (n int64, err error) {
	var v Long
	n, err = v.ReadFrom(r)
	if err != nil {
		return
	}

	*d = Double(math.Float64frombits(uint64(v)))
	return
}

// WriteTo writes Double data to w until an error occurs.
// The return value n is the number of bytes written.
// Any error encountered during the write is also returned.
func (d Double) WriteTo(w io.Writer) (n int64, err error) {
	return Long(math.Float64bits(float64(d))).WriteTo(w)
}

// --- String ---

// String is a sequence of Unicode scalar values.
// Implements proto.Type interface (Minecraft protocol data type).
type String string

// ReadFrom reads String data from r until an error occurs.
// The return value n is the number of bytes read.
// Any error encountered during the read is also returned.
func (s *String) ReadFrom(r io.Reader) (n int64, err error) {
	var l VarInt //String length

	nn, err := l.ReadFrom(r)
	if err != nil {
		return nn, err
	}
	n += nn

	bs := make([]byte, l)
	if _, err := io.ReadFull(r, bs); err != nil {
		return n, err
	}
	n += int64(l)

	*s = String(bs)
	return n, nil
}

// WriteTo writes String data to w until an error occurs.
// The return value n is the number of bytes written.
// Any error encountered during the write is also returned.
func (s String) WriteTo(w io.Writer) (int64, error) {
	byteStr := []byte(s)
	n1, err := VarInt(len(byteStr)).WriteTo(w)
	if err != nil {
		return n1, err
	}
	n2, err := w.Write(byteStr)
	return n1 + int64(n2), err
}

// --- Chat ---

// Chat supports two-way chat communication.
// Implements proto.Type interface (Minecraft protocol data type).
type Chat = String

// --- Identifier ---

// Identifier is a namespaced location.
// Implements proto.Type interface (Minecraft protocol data type).
type Identifier = String

// --- VarInt ---

// VarInt is variable-length data encoding a two's complement signed 32-bit integer.
// Implements proto.Type interface (Minecraft protocol data type).
type VarInt int32

// MaxVarIntLen is Maximum VarInt length.
const MaxVarIntLen = 5

// ReadFrom reads VarInt data from r until an error occurs.
// The return value n is the number of bytes read.
// Any error encountered during the read is also returned.
func (v *VarInt) ReadFrom(r io.Reader) (n int64, err error) {
	var vi uint32
	for sec := byte(0x80); sec&0x80 != 0; n++ {
		if n > MaxVarIntLen {
			return n, errors.New("VarInt is too big")
		}

		sec, err = readByte(r)
		if err != nil {
			return n, err
		}

		vi |= uint32(sec&0x7F) << uint32(7*n)
	}

	*v = VarInt(vi)
	return
}

// WriteTo writes VarInt data to w until an error occurs.
// The return value n is the number of bytes written.
// Any error encountered during the write is also returned.
func (v VarInt) WriteTo(w io.Writer) (n int64, err error) {
	var vi = make([]byte, 0, MaxVarIntLen)
	num := uint32(v)
	for {
		b := num & 0x7F
		num >>= 7
		if num != 0 {
			b |= 0x80
		}
		vi = append(vi, byte(b))
		if num == 0 {
			break
		}
	}
	nn, err := w.Write(vi)
	return int64(nn), err
}

// --- VarLong ---

// VarLong is variable-length data encoding a two's complement signed 64-bit integer.
// Implements proto.Type interface (Minecraft protocol data type).
type VarLong int64

// MaxVarLongLen is Maximum VarLong length.
const MaxVarLongLen = 10

// ReadFrom reads VarLong data from r until an error occurs.
// The return value n is the number of bytes read.
// Any error encountered during the read is also returned.
func (v *VarLong) ReadFrom(r io.Reader) (n int64, err error) {
	var V uint64
	for sec := byte(0x80); sec&0x80 != 0; n++ {
		if n >= MaxVarLongLen {
			return n, errors.New("VarLong is too big")
		}
		sec, err = readByte(r)
		if err != nil {
			return
		}

		V |= uint64(sec&0x7F) << uint64(7*n)
	}

	*v = VarLong(V)
	return
}

// WriteTo writes VarLong data to w until an error occurs.
// The return value n is the number of bytes written.
// Any error encountered during the write is also returned.
func (v VarLong) WriteTo(w io.Writer) (n int64, err error) {
	var vi = make([]byte, 0, MaxVarLongLen)
	num := uint64(v)
	for {
		b := num & 0x7F
		num >>= 7
		if num != 0 {
			b |= 0x80
		}
		vi = append(vi, byte(b))
		if num == 0 {
			break
		}
	}
	nn, err := w.Write(vi)
	return int64(nn), err
}

// --- EntityMetadata ---

// EntityMetadata represents miscellaneous information about an entity
// Implements proto.Type interface (Minecraft protocol data type).
// WARNING: EntityMetadata is not implemented in proto.
type EntityMetadata struct{}

// ReadFrom reads EntityMetadata data from r until an error occurs.
// The return value n is the number of bytes read.
// Any error encountered during the read is also returned.
// WARNING: EntityMetadata is not implemented in proto.
func (e *EntityMetadata) ReadFrom(r io.Reader) (n int64, err error) {
	return 0, errors.New("proto.EntityMetadata is not implemented")
}

// WriteTo writes EntityMetadata data to w until an error occurs.
// The return value n is the number of bytes written.
// Any error encountered during the write is also returned.
// WARNING: EntityMetadata is not implemented in proto.
func (e *EntityMetadata) WriteTo(w io.Writer) (n int64, err error) {
	return 0, errors.New("proto.EntityMetadata is not implemented")
}

// --- Slot ---

// Slot represents an item stack in an inventory or container.
// Implements proto.Type interface (Minecraft protocol data type).
// WARNING: Slot is not implemented in proto.
type Slot struct{}

// ReadFrom reads Slot data from r until an error occurs.
// The return value n is the number of bytes read.
// Any error encountered during the read is also returned.
// WARNING: Slot is not implemented in proto.
func (s *Slot) ReadFrom(r io.Reader) (n int64, err error) {
	return 0, errors.New("proto.Slot is not implemented")
}

// WriteTo writes Slot data to w until an error occurs.
// The return value n is the number of bytes written.
// Any error encountered during the write is also returned.
// WARNING: Slot is not implemented in proto.
func (s *Slot) WriteTo(w io.Writer) (n int64, err error) {
	return 0, errors.New("proto.Slot is not implemented")
}

// --- NBTTag ---

// NBTTag represents an item and its associated data.
// Implements proto.Type interface (Minecraft protocol data type).
// WARNING: NBTTag is not implemented in proto.
type NBTTag struct{}

// ReadFrom reads NBTTag data from r until an error occurs.
// The return value n is the number of bytes read.
// Any error encountered during the read is also returned.
// WARNING: NBTTag is not implemented in proto.
func (t *NBTTag) ReadFrom(r io.Reader) (n int64, err error) {
	return 0, errors.New("proto.NBTTag is not implemented")
}

// WriteTo writes NBTTag data to w until an error occurs.
// The return value n is the number of bytes written.
// Any error encountered during the write is also returned.
// WARNING: NBTTag is not implemented in proto.
func (t *NBTTag) WriteTo(w io.Writer) (n int64, err error) {
	return 0, errors.New("proto.NBTTag is not implemented")
}

// --- Position ---

// Position is an integer/block position: x,y,z.
// Implements proto.Type interface (Minecraft protocol data type).
type Position struct {
	X, Y, Z int
}

// ReadFrom reads Position data from r until an error occurs.
// The return value n is the number of bytes read.
// Any error encountered during the read is also returned.
func (p *Position) ReadFrom(r io.Reader) (n int64, err error) {
	var v Long
	nn, err := v.ReadFrom(r)
	if err != nil {
		return nn, err
	}
	n += nn

	x := int(v >> 38)
	y := int(v & 0xFFF)
	z := int(v << 26 >> 38)

	if x >= 1<<25 {
		x -= 1 << 26
	}
	if y >= 1<<11 {
		y -= 1 << 12
	}
	if z >= 1<<25 {
		z -= 1 << 26
	}

	p.X, p.Y, p.Z = x, y, z
	return
}

// WriteTo writes Position data to w until an error occurs.
// The return value n is the number of bytes written.
// Any error encountered during the write is also returned.
func (p Position) WriteTo(w io.Writer) (n int64, err error) {
	var b [8]byte
	position := uint64(p.X&0x3FFFFFF)<<38 | uint64((p.Z&0x3FFFFFF)<<12) | uint64(p.Y&0xFFF)
	for i := 7; i >= 0; i-- {
		b[i] = byte(position)
		position >>= 8
	}
	nn, err := w.Write(b[:])
	return int64(nn), err
}

// --- Angle ---

// Angle is rotation angle in steps of 1/256 of a full turn.
// Implements proto.Type interface (Minecraft protocol data type).
type Angle Byte

// ReadFrom reads Angle data from r until an error occurs.
// The return value n is the number of bytes read.
// Any error encountered during the read is also returned.
func (a *Angle) ReadFrom(r io.Reader) (int64, error) {
	return (*Byte)(a).ReadFrom(r)
}

// WriteTo writes Angle data to w until an error occurs.
// The return value n is the number of bytes written.
// Any error encountered during the write is also returned.
func (a Angle) WriteTo(w io.Writer) (int64, error) {
	return Byte(a).WriteTo(w)
}

// ToDeg convert Angle to Degree
func (a Angle) ToDeg() float64 {
	return 360 * float64(a) / 256
}

// ToRad convert Angle to Radian
func (a Angle) ToRad() float64 {
	return 2 * math.Pi * float64(a) / 256
}

// --- UUID ---

// UUID is an uuid.
// Implements proto.Type interface (Minecraft protocol data type).
type UUID uuid.UUID

// ReadFrom reads UUID data from r until an error occurs.
// The return value n is the number of bytes read.
// Any error encountered during the read is also returned.
func (u *UUID) ReadFrom(r io.Reader) (n int64, err error) {
	nn, err := io.ReadFull(r, (*u)[:])
	return int64(nn), err
}

// WriteTo writes UUID data to w until an error occurs.
// The return value n is the number of bytes written.
// Any error encountered during the write is also returned.
func (u UUID) WriteTo(w io.Writer) (n int64, err error) {
	nn, err := w.Write(u[:])
	return int64(nn), err
}

// --- ByteArray ---

// ByteArray is just a sequence of zero or more bytes.
// Its meaning should be explained somewhere else (e.g. in the packet description).
// Implements proto.Type interface (Minecraft protocol data type).
type ByteArray []byte

// ReadFrom reads ByteArray data from r until an error occurs.
// The return value n is the number of bytes read.
// Any error encountered during the read is also returned.
func (b *ByteArray) ReadFrom(r io.Reader) (n int64, err error) {
	var Len VarInt
	n1, err := Len.ReadFrom(r)
	if err != nil {
		return n1, err
	}
	buf := bytes.NewBuffer(*b)
	buf.Reset()
	n2, err := io.CopyN(buf, r, int64(Len))
	*b = buf.Bytes()
	return n1 + n2, err
}

// WriteTo writes ByteArray data to w until an error occurs.
// The return value n is the number of bytes written.
// Any error encountered during the write is also returned.
func (b ByteArray) WriteTo(w io.Writer) (n int64, err error) {
	n1, err := VarInt(len(b)).WriteTo(w)
	if err != nil {
		return n1, err
	}
	n2, err := w.Write(b)
	return n1 + int64(n2), err
}
