package main

import (
	"flag"
	"fmt"

	"github.com/kaleofeng/gometazion/pkg/homing"
	"github.com/kaleofeng/gometazion/pkg/kit/file"
)

type CmdParams struct {
	srcDir         string
	dstDir         string
	srcFilePattern string
	dstFilePattern string
}

func Usage() {
	fmt.Println("Usage: homing [OPTION]... SOURCE DEST")
	fmt.Println("Load key-values from ini files in SOURCE directory")
	fmt.Println("and replace placeholders in text files in DEST directory.")
	fmt.Printf("  %s\t%s\n", "-s", "source ini file pattern, default .ini$")
	fmt.Printf("  %s\t%s\n", "-d", "dest text file pattern, default .*")
}

func ParseParams(params *CmdParams) bool {
	flag.Usage = Usage

	srcFilePattern := flag.String("s", ".ini$", "source ini file pattern, default .ini$")
	dstFilePattern := flag.String("d", ".*", "dest text file pattern, default .*")
	flag.Parse()

	restArgs := flag.Args()
	restArgNum := flag.NArg()
	if restArgNum < 2 {
		flag.Usage()
		return false
	}

	srcDir := restArgs[0]
	dstDir := restArgs[1]

	isSrcDir, _ := file.IsDir(srcDir)
	isDstDir, _ := file.IsDir(dstDir)
	if !isSrcDir || !isDstDir {
		flag.Usage()
		return false
	}

	params.srcDir = srcDir
	params.dstDir = dstDir
	params.srcFilePattern = *srcFilePattern
	params.dstFilePattern = *dstFilePattern
	return true
}

func main() {
	var cmdParams CmdParams
	if !ParseParams(&cmdParams) {
		return
	}

	fmt.Printf("Command params: %+v\n", cmdParams)

	aide := homing.NewAide()

	srcFilePaths, err := file.ListFiles(cmdParams.srcDir, cmdParams.srcFilePattern)
	if err != nil {
		return
	}

	dstFilePaths, err := file.ListFiles(cmdParams.dstDir, cmdParams.dstFilePattern)
	if err != nil {
		return
	}

	for _, filePath := range srcFilePaths {
		fmt.Printf("Load from file[%s]\n", filePath)
		err = aide.LoadFromIni(filePath)
		if err != nil {
			fmt.Printf("Error Occurred[%v] while load file[%s]\n", err, filePath)
		}
	}

	for _, filePath := range dstFilePaths {
		fmt.Printf("Repalce text file[%s]\n", filePath)
		err = aide.ReplaceTextFile(filePath)
		if err != nil {
			fmt.Printf("Error Occurred[%v] while replace file[%s]\n", err, filePath)
		}
	}
}
