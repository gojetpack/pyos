package pyos

// #include <stdio.h>
// #include <stdlib.h>
// #include <termios.h>
// #include <unistd.h>
// #include <fcntl.h>
// #include <sys/ioctl.h>
//
// /* because golang doesn't like the ... param of ioctl */
// int ioctl_winsize(int d, unsigned long request, void *buf) {
//   return ioctl(d, request, buf);
// }
//
import "C"

import (
	"fmt"
	"unsafe"
)

// ctermid
func getTTYName() (string, error) {
	ttyname := C.ctermid((*C.char)(unsafe.Pointer(nil)))
	if p := (*C.char)(unsafe.Pointer(ttyname)); p == nil {
		err := fmt.Errorf("failed to get tty name")
		return "", err
	}
	return C.GoString(ttyname), nil
}
