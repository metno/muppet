// Package step provides a model for logical steps in the pipeline.
package step

import (
	"github.com/metno/muppet/pipeline"
	uuid "github.com/satori/go.uuid"
)

// Step is a node in the pipeline. A step can have parents and children, and
// fans in and out.
type Step struct {
	id          uuid.UUID          // Unique identifier for this step.
	pipeline    *pipeline.Pipeline // Which pipeline the step is in.
	conflicts   []*Step            // Any steps that conflicts with this step are automatically excluded.
	parents     []*Step            // Parent steps for predecession or fan-in.
	children    []*Step            // Child steps for succession or fan-out.
	concurrency int                // How many instances of this step that can be run at the same time.
	command     string             // The command to run on the executor.
	envVars     map[string]string  // Environment variables to pass to the executor.
	numJobs     int                // Number of jobs to generate from this step.
	retryPolicy struct {
		numTries        int     // Number of times to try running this step.
		backoffFactor   float64 // Multiply by this number each time a job fails.
		backoffInterval float64 // The time to sleep between job retries, in seconds.
	}
}
