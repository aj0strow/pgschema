package plan

import (
	"fmt"
)

// Change is a single atomic schema change.
type Change interface {
	String() string
}

// The search path needs to be updated depending on the active schema.
type SetSearchPath struct {
	SchemaName string
}

func (sp SetSearchPath) String() string {
	return fmt.Sprintf(`SET search_path TO %s`, sp.SchemaName)
}

var _ Change = (*SetSearchPath)(nil)

// Create a new schema. This change occurs when you have a schema
// but it doesn't exist in the database yet.
type CreateSchema struct {
	SchemaName string
}

func (cs CreateSchema) String() string {
	return fmt.Sprintf(`CREATE SCHEMA %s`, cs.SchemaName)
}

var _ Change = (*CreateSchema)(nil)
