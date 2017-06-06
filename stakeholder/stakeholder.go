// Package stakeholder provides a data model for stakeholders.
package stakeholder

import uuid "github.com/satori/go.uuid"

// Stakeholder is a human or group who depend on data.
type Stakeholder struct {
	id    uuid.UUID // Unique identifier for this stakeholder.
	name  string    // Actor name.
	email string    // Actor e-mail address.
}
