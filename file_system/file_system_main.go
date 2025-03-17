package filesystem

import (
	"fmt"
	"io"
	"os"
)

// ListFiles lists all files in the specified directory
func ListFiles(dirPath string) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	fmt.Println("Files in directory:", dirPath)
	for _, file := range files {
		fmt.Println(" -", file.Name())
	}
}

// ReadFile reads and prints the contents of a file
func ReadFile(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File contents:\n", string(data))
}

// WriteFile creates a new file and writes data to it
func WriteFile(filePath, content string) {
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("File written successfully:", filePath)
}

// DeleteFile removes a file
func DeleteFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
	fmt.Println("File deleted successfully:", filePath)
}

// CopyFile copies a file from source to destination
func CopyFile(src, dst string) {
	sourceFile, err := os.Open(src)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}

	fmt.Println("File copied successfully:", src, "→", dst)
}
