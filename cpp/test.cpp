#include <iostream>
#include "test.h"
#include "rocksdb/db.h"

using namespace std;
using rocksdb::DB;
using rocksdb::ColumnFamilyHandle;
using rocksdb::ReadOptions;
using rocksdb::Slice;
using rocksdb::Status;
using rocksdb::Iterator;

extern "C" {
  struct rocksdb_t{DB* rep;};
  struct rocksdb_readoptions_t{
    ReadOptions rep;
    Slice upper_bound; // stack variable to set pointer to in ReadOptions
  };
  struct rocksdb_column_family_handle_t{ColumnFamilyHandle* rep;};
  struct rocksdb_iterator_t{Iterator* rep;};
}

// copied from c.cc rocksdb
static bool SaveError(char** errptr, const Status& s) {
  assert(errptr != nullptr);
  if (s.ok()) {
    return false;
  } else if (*errptr == nullptr) {
    *errptr = strdup(s.ToString().c_str());
  } else {
    // TODO(sanjay): Merge with existing error?
    // This is a bug if *errptr is not created by malloc()
    free(*errptr);
    *errptr = strdup(s.ToString().c_str());
  }
  return true;
}

 void getIterators(
    rocksdb_readoptions_t* opts,
    rocksdb_t *db,
    rocksdb_column_family_handle_t** cfs,
    rocksdb_iterator_t** iters,
    int num,
    char** err) {
  vector<ColumnFamilyHandle*> cfs_vec(num);
  for (int i = 0; i < num; i++) {
    cfs_vec.push_back(cfs[i]->rep);
  }
  vector<Iterator*> res;

  Status status = db->rep->NewIterators(opts->rep,
      cfs_vec,
      &res);
  assert(res.size() == num);
  SaveError(err, status);
  if (*err != nullptr) {
    return;
  }

  for (int i = 0; i < num; i++) {
    iters[i] = new rocksdb_iterator_t;
    iters[i]->rep = res[i];
  }
}
