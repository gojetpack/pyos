package pyos

import "runtime"

var Path path

// The name of the operating system dependent module imported.
// The following names have currently been registered: 'posix', 'nt', 'java'.
// see: https://docs.python.org/3/library/os.html#os.name
func Name() string {
	switch runtime.GOOS {
	case "windows":
		return "nt"
	case "linux":
	case "freebsd":
	case "android":
	case "aix":
	case "darwin":
	case "dragonfly":
	case "hurd":
	case "illumos":
	case "js":
	case "nacl":
	case "netbsd":
	case "openbsd":
	case "plan9":
	case "solaris":
	case "zos":
		return "posix"
	}
	return ""
}

// Return the filename corresponding to the controlling terminal of the process.
// see: https://docs.python.org/3/library/os.html#os.ctermid
func Ctermid() (string, error) {
	return getTTYName()
}
