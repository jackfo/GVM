package cmd

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	HelpFlag    bool
	VersionFlag bool
	CpOption    string
	Class       string
	XjreOption  string
	Args        []string
}

func PrintUsage(){
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}

/**
*将相应的cmd设置对应的参数
 */
func ParseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = PrintUsage
	flag.BoolVar(&cmd.HelpFlag, "help", false, "输出帮助信息")
	//flag.BoolVar(&cmd.HelpFlag, "?", false, "输出帮助信息")
	flag.BoolVar(&cmd.VersionFlag, "version", false, "输出相应的版本且退出")
	flag.StringVar(&cmd.CpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.CpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "XjreOption", "", "jre")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.Class = args[0]
		cmd.Args = args[:1]
	}
	return cmd
}
