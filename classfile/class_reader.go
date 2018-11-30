package classfile

import "encoding/binary"

/**字节码数据读取器*/
type ClassReader struct {
	data []byte
}
//u1
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	//将数组第一个元素干掉,保留后面元素到新数组
	self.data = self.data[1:]
	return val
}

//u2
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}


//u4
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16() }
	return s
}


/**读取指定数量的字节*/
func (self *ClassReader) readBytes(n uint32) []byte { bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}