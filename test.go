package main

// #cgo LDFLAGS: -L./cpp -lfoo
// #cgo CFLAGS: -I./cpp
// #include <stdlib.h>
// #include "test.h"
// #include "rocksdb/c.h"
import "C"

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/tecbot/gorocksdb"
)

func printNumLevel() {
	cs := C.CString("/tmp/yolo")
	defer C.free(unsafe.Pointer(cs))
	val := int(C.getNumLevels(cs))
	fmt.Println(val)
}

func printDBName(name string) {
	opts := gorocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)
	db, err := gorocksdb.OpenDb(opts, name)
	if err != nil {
		log.Fatal(err)
	}
	c_db := (*C.struct_rocksdb_t)(db.GetDb())
	C.printDBName(c_db)
}

func main() {
	printDBName("/tmp/yolo")
}
