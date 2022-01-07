//go:build !(amd64 || arm64) || purego

// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"encoding/binary"

	"github.com/go-faster/errors"
)

var _ = binary.LittleEndian // clickHouse uses LittleEndian

// DecodeColumn decodes Int8 rows from *Reader.
func (c *ColInt8) DecodeColumn(r *Reader, rows int) error {
	if rows == 0 {
		return nil
	}
	data, err := r.ReadRaw(rows)
	if err != nil {
		return errors.Wrap(err, "read")
	}
	v := *c
	v = append(v, make([]int8, rows)...)
	for i := range data {
		v[i] = int8(data[i])
	}
	*c = v
	return nil
}

// EncodeColumn encodes Int8 rows to *Buffer.
func (c ColInt8) EncodeColumn(b *Buffer) {
	if len(c) == 0 {
		return
	}
	start := len(b.Buf)
	b.Buf = append(b.Buf, make([]byte, len(c))...)
	for i := range c {
		b.Buf[i+start] = uint8(c[i])
	}
}
