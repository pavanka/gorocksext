#include "rocksdb/c.h"

#ifdef __cplusplus
extern "C" {
#endif
extern void getIterators(
    rocksdb_readoptions_t* opts,
    rocksdb_t *db,
    rocksdb_column_family_handle_t** cfs,
    rocksdb_iterator_t** iters,
    int num,
    char** err);
#ifdef __cplusplus
}
#endif
