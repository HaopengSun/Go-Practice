package tests

import (
	filesystem "go-practice/file_system"
	"os"
	"path/filepath"
	"testing"
)

// TestWriteFile tests the WriteFile function
func TestWriteFile(t *testing.T) {
	testDir := "./test_dir"
	testFile := filepath.Join(testDir, "testfile.txt")
	content := "Hello, test file!"

	// Ensure test directory exists
	if err := os.MkdirAll(testDir, 0755); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	// Write file
	filesystem.WriteFile(testFile, content)

	// Read the file back
	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	// Check content
	if string(data) != content {
		t.Errorf("Expected '%s', but got '%s'", content, string(data))
	}

	// Cleanup
	os.Remove(testFile)
}

// TestListFiles tests the ListFiles function
func TestListFiles(t *testing.T) {
	testDir := "./test_dir"

	// Ensure test directory exists
	if err := os.MkdirAll(testDir, 0755); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	// Create test files
	testFile1 := filepath.Join(testDir, "file1.txt")
	testFile2 := filepath.Join(testDir, "file2.txt")
	os.WriteFile(testFile1, []byte("file1"), 0644)
	os.WriteFile(testFile2, []byte("file2"), 0644)

	// Run ListFiles (manually verify by checking output)
	filesystem.ListFiles(testDir)

	// Cleanup
	os.Remove(testFile1)
	os.Remove(testFile2)
}

// TestCopyFile tests the CopyFile function
func TestCopyFile(t *testing.T) {
	testDir := "./test_dir"
	srcFile := filepath.Join(testDir, "source.txt")
	dstFile := filepath.Join(testDir, "copy.txt")
	content := "File copy test!"

	// Ensure test directory exists
	if err := os.MkdirAll(testDir, 0755); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	// Write the source file
	filesystem.WriteFile(srcFile, content)

	// Copy the file
	filesystem.CopyFile(srcFile, dstFile)

	// Read copied file
	data, err := os.ReadFile(dstFile)
	if err != nil {
		t.Fatalf("Failed to read copied file: %v", err)
	}

	// Check content
	if string(data) != content {
		t.Errorf("Expected copied content '%s', but got '%s'", content, string(data))
	}

	// Cleanup
	os.Remove(srcFile)
	os.Remove(dstFile)
}

// TestDeleteFile tests the DeleteFile function
func TestDeleteFile(t *testing.T) {
	testDir := "./test_dir"
	testFile := filepath.Join(testDir, "delete_me.txt")

	// Ensure test directory exists
	if err := os.MkdirAll(testDir, 0755); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	// Create file
	filesystem.WriteFile(testFile, "To be deleted")

	// Delete file
	filesystem.DeleteFile(testFile)

	// Verify deletion
	if _, err := os.Stat(testFile); !os.IsNotExist(err) {
		t.Errorf("File %s was not deleted", testFile)
	}
}
