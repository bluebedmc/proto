package proto

import "fmt"

type wrongPacketErr struct {
	expect int32
	get    int32
}

func (w wrongPacketErr) Error() string {
	return fmt.Sprintf("wrong packet id: expect %#02X, get %#02X", w.expect, w.get)
}

func WrongPacketError(expect, get int32) error {
	return wrongPacketErr{expect, get}
}
