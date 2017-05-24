package main

// UUID is a 128-bit universally unique identifier.
type UUID [16]int8

// Pipeline is a logical grouping of processing steps, and connects the steps
// to product owner(s).
type Pipeline struct {
	id       UUID
	owner    UUID
	requires []UUID
	produces []UUID
	title    string
}

// Step is a node in the pipeline. A step can have parents and children, and
// fans in and out.
type Step struct {
	id          UUID              // Unique identifier for this step.
	pipeline    *Pipeline         // Which pipeline the step is in.
	conflicts   []*Step           // Any steps that conflicts with this step are automatically excluded.
	parents     []*Step           // Parent steps for predecession or fan-in.
	children    []*Step           // Child steps for succession or fan-out.
	concurrency int               // How many instances of this step that can be run at the same time.
	command     string            // The command to run on the executor.
	envVars     map[string]string // Environment variables to pass to the executor.
	numJobs     int               // Number of jobs to generate from this step.
	retryPolicy struct {
		numTries        int     // Number of times to try running this step.
		backoffFactor   float64 // Multiply by this number each time a job fails.
		backoffInterval float64 // The time to sleep between job retries, in seconds.
	}
}

// Job is an instance of a pipeline step. When the first step in a pipeline is
// activated, all jobs for the pipeline are created, according to how much data
// is expected.
type Job struct {
	id          UUID              // Unique identifier for this job.
	pipeline    *Pipeline         // Which pipeline this job originates from.
	step        *Step             // Which step this job originates from.
	status      int               // Status code of the job, defined in the Status* constants.
	stdout      string            // Unix STDOUT output from the job.
	stderr      string            // Unix STDERR output from the job.
	successor   *Job              // If this job failed, successor is a reference to the next retry.
	predecessor *Job              // If this job previously failed, predecessor is a reference to the previous try.
	failures    int               // Number of times this series of jobs has failed.
	command     string            // Shell command to run.
	envVars     map[string]string // Environment variables to pass to the executor.
}

// Stakeholder is a human or group who depend on data.
type Stakeholder struct {
	id    UUID   // Unique identifier for this stakeholder.
	name  string // Actor name.
	email string // Actor e-mail address.
}

// Job status codes.
const (
	NEW = iota // Job is created, but hasn't been scheduled.
)

func main() {
}
