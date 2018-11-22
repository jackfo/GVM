package classpath

import (
	"path/filepath"
	"archive/zip"
	"io/ioutil"
	"errors"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

/**
* 1.打开相应的zip文件
* 2.
*/
func (self *ZipEntry) ReadClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err }
	//设置延时调用函数 关闭文件
	defer r.Close()

	//第一个_代表索引 第二个代表值
	for _, f := range r.File {
		if f.Name == className {
			//打开文件
			rc, err := f.Open()
			if err != nil {
			   return nil, nil, err
			}
			defer rc.Close()
			//读取文件相关数据
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}


func (self *ZipEntry) String() string {
	return self.absPath
}