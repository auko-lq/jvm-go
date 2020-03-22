package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

func (self *ClassReader) readUint8() uint8 {
	// u1
	value := self.data[0]
	self.data = self.data[1:]
	return value
}

func (self *ClassReader) readUint16() uint16 {
	// u2
	value := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return value
}

func (self *ClassReader) readUint32() uint32 {
	// u4
	value := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return value
}

func (self *ClassReader) readUint64() uint64 {
	value := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return value
}

func (self *ClassReader) readUint16s() []uint16 {
	// uint16 table
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

func (self *ClassReader) readBytes(length uint32) []byte {
	bytes := self.data[:length]
	self.data = self.data[length:]
	return bytes
}
