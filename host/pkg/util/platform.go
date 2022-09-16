package util

import (
	"fmt"
	"runtime"
)

// OSArchPair is wrapper around operating system and architecture
type OSArchPair struct {
	OS, Arch string
}

// String converts environment into a string
func (p OSArchPair) String() string {
	return fmt.Sprintf("%s/%s", p.OS, p.Arch)
}

// OSArch returns the OS/arch combination to be used on the current system
func OSArch() OSArchPair {
	return OSArchPair{
		OS:   runtime.GOOS,
		Arch: runtime.GOARCH,
	}
}
