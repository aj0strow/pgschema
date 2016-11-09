package psql

import (
	"testing"

	"github.com/aj0strow/pgschema/temp"
)

func TestLoadExtensionNodes(t *testing.T) {
	conn, err := temp.Connect("pgschema")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	extNodes, err := LoadExtensionNodes(conn)
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, node := range extNodes {
		if node.Extension.ExtName == "plpgsql" {
			found = true
		}
	}
	if !found {
		t.Fatal("missing extension node plpgsql")
	}
}

func TestLoadExtensions(t *testing.T) {
	conn, err := temp.Connect("pgschema")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	exts, err := LoadExtensions(conn)
	if err != nil {
		t.Fatal(err)
	}
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
