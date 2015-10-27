#include "rocksdb/c.h"
#include "rocksdb/db.h"
#include "test.h"

using rocksdb::DB;

extern "C" {
  struct rocksdb_t{DB* rep;};
}

int getNumLevels(const char *name) {
  rocksdb_options_t *options = rocksdb_options_create();
  rocksdb_options_set_create_if_missing(options, 1);
  char* err = NULL;
  rocksdb_t* db = rocksdb_open(options, name, &err);
	return db->rep->NumberLevels();
}
