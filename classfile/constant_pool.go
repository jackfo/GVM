package classfile

//定义常量池,其实就是常量信息的一个集合
type ConstantPool []ConstantInfo


/**根据类阅读器,阅读常量池相关信息*/
func readConstantPool(reader *ClassReader) ConstantPool{

}


func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo{

}

func (self ConstantPool) getNameAndType(index uint16) (string,string){

}

func (self ConstantPool) getClassName(index uint16) string{

}

/**获取字节码以utf-8的形式*/
func (self ConstantPool) getUtf8(index uint16){
	
}