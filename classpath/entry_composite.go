package classpath

import (
	"strings"
	"errors"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := NewEntry(path)
		compositeEntry = append(compositeEntry, entry) }
	return compositeEntry
}

/**读取到第一个正常的字节码*/
func (self CompositeEntry) ReadClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, from, err := entry.ReadClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

/**遍历所有的Entry*/
func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String() }
	return strings.Join(strs, pathListSeparator)
}
