package bio

import (
	"encoding/binary"
	"io"
	"math"
	"time"
)

// BinaryReader read from io.Reader with special byte order
type BinaryReader struct {
	io.Reader
	order binary.ByteOrder
}

// NewBinaryReader create new BinaryReader
func NewBinaryReader(r io.Reader) *BinaryReader {
	return NewBinaryReaderOrder(r, binary.BigEndian)
}

// NewBinaryReaderOrder create new BinaryReader with special byte order
func NewBinaryReaderOrder(r io.Reader, order binary.ByteOrder) *BinaryReader {
	return &BinaryReader{r, order}
}

// Bool read a bool
func (r BinaryReader) Bool() (bool, error) {
	value, err := r.UInt8()
	if err != nil {
		return false, err
	}

	if value > 0 {
		return true, nil
	}

	return false, nil
}

// Bytes read n bytes
func (r BinaryReader) Bytes(n int) ([]byte, error) {
	buffer := make([]byte, n)
	_, err := r.Read(buffer)
	return buffer, err
}

// UInt8 read a uint8
func (r BinaryReader) UInt8() (uint8, error) {
	buffer, err := r.Bytes(1)
	if err != nil {
		return 0, err
	}

	return uint8(buffer[0]), nil
}

// UInt16 read a uint16
func (r BinaryReader) UInt16() (uint16, error) {
	buffer, err := r.Bytes(2)
	if err != nil {
		return 0, err
	}

	return r.order.Uint16(buffer), nil
}

// UInt32 read a uint32
func (r BinaryReader) UInt32() (uint32, error) {
	buffer, err := r.Bytes(4)
	if err != nil {
		return 0, err
	}

	return r.order.Uint32(buffer), nil
}

// UInt64 read a uint64
func (r BinaryReader) UInt64() (uint64, error) {
	buffer, err := r.Bytes(8)
	if err != nil {
		return 0, err
	}

	return r.order.Uint64(buffer), nil
}

// Int8 read a int8
func (r BinaryReader) Int8() (int8, error) {
	value, err := r.UInt8()
	return int8(value), err
}

// Int16 read a int16
func (r BinaryReader) Int16() (int16, error) {
	value, err := r.UInt16()
	return int16(value), err
}

// Int32 read a int32
func (r BinaryReader) Int32() (int32, error) {
	value, err := r.UInt32()
	return int32(value), err
}

// Int64 read a int64
func (r BinaryReader) Int64() (int64, error) {
	value, err := r.UInt64()
	return int64(value), err
}

// Int read a int
func (r BinaryReader) Int() (int, error) {
	value, err := r.Int32()
	return int(value), err
}

// Float32 read a float32
func (r BinaryReader) Float32() (float32, error) {
	value, err := r.UInt32()
	if err != nil {
		return 0, nil
	}

	return math.Float32frombits(value), nil
}

// Float64 read a float64
func (r BinaryReader) Float64() (float64, error) {
	value, err := r.UInt64()
	if err != nil {
		return 0, nil
	}

	return math.Float64frombits(value), nil
}

// String read a string
func (r BinaryReader) String() (string, error) {
	size, err := r.Int()
	if err != nil {
		return "", err
	}

	buffer := make([]byte, size)
	_, err = r.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}

// Time read a time.Time
func (r BinaryReader) Time() (time.Time, error) {
	value := time.Time{}

	size, err := r.UInt8()
	if err != nil {
		return value, err
	}

	buffer := make([]byte, size)
	_, err = r.Read(buffer)
	if err != nil {
		return value, err
	}

	err = value.UnmarshalBinary(buffer)
	if err != nil {
		return value, err
	}

	return value, nil
}

func (r BinaryReader) ByteOrder() binary.ByteOrder {
	return r.order
}
