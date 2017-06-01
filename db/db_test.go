package db_test

import (
	"testing"

	"github.com/metno/muppet/db"
	"github.com/metno/muppet/job"
)

// Test that the same amount of input jobs are output back on the job queue.
func TestEcho(t *testing.T) {
	max := db.MAX_QUEUE
	db := db.New()
	go db.Run()
	for i := 0; i < max; i++ {
		j := job.New()
		db.In <- j
	}
	for i := 0; i < max; i++ {
		select {
		case <-db.Out:
		default:
			t.Fail()
		}
	}
}
