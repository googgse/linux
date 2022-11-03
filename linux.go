package linux

import (
	// import built-in packages
	"os"
	"regexp"
	"runtime"

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
	regDebian := regexp.MustCompile(`(?i)(^|([\s\t\n]+))(debian)($|([\s\t\n]+))`)
	regUbuntu := regexp.MustCompile(`(?i)(^|([\s\t\n]+))(ubuntu)($|([\s\t\n]+))`)
	regCentos := regexp.MustCompile(`(?i)(^|([\s\t\n]+))(centos)($|([\s\t\n]+))`)
	regRedHat := regexp.MustCompile(`(?i)(^|([\s\t\n]+))(redhat)($|([\s\t\n]+))`)
	regRedHat2 := regexp.MustCompile(`(?i)(^|([\s\t\n]+))(red hat)($|([\s\t\n]+))`)
	// control flow
	if err != nil {
		bytes, err := os.ReadFile("/proc/version")
		// control flow
		if err != nil {
			// return value
			return "unknown"
		} else if regDebian.MatchString(string(bytes)) {
			// return value
			return "debian"
		} else if regUbuntu.MatchString(string(bytes)) {
			// return value
			return "ubuntu"
		} else if regCentos.MatchString(string(bytes)) {
			// return value
			return "centos"
		} else if regRedHat.MatchString(string(bytes)) {
			// return value
			return "redhat"
		} else if regRedHat2.MatchString(string(bytes)) {
			// return value
			return "redhat"
		}
	} else if regUbuntu.MatchString(string(bytes)) {
		// return value
		return "ubuntu"
	} else if regCentos.MatchString(string(bytes)) {
		// return value
		return "centos"
	} else if regRedHat.MatchString(string(bytes)) {
		// return value
		return "redhat"
	} else if regRedHat2.MatchString(string(bytes)) {
		// return value
		return "redhat"
	}
	// return value
	return "unknown"
}
