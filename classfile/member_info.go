package classfile

type MemberInfo struct {
	//保存常量池地址
	cp ConstantPool
	accessFlags uint16
	nameIndex uint16
	descriptorIndex uint16
	attributes []AttributeInfo
}

/**读取字段方法和数据*/
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount:=reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp) }
	return members
}

/***/
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:cp,
		accessFlags:reader.readUint16(),
		nameIndex:reader.readUint16(),
		descriptorIndex:reader.readUint16(),
		attributes:readAttributes(reader, cp),
	}
}


func (self *MemberInfo) AccessFlags() uint16 {

}
func (self *MemberInfo) Name() string {

}
func (self *MemberInfo) Descriptor() string {

}