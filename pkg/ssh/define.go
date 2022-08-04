package ssh

import "os"

const (
	ShellExit = "shell_exit"
)

type StatCallback = func(srcInfo os.FileInfo, srcErr error, dstInfo os.FileInfo, dstErr error)
