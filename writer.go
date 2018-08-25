package bio

import (
	"encoding/binary"
	"io"
	"math"
	"time"
)

// BinaryWriter write to io.Writer with special byte order
type BinaryWriter struct {
	io.Writer
	order binary.ByteOrder
}

// NewBinaryWriter create new BinaryWriter
func NewBinaryWriter(w io.Writer) *BinaryWriter {
	return NewBinaryWriterOrder(w, binary.BigEndian)
}

// NewBinaryWriterOrder create new BinaryWriter with special byte order
func NewBinaryWriterOrder(w io.Writer, order binary.ByteOrder) *BinaryWriter {
	return &BinaryWriter{w, order}
}

// Bool write a bool
func (w BinaryWriter) Bool(value bool) (int, error) {
	var v uint8
	if value {
		v = 1
	}

	return w.UInt8(v)
}

// UInt8 write a uint8
func (w BinaryWriter) UInt8(value uint8) (int, error) {
	return w.Write([]byte{byte(value)})
}

// UInt16 write a uint16
func (w BinaryWriter) UInt16(value uint16) (int, error) {
	buffer := make([]byte, 2)
	w.order.PutUint16(buffer, value)

	return w.Write(buffer)
}

// UInt32 write a uint32
func (w BinaryWriter) UInt32(value uint32) (int, error) {
	buffer := make([]byte, 4)
	w.order.PutUint32(buffer, value)

	return w.Write(buffer)
}

// UInt64 write a uint64
func (w BinaryWriter) UInt64(value uint64) (int, error) {
	buffer := make([]byte, 8)
	w.order.PutUint64(buffer, value)

	return w.Write(buffer)
}

// Int8 write a int8
func (w BinaryWriter) Int8(value int8) (int, error) {
	return w.UInt8(uint8(value))
}

// Int16 write a int16
func (w BinaryWriter) Int16(value int16) (int, error) {
	return w.UInt16(uint16(value))
}

// Int32 write a int32
func (w BinaryWriter) Int32(value int32) (int, error) {
	return w.UInt32(uint32(value))
}

// Int64 write a int64
func (w BinaryWriter) Int64(value int64) (int, error) {
	return w.UInt64(uint64(value))
}

// Int write a int
func (w BinaryWriter) Int(value int) (int, error) {
	return w.UInt32(uint32(value))
}

// Float32 write a float32
func (w BinaryWriter) Float32(value float32) (int, error) {
	return w.UInt32(math.Float32bits(value))
}

// Float64 write a float64
func (w BinaryWriter) Float64(value float64) (int, error) {
	return w.UInt64(math.Float64bits(value))
}

// String write a string
func (w BinaryWriter) String(value string) (int, error) {
	buffer := []byte(value)
	_, err := w.Int(len(buffer))
	if err != nil {
		return 0, err
	}

	return w.Write(buffer)
}

// Time write a time.Time
func (w BinaryWriter) Time(value time.Time) (int, error) {
	_, err := w.UInt64(uint64(value.Unix()))
	if err != nil {
		return 0, err
	}

	return w.String(value.Location().String())
}
