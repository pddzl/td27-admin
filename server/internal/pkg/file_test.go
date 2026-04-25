package pkg

import (
	"testing"
)

func TestGetFileAndExt(t *testing.T) {
	tests := []struct {
		name         string
		filePath     string
		expectName   string
		expectExt    string
	}{
		{
			name:       "normal path",
			filePath:   "/path/to/your/file.txt",
			expectName: "file",
			expectExt:  ".txt",
		},
		{
			name:       "no directory",
			filePath:   "file.txt",
			expectName: "file",
			expectExt:  ".txt",
		},
		{
			name:       "multiple dots",
			filePath:   "/path/to/file.tar.gz",
			expectName: "file.tar",
			expectExt:  ".gz",
		},
		{
			name:       "no extension",
			filePath:   "/path/to/file",
			expectName: "file",
			expectExt:  "",
		},
		{
			name:       "hidden file",
			filePath:   "/path/to/.hidden",
			expectName: "",
			expectExt:  ".hidden",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name, ext := GetFileAndExt(tt.filePath)
			if name != tt.expectName {
				t.Errorf("GetFileAndExt(%q) name = %q, want %q", tt.filePath, name, tt.expectName)
			}
			if ext != tt.expectExt {
				t.Errorf("GetFileAndExt(%q) ext = %q, want %q", tt.filePath, ext, tt.expectExt)
			}
		})
	}
}
