#include "rocksdb/c.h"

#ifdef __cplusplus
extern "C" {
#endif
  extern int getNumLevels(const char*);
  extern void printDBName(rocksdb_t* db);
#ifdef __cplusplus
}
#endif
