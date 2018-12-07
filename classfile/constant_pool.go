package classfile

//定义常量池,其实就是常量信息的一个集合
type ConstantPool []ConstantInfo


/**根据类阅读器,阅读常量池相关信息*/
func readConstantPool(reader *ClassReader) ConstantPool{
    cpCount := int(reader.readUint16())
    cp := make([]ConstantInfo, cpCount)
    for i := 1; i < cpCount; i++ {
    	cp[i] = readConstantInfo(reader, cp)
    	switch cp[i].(type){
    		case *ConstantLongInfo, *ConstantDoubleInfo:
    			i++ 
    	}
    }
    return cp
}

/**根据索引查找常量*/
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo{
     if cpInfo := self[index]; cpInfo != nil{
     	return cpInfo
     }
     panic("Invalid constant pool index!")
}

/**从常量池查找字段或方法的名字和描述符*/
func (self ConstantPool) getNameAndType(index uint16) (string,string){
     ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
     name := self.getUtf8(ntInfo.nameIndex)
     _type := self.getUtf8(ntInfo.descriptorIndex)
     return name, _type
}

func (self ConstantPool) getClassName(index uint16) string{
    classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
    return self.getUtf8(classInfo.nameIndex)
}

/**获取字节码以utf-8的形式*/
func (self ConstantPool) getUtf8(index uint16){
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}