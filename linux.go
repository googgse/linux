package linux

import (
	"runtime"
)

// 是否为 linux
func IsLinux() bool {
	if runtime.GOOS != "linux" {
		return false
	} else {
		return true
	}
}
