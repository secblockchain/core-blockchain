//go:build freebsd || darwin
// +build freebsd darwin

package reexec

// Self returns the path to the current process's binary.
func Self() string {
	return naiveSelf()
}
