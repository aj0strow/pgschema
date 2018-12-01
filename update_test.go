package main

import (
	"bytes"
	"github.com/aj0strow/pgschema/order"
	"github.com/aj0strow/pgschema/source/hcl"
	"github.com/aj0strow/pgschema/source/psql"
	"github.com/aj0strow/pgschema/temp"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExamples(t *testing.T) {
	fds, err := ioutil.ReadDir("testdata")
	if err != nil {
		t.Fatal(err)
	}
	var paths []string
	for _, fd := range fds {
		paths = append(paths, filepath.Join("testdata", fd.Name()))
	}
	for _, path := range paths {
		RunExample(t, path)
	}
}

func RunExample(t *testing.T, path string) {
	// New ephemeral schema using the `pgschema/temp` package.
	conn, err := temp.Connect("pgschema")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	// Setup SQL database.
	f1, err := ioutil.ReadFile(filepath.Join(path, "setup.sql"))
	if err != nil {
		t.Fatal(err)
	}
	setupSQLFile := bytes.Replace(f1, []byte("_schema_"), []byte(conn.SchemaName), 1)
	if _, err := conn.Exec(string(setupSQLFile)); err != nil {
		t.Fatal(err)
	}

	// Load HCL schema from file.
	f2, err := ioutil.ReadFile(filepath.Join(path, "schema.hcl"))
	if err != nil {
		t.Fatal(err)
	}
	hclFile := bytes.Replace(f2, []byte("_schema_"), []byte(conn.SchemaName), 1)
	a, err := hcl.ParseBytes(hclFile)
	if err != nil {
		t.Fatal(err)
	}

	// Load PSQL schema from database.
	b, err := psql.LoadDatabase(conn)
	if err != nil {
		t.Fatal(err)
	}

	// Compare expected changes.
	wantChanges := []string{}
	if _, err := os.Stat(filepath.Join(path, "changes.sql")); !os.IsNotExist(err) {
		f3, err := ioutil.ReadFile(filepath.Join(path, "changes.sql"))
		if err != nil {
			t.Fatal(err)
		}
		changesSQLFile := bytes.Replace(f3, []byte("_schema_"), []byte(conn.SchemaName), -1)
		wantChanges = parseChangesSQLFile(changesSQLFile)
	}
	haveChanges := order.Changes(a, b)
	for i := range haveChanges {
		if i >= len(wantChanges) {
			t.Errorf("Unexpected change:\nhave: %s\n", haveChanges[i].String())
		} else if haveChanges[i].String() != wantChanges[i] {
			t.Errorf("Unexpected change:\nhave: %s\nwant: %s\n", haveChanges[i], wantChanges[i])
		}
	}
	for i := range wantChanges {
		if i >= len(haveChanges) {
			t.Errorf("Exected change:\nwant: %s\n", wantChanges[i])
		}
	}
}

func parseChangesSQLFile(contents []byte) []string {
	changes := strings.Split(string(contents), ";\n")
	// Remove empty string at end of slice.
	if len(changes) > 0 && changes[len(changes)-1] == "" {
		changes = changes[0 : len(changes)-1]
	}
	return changes
}
