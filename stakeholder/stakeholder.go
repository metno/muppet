// Package stakeholder provides a data model for stakeholders.
package stakeholder

import "github.com/metno/muppet/types"

// Stakeholder is a human or group who depend on data.
type Stakeholder struct {
	id    types.UUID // Unique identifier for this stakeholder.
	name  string     // Actor name.
	email string     // Actor e-mail address.
}
