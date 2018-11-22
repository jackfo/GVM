package classpath

import (
	"os"
	"strings"
)

const pathListSeparator  = string(os.PathListSeparator)

/**表示类路径的接口*/
type Entry interface {
	ReadClass(className string) ([]byte,Entry,error)
	String() string
}


/**创建一个新的Entry*/
func NewEntry(path string)  Entry{
	//如果存在多个路径,创建组合Entry
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
    }

    //如果存在通配符,创建统配的Entry
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
    //如果是压缩文件,则采用压缩Entry
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
    }
    //创建类目录Entry
	return newDirEntry(path)
}