package main

// #cgo LDFLAGS: -L./cpp -lfoo
// #cgo CFLAGS: -I./cpp
// #include <stdlib.h>
// #include "test.h"
// #include "rocksdb/c.h"
import "C"

import (
	"errors"
	"unsafe"

	"github.com/pavanka/gorocksdb"
)

func NewIterators(opts *gorocksdb.ReadOptions, db *gorocksdb.DB, cfs []*gorocksdb.ColumnFamilyHandle) ([]*gorocksdb.Iterator, error) {
	size := len(cfs)
	cfsC := make([]*C.rocksdb_column_family_handle_t, size)
	for _, cf := range cfs {
		cfsC = append(cfsC, (*C.rocksdb_column_family_handle_t)(cf.UnsafeGetCFHandler()))
	}

	iters := make([]*C.rocksdb_iterator_t, size)
	var cErr *C.char
	C.getIterators(
		(*C.rocksdb_readoptions_t)(opts.UnsafeGetReadOptions()),
		(*C.rocksdb_t)(db.UnsafeGetDB()),
		&cfsC[0],
		&iters[0],
		C.int(size),
		&cErr,
	)
	if cErr != nil {
		defer C.free(unsafe.Pointer(cErr))
		return nil, errors.New(C.GoString(cErr))
	}

	var iterators []*gorocksdb.Iterator
	for _, iter := range iters {
		iterators = append(iterators, gorocksdb.NewNativeIterator(iter))
	}
	return iterators, nil
}

func main() {
}
