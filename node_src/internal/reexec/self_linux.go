//go:build linux
// +build linux

package reexec

// Self returns the path to the current process's binary.
func Self() string {
	return "/proc/self/exe"
}
