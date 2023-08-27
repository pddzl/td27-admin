package fileM

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"server/global"
	fileMModel "server/model/fileM"
	fileMReq "server/model/fileM/request"
	"server/utils"
)

type FileService struct{}

// Upload 上传文件
func (fs *FileService) Upload(file *multipart.FileHeader) (string, error) {
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

// GetFileList 分页获取文件信息
func (fs *FileService) GetFileList(params fileMReq.FileSearchParams) ([]fileMModel.FileModel, int64, error) {
	limit := params.PageSize
	offset := params.PageSize * (params.Page - 1)
	db := global.TD27_DB.Model(&fileMModel.FileModel{})
	var fileList []fileMModel.FileModel

	if params.Name != "" {
		db = db.Where("file_name LIKE ?", "%"+params.Name+"%")
	}

	var total int64
	err := db.Count(&total).Error

	if err != nil {
		return fileList, total, err
	} else {
		db = db.Limit(limit).Offset(offset)
		if params.OrderKey != "" {
			var orderStr string
			// 设置有效排序key 防止sql注入
			orderMap := make(map[string]bool, 1)
			orderMap["ID"] = true
			if orderMap[params.OrderKey] {
				if params.Desc {
					orderStr = params.OrderKey + " desc"
				} else {
					orderStr = params.OrderKey
				}
			} else { // didn't match any order key in `orderMap`
				err = fmt.Errorf("非法的排序字段: %v", params.OrderKey)
				return fileList, total, err
			}

			err = db.Order(orderStr).Find(&fileList).Error
		} else {
			err = db.Find(&fileList).Error
		}
	}
	return fileList, total, err
}
