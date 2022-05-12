package util

import (
	"encoding/binary"
	"errors"
	"io"
)

var ErrShortRead = errors.New("short read")

// len + body
func Decode(r io.Reader) ([]byte, error) {
	var totalLen int32
	err := binary.Read(r, binary.BigEndian, &totalLen)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, totalLen-4)
	n, err := io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}
	if n != int(totalLen-4) {
		return nil, ErrShortRead
	}
	return buf, nil
}
