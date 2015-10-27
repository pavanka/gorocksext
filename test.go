package main

// #cgo LDFLAGS: -L./cpp -lfoo
// #cgo CFLAGS: -I./cpp
// #include <stdlib.h>
// #include "test.h"
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	cs := C.CString("/tmp/yolo")
	defer C.free(unsafe.Pointer(cs))
	val := int(C.getNumLevels(cs))
	fmt.Println(val)
}
