package pkg

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPathExists(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()

	// Test case: directory exists
	t.Run("directory exists", func(t *testing.T) {
		exists, err := PathExists(tmpDir)
		if err != nil {
			t.Errorf("PathExists(%q) unexpected error: %v", tmpDir, err)
		}
		if !exists {
			t.Errorf("PathExists(%q) = %v, want true", tmpDir, exists)
		}
	})

	// Test case: file exists (not directory)
	t.Run("file exists", func(t *testing.T) {
		tmpFile := filepath.Join(tmpDir, "testfile.txt")
		if err := os.WriteFile(tmpFile, []byte("test"), 0644); err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}

		exists, err := PathExists(tmpFile)
		if err == nil {
			t.Errorf("PathExists(%q) expected error for file, got nil", tmpFile)
		}
		if exists {
			t.Errorf("PathExists(%q) = %v, want false for file", tmpFile, exists)
		}
	})

	// Test case: path does not exist
	t.Run("path not exists", func(t *testing.T) {
		nonExistent := filepath.Join(tmpDir, "nonexistent")
		exists, err := PathExists(nonExistent)
		if err != nil {
			t.Errorf("PathExists(%q) unexpected error: %v", nonExistent, err)
		}
		if exists {
			t.Errorf("PathExists(%q) = %v, want false", nonExistent, exists)
		}
	})
}
