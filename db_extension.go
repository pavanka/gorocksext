package gorocksext

// #include <stdlib.h>
// #include "extensions.h"
import "C"

import (
	"errors"
	"unsafe"

	"github.com/tecbot/gorocksdb"
)

// NewIterators returns an array of iterators
// corresponding to the given column families for a given db
// iterators is allocated by this method -> responsibility of
// caller to destroy the underlying c objects
// opts, db, cfs will not be destroyed -> resonsibility of caller
func NewIterators(
	opts *gorocksdb.ReadOptions,
	db *gorocksdb.DB,
	cfHandles []*gorocksdb.ColumnFamilyHandle,
) ([]*gorocksdb.Iterator, error) {
	size := len(cfHandles)
	cCFHandles := make([]*C.rocksdb_column_family_handle_t, size, size)
	for i, cfHandle := range cfHandles {
		cCFHandles[i] = (*C.rocksdb_column_family_handle_t)(cfHandle.UnsafeGetCFHandler())
	}

	iters := make([]*C.rocksdb_iterator_t, size, size)
	var cErr *C.char
	C.get_iterators(
		(*C.rocksdb_readoptions_t)(opts.UnsafeGetReadOptions()),
		(*C.rocksdb_t)(db.UnsafeGetDB()),
		&cCFHandles[0],
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
		iterators = append(
			iterators,
			gorocksdb.NewNativeIterator(unsafe.Pointer(iter)),
		)
	}
	return iterators, nil
}

// CreateCheckpoint creates an openable snapshot of the db
// db is not destroyed -> responsibility of the caller
func CreateCheckpoint(db *gorocksdb.DB, checkpointDir string) error {
	cCheckpointDir := C.CString(checkpointDir)
	defer C.free(unsafe.Pointer(cCheckpointDir))
	var cErr *C.char
	C.create_checkpoint(
		(*C.rocksdb_t)(db.UnsafeGetDB()),
		cCheckpointDir,
		&cErr,
	)
	if cErr != nil {
		defer C.free(unsafe.Pointer(cErr))
		return errors.New(C.GoString(cErr))
	}
	return nil
}
