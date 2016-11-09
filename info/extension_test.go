package info

import (
	"testing"

	"github.com/aj0strow/pgschema/temp"
)

func TestLoadExtensions(t *testing.T) {
	conn, err := temp.Connect("pgschema")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	exts, err := LoadExtensions(conn)
	found := false
	for _, ext := range exts {
		if ext.ExtName == "plpgsql" {
			found = true
		}
	}
	if !found {
		t.Fatal("missing extension plpgsql")
	}
}
