package gorocksext

// #include <stdlib.h>
// #include "extensions.h"
import "C"

import (
	"errors"
	"unsafe"

	"github.com/tecbot/gorocksdb"
)

// PurgeOldBackups purges all old backups except the 'num' latest
// engine will not be destroyed -> responsibility of caller
func PurgeOldBackups(engine *gorocksdb.BackupEngine, num uint) error {
	var cErr *C.char
	C.purge_old_backups(
		(*C.rocksdb_backup_engine_t)(engine.UnsafeGetBackupEngine()),
		C.uint32_t(num),
		&cErr,
	)
	if cErr != nil {
		defer C.free(unsafe.Pointer(cErr))
		return errors.New(C.GoString(cErr))
	}
	return nil
}
