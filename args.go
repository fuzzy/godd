package main

import (
	"os"
	"strconv"
	"strings"
)

func quickCheck(args []string) bool {
	for _, a := range args {
		for _, arg := range os.Args {
			if a == arg {
				return true
			}
		}
	}
	return false
}

func argCheck(arg string) string {
	for _, v := range os.Args {
		if arg == strings.Split(v, "=")[0] {
			return strings.Split(v, "=")[1]
		}
	}
	return ""
}

func intArgCheck(arg string) int {
	ret_s := argCheck(arg)
	ret := 0
	if len(ret_s) > 0 {
		ret, _ = strconv.Atoi(ret_s)
	}
	return ret
}

func ioArgCheck(arg string) GoddMover {
	for _, v := range os.Args {
		if arg == strings.Split(v, "=")[0] {
			d := strings.Split(v, "=")
			switch d[0] {
			case "if":
				if exists(d[1]) {
					return NewFileReader(d[1])
				}
			case "of":
				return NewFileWriter(d[1])
			}
		}
	}
	return GoddFileMover{}
}

func ArgParser() (*GoddOpts, GoddReader, GoddWriter) {
	retv := &GoddOpts{BlockSize: 1024576}
	retv.Quiet = quickCheck([]string{"q", "qu", "quiet"})
	retv.Progress = quickCheck([]string{"np", "nop", "nopr", "noprog", "noprogress"})
	retv.BlockSize = intArgCheck("bs")
	if retv.BlockSize == 0 {
		retv.BlockSize = 1024576
	}
	retv.Count = intArgCheck("count")
	retv.Skip = intArgCheck("skip")
	retv.Seek = intArgCheck("seek")
	input := ioArgCheck("if")
	retv.Size = input.Size()
	output := ioArgCheck("of")
	return retv, input, output
}
