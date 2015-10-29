#include "rocksdb/c.h"

#ifdef __cplusplus
extern "C" {
#endif
  extern void get_iterators(
      rocksdb_readoptions_t* opts,
      rocksdb_t *db,
      rocksdb_column_family_handle_t** cfs,
      rocksdb_iterator_t** iters,
      int num,
      char** errptr);

  extern void purge_old_backups(
      rocksdb_backup_engine_t* backup_engine,
      uint32_t num_backups_to_keep,
      char** errptr);

  extern void create_checkpoint(
      rocksdb_t *db,
      const char* checkpoint_dir,
      char **errptr);
#ifdef __cplusplus
}
#endif
