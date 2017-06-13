// Package scheduler provides data flow.
package scheduler

import (
	"github.com/metno/muppet/job"
	"github.com/metno/muppet/pipeline"
	"github.com/metno/muppet/step"
)

// Scheduler provides job data flow, and keeps internal state about job state.
type Scheduler struct {
	pipelines []pipeline.Pipeline
	steps     []step.Step
	jobs      []job.Job
}

// New returns Scheduler.
func New() *Scheduler {
	return &Scheduler{
		pipelines: make([]pipeline.Pipeline, 0),
		steps:     make([]step.Step, 0),
		jobs:      make([]job.Job, 0),
	}
}
