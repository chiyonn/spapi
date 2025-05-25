package testutil

import (

	"encoding/json"
	"os"
	"testing"
	"path/filepath"
)

func LoadResponseJSON(t *testing.T, name string) string {
	t.Helper()

	path := filepath.Join("testdata", name)
	bytes, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read test data file %s: %v", path, err)
	}
	return string(bytes)
}


func LoadResponseStruct[T any](t *testing.T, name string) T {
	t.Helper()
	path := filepath.Join("testdata", name)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read test JSON: %v", err)
	}
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("failed to unmarshal test JSON: %v", err)
	}
	return result
}
