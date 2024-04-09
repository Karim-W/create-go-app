package rsync

import (
	"os"
	"testing"
)

func TestRsync(t *testing.T) {
	rdb := os.Getenv("REDIS_URI")
	if rdb == "" {
		t.Skip("skipping test; $REDIS_URI not set")
	}

	rs := Init(rdb)

	mtx := rs.NewMutex("test")
	mtx.Lock()
	mtx.Unlock()
}

func TestRsyncThreeLayerLock(t *testing.T) {
	rdb := os.Getenv("REDIS_URI")
	if rdb == "" {
		t.Skip("skipping test; $REDIS_URI not set")
	}

	rs := Init(rdb)

	mtx1 := rs.NewMutex("test1")
	mtx2 := rs.NewMutex("test2")

	// mutex 1
	mtx1.Lock()
	mtx1.Unlock()

	// mutex 2
	mtx2.Lock()
	mtx2.Unlock()

	// mutex 1
	renewedMtx1 := rs.NewMutex("test1")
	renewedMtx1.Lock()
	renewedMtx1.Unlock()
}
