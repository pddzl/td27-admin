package fileM

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"server/global"
	fileMModel "server/model/fileM"
	"server/utils"
)

type FileService struct{}

// Upload 上传文件
func (u *FileService) Upload(file *multipart.FileHeader) (string, error) {
	var uploadModel fileMModel.FileModel
	uploadModel.Mime = file.Header.Get("Content-Type")
	// 读取文件、文件后缀
	fileName, fileExt := utils.GetFileAndExt(file.Filename)
	// 转换后的文件名
	transformName := fmt.Sprintf("%s_%s%s", fileName, uuid.New(), fileExt)
	fullPath := fmt.Sprintf("%s/%s", global.TD27_CONFIG.System.Upload, transformName)
	uploadModel.FileName = transformName
	uploadModel.FullPath = fullPath

	// 读取文件内容
	f, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("open file failed: %s", err.Error())
	}
	defer f.Close()

	// 创建目标文件
	destFile, err := os.Create(fullPath)
	if err != nil {
		return "", fmt.Errorf("create dest file failed: %s", err.Error())
	}
	defer destFile.Close()

	// copy内容到目标文件
	_, err = io.Copy(destFile, f)
	if err != nil {
		return "", fmt.Errorf("copy file failed: %s", err.Error())
	}

	global.TD27_DB.Save(&uploadModel)

	return fullPath, nil
}
