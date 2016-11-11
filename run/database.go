package run

import (
	"github.com/aj0strow/pgschema/ab"
)

type UpdateDatabase struct {
	CreateExtensions []CreateExtension
	DropExtensions   []DropExtension
}

func updateDatabase(db ab.DatabaseMatch) UpdateDatabase {
	return UpdateDatabase{
		CreateExtensions: createExtensions(db.ExtensionMatches),
		DropExtensions:   dropExtensions(db.ExtensionMatches),
	}
}
