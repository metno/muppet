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
	id          uuid.UUID          // Unique identifier for this job.
	pipeline    *pipeline.Pipeline // Which pipeline this job originates from.
	step        *step.Step         // Which step this job originates from.
	status      int                // Status code of the job, defined in the Status* constants.
	stdout      string             // Unix STDOUT output from the job.
	stderr      string             // Unix STDERR output from the job.
	successor   *Job               // If this job failed, successor is a reference to the next retry.
	predecessor *Job               // If this job previously failed, predecessor is a reference to the previous try.
	failures    int                // Number of times this series of jobs has failed.
	command     string             // Shell command to run.
	envVars     map[string]string  // Environment variables to pass to the executor.
}

// Job status codes.
const (
	NEW = iota // Job is created, but hasn't been scheduled.
)

// New returns a new Job object.
func New() Job {
	return Job{
		id:      uuid.NewV4(),
		envVars: make(map[string]string, 0),
	}
}
