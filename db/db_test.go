package db_test

import (
	"testing"
	"time"

	"github.com/metno/muppet/db"
	"github.com/metno/muppet/job"
	"github.com/stretchr/testify/assert"
)

// Test that input jobs are echoed back on the job queue after persisting.
func TestEcho(t *testing.T) {

	// Generate MAX_QUEUE jobs.
	jobs := make([]job.Job, db.MAX_QUEUE)
	for i := 0; i < db.MAX_QUEUE; i++ {
		jobs[i] = job.New()
	}

	// Submit all jobs asynchronously to the database thread.
	dbThread := db.New()
	go dbThread.Run()
	go func() {
		for i := 0; i < db.MAX_QUEUE; i++ {
			dbThread.In <- jobs[i]
		}
	}()

	// Receive all messages from the DB channel, but wait at most 1 second.
	// Assert equality on all job objects.
	timeout := time.After(1 * time.Second)
	for i := 0; i < db.MAX_QUEUE; i++ {
		select {
		case j := <-dbThread.Out:
			assert.Equal(t, jobs[i], j)
		case <-timeout:
			t.Logf("Not enough data available on output queue.")
			t.Fail()
		}
	}
}
