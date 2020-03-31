package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag         bool
	versionFlag      bool
	verboseClassFlag bool
	verboseInstFlag  bool
	cpOption         string
	XjreOption       string
	XmsOption        string
	class            string
	args             []string
}

func parseCmd() *Cmd {
	cmd := new(Cmd)
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose", false, "enable verbose output")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose:class", false, "enable verbose output")
	flag.BoolVar(&cmd.verboseInstFlag, "verbose:inst", false, "enable verbose output")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "E:/project/java/pebble/stone/out/production/pebble/pebble", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.StringVar(&cmd.XjreOption, "Xms", "", "path to jre")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		// cmd.class = args[0]
		cmd.class = "Test"
		cmd.args = args[1:]
	}
	cmd.class = "Test"
	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
