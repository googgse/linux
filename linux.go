package linux

import (
	// import built-in packages
	"os"
	"runtime"
	"strings"

	// import third-party packages
	"github.com/googgse/path"
)

// 是否为 Linux
func IsLinux() bool {
	// control flow
	if runtime.GOOS != "linux" {
		// return value
		return false
	} else {
		// return value
		return true
	}
}

// 查看 Linux 发行版本
func Release() string {
	// control flow
	if !IsLinux() {
		// return value
		return "notLinux"
	} else if path.Exists("/etc/redhat-release") {
		// return value
		return "centos"
	}
	bytes, err := os.ReadFile("/etc/issue")
	// control flow
	if err != nil {
		bytes, err := os.ReadFile("/proc/version")
		// control flow
		if err != nil {
			// return value
			return "unknown"
		} else if strings.Contains(strings.ToLower(string(bytes)), "debian") {
			// return value
			return "debian"
		} else if strings.Contains(strings.ToLower(string(bytes)), "ubuntu") {
			// return value
			return "ubuntu"
		} else if strings.Contains(strings.ToLower(string(bytes)), "centos") {
			// return value
			return "centos"
		} else if strings.Contains(strings.ToLower(string(bytes)), "redhat") {
			// return value
			return "redhat"
		} else if strings.Contains(strings.ToLower(string(bytes)), "red hat") {
			// return value
			return "redhat"
		} else if strings.Contains(strings.ToLower(string(bytes)), "suse") {
			// return value
			return "suse"
		} else if strings.Contains(strings.ToLower(string(bytes)), "kail") {
			// return value
			return "kail"
		}
	} else if strings.Contains(strings.ToLower(string(bytes)), "debian") {
		// return value
		return "debian"
	} else if strings.Contains(strings.ToLower(string(bytes)), "ubuntu") {
		// return value
		return "ubuntu"
	} else if strings.Contains(strings.ToLower(string(bytes)), "centos") {
		// return value
		return "centos"
	} else if strings.Contains(strings.ToLower(string(bytes)), "redhat") {
		// return value
		return "redhat"
	} else if strings.Contains(strings.ToLower(string(bytes)), "red hat") {
		// return value
		return "redhat"
	} else if strings.Contains(strings.ToLower(string(bytes)), "suse") {
		// return value
		return "suse"
	} else if strings.Contains(strings.ToLower(string(bytes)), "kail") {
		// return value
		return "kail"
	}
	// return value
	return "unknown"
}
