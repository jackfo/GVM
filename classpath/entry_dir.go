package classpath

import (
	"path/filepath"
	"io/ioutil"
)

/**目录形式类路径*/
type DirEntry struct {
	absDir string
}

/**将相对路径转化为决定路径*/
func newDirEntry(path string)  *DirEntry{
	absDir,err := filepath.Abs(path)
	if err!=nil{
		panic(err)
	}
	return &DirEntry{absDir}
}

/**根据路径读取相关数据*/
func (self *DirEntry)  ReadClass(className string)([]byte , Entry, error){
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string {
    return self.absDir
}