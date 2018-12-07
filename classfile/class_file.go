package classfile

import "fmt"

type ClassFile struct {
	magic uint32
    minorVersion uint16
    majorVersion uint16
    constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

/**
* 将二进制文件进行解析,解析了读入到classfile
*/
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r:=recover();r!=nil{
			var ok bool
			err,ok =r.(error)
			if !ok{
				err = fmt.Errorf("%V",r)
			}
		}
	}()
	cr:=&ClassReader{classData}
	cf =&ClassFile{}
	cf.read(cr)
	return
}

func (self *ClassFile) read(reader *ClassReader) {

}
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic :=reader.readUint32()
	if magic !=0xCAFEBABE{
		panic("java.lang.ClassFormatter:magic")
	}
}

/**
*读取并检查版本信息
*/
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45: return
	case 46, 47, 48, 49, 50, 51, 52: if self.minorVersion == 0 {
		return }
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (self *ClassFile) MinorVersion() uint16 {

} // getter
func (self *ClassFile) MajorVersion() uint16 {

} // getter
func (self *ClassFile) ConstantPool() ConstantPool {

} // getter func (self *ClassFile) AccessFlags() uint16 {...} // getter
func (self *ClassFile) Fields() []*MemberInfo {

} // getter
func (self *ClassFile) Methods() []*MemberInfo {

} // getter
func (self *ClassFile) ClassName() string {

}
func (self *ClassFile) SuperClassName() string {

}
func (self *ClassFile) InterfaceNames() []string {

}
