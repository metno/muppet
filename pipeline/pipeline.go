// Package pipeline provides a model for pipelines.
package pipeline

import uuid "github.com/satori/go.uuid"

// Pipeline is a logical grouping of processing steps, and connects the steps
// to product owner(s).
type Pipeline struct {
	id       uuid.UUID
	owner    uuid.UUID
	requires []uuid.UUID
	produces []uuid.UUID
	title    string
}
