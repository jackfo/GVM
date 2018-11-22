package main

import (
	"GVM/cmd"
	"GVM/classpath"
	"fmt"
	"strings"
)

/**启动JVM*/
func startJVM(cmd *cmd.Cmd) {

	//解析到相应的XjreOption和CpOption
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)

	fmt.Printf("classpath:%v class:%v args:%v\n",
		cp, cmd.Class, cmd.Args)
	className := strings.Replace(cmd.Class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.Class)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}

func main() {
	cmdLine := cmd.ParseCmd()
	if cmdLine.VersionFlag {
		fmt.Print("version 0.0.1")
	} else if cmdLine.HelpFlag || cmdLine.Class == "" {
		cmd.PrintUsage()
	} else {
		startJVM(cmdLine)
	}
}
