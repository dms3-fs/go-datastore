package sync

import (
	"testing"

	ds "github.com/dms3-fs/go-datastore"
	dstest "github.com/dms3-fs/go-datastore/test"
)

func TestSync(t *testing.T) {
	dstest.SubtestAll(t, MutexWrap(ds.NewMapDatastore()))
}
