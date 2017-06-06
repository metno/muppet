// Package db provides persistence for scheduled and running jobs.
package db

import "github.com/metno/muppet/job"

// MAX_QUEUE is the maximum number of jobs that can be queued in the I/O buffers.
const MAX_QUEUE int = 4096

// Database implements persistence for jobs. The channel `In` receives job data
// that should be persisted to the database, while the `Out` channel sends
// objects that have already been persisted.
type Database struct {
	In  chan job.Job
	Out chan job.Job
}

// New returns Database.
func New() *Database {
	return &Database{
		In:  make(chan job.Job, MAX_QUEUE),
		Out: make(chan job.Job, MAX_QUEUE),
	}
}

// Run provides the main loop for the database process. It reads jobs from the
// input channel, persists them to the database, then sends a copy of the
// persisted object on the output channel.
func (db *Database) Run() {
	for {
		j := <-db.In
		db.persist(j)
		db.Out <- j
	}
}

// persist stores an object in the database. This function must never fail, so
// if persistence fails, it must be retried until success.
// FIXME: jobs are not persisted, figure out a database backend to use.
func (db *Database) persist(j job.Job) {
}
