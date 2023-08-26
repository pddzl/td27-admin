package utils

import "path/filepath"

// GetFileAndExt 获取文件名、文件后缀
// "/path/to/your/file.txt" -> file, txt
func GetFileAndExt(filePath string) (string, string) {
	fileName := filepath.Base(filePath) // 获取完整文件名，包括后缀
	ext := filepath.Ext(fileName)
	return fileName[:len(fileName)-len(ext)], ext
}
