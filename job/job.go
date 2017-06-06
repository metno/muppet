// Package job provides the data model for a job.
package job

import (
	"github.com/metno/muppet/pipeline"
	"github.com/metno/muppet/step"
	uuid "github.com/satori/go.uuid"
)

// Job is an instance of a pipeline step. When the first step in a pipeline is
// activated, all jobs for the pipeline are created, according to how much data
// is expected.
type Job struct {
	ID          uuid.UUID          // Unique identifier for this job.
	Pipeline    *pipeline.Pipeline // Which pipeline this job originates from.
	Step        *step.Step         // Which step this job originates from.
	Status      int                // Status code of the job, defined in the Status* constants.
	Stdout      string             // Unix STDOUT output from the job.
	Stderr      string             // Unix STDERR output from the job.
	Successor   *Job               // If this job failed, successor is a reference to the next retry.
	Predecessor *Job               // If this job previously failed, predecessor is a reference to the previous try.
	Failures    int                // Number of times this series of jobs has failed.
	Command     string             // Shell command to run.
	EnvVars     map[string]string  // Environment variables to pass to the executor.
}

// Job status codes.
const (
	NEW = iota // Job is created, but hasn't been scheduled.
)

// New returns a new Job object.
func New() Job {
	return Job{
		ID:      uuid.NewV4(),
		EnvVars: make(map[string]string, 0),
	}
}
