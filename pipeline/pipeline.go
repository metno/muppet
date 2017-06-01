// Package pipeline provides a model for pipelines.
package pipeline

import "github.com/metno/muppet/types"

// Pipeline is a logical grouping of processing steps, and connects the steps
// to product owner(s).
type Pipeline struct {
	id       types.UUID
	owner    types.UUID
	requires []types.UUID
	produces []types.UUID
	title    string
}
