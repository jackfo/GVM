package classpath
import "os"
import "path/filepath"

type Classpath struct {
	//引导类类加载器
	bootClasspath Entry
	//扩展类类加载器
	extClasspath Entry
	//用户类类加载器
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.ReadClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.ReadClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.ReadClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
    }
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

/**检测路径在操作系统中是否存在*/
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

/**解析用户类路径*/
func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
	    cpOption = "."
	}
	self.userClasspath = NewEntry(cpOption)
}