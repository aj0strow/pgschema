package next

import (
	"github.com/aj0strow/pgschema/ab"
	"github.com/aj0strow/pgschema/db"
)

type CreateExtension struct {
	*db.Extension
}

func createExtensions(exts []ab.ExtensionMatch) []CreateExtension {
	var xs []CreateExtension
	for _, ext := range exts {
		if ext.B == nil {
			xs = append(xs, CreateExtension{ext.A})
		}
	}
	return xs
}

type DropExtension struct {
	*db.Extension
}

func dropExtensions(exts []ab.ExtensionMatch) []DropExtension {
	var xs []DropExtension
	var keep = map[string]bool{
		"plpgsql": true,
	}
	for _, ext := range exts {
		if ext.A == nil && !keep[ext.B.ExtName] {
			xs = append(xs, DropExtension{ext.B})
		}
	}
	return xs
}
