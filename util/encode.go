package util

import (
	"encoding/binary"
	"errors"
	"io"
)

var ErrShortWrite = errors.New("short write")

func Encode(w io.Writer, buf []byte) error {
	var totalLen int32 = int32(len(buf)) + 4

	err := binary.Write(w, binary.BigEndian, &totalLen)
	if err != nil {
		return err
	}
	n, err := w.Write([]byte(buf)) // write the frame payload to outbound stream
	if err != nil {
		return err
	}
	if n != len(buf) {
		return ErrShortWrite
	}
	return nil
}
