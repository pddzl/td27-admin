package sysTool

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"server/internal/global"
	"server/internal/pkg"
)

type FileRepository interface {
	Upload(context.Context, *multipart.FileHeader) (*FileModel, error)
	List(context.Context, ListFileReq) ([]*FileModel, int64, error)
	Delete(context.Context, string) error
}

type fileEntity struct {
	conn *gorm.DB
}

func NewFileEntity(conn *gorm.DB) FileRepository {
	return &fileEntity{conn: conn}
}

// Upload 上传文件
func (e *fileEntity) Upload(ctx context.Context, file *multipart.FileHeader) (*FileModel, error) {
	if file == nil {
		return nil, fmt.Errorf("file is nil")
	}

	// Prepare model
	uploadModel := &FileModel{
		Mime: file.Header.Get("Content-Domain"),
	}

	// Parse filename
	fileName, fileExt := pkg.GetFileAndExt(file.Filename)
	transformName := fmt.Sprintf("%s_%s%s", fileName, uuid.NewString(), fileExt)

	// Build path safely
	uploadDir := global.TD27_CONFIG.File.Upload
	fullPath := filepath.Join(uploadDir, transformName)

	uploadModel.FileName = transformName
	uploadModel.FullPath = fullPath

	// Ensure upload directory exists
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return nil, fmt.Errorf("create upload dir failed: %w", err)
	}

	// Open source file
	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("open upload file failed: %w", err)
	}
	defer src.Close()

	// Create destination file
	dst, err := os.Create(fullPath)
	if err != nil {
		return nil, fmt.Errorf("create dest file failed: %w", err)
	}
	defer dst.Close()

	// Copy content
	if _, err = io.Copy(dst, src); err != nil {
		return nil, fmt.Errorf("copy file failed: %w", err)
	}

	// Persist metadata
	if err = e.conn.WithContext(ctx).Create(uploadModel).Error; err != nil {
		return nil, fmt.Errorf("save file metadata failed: %w", err)
	}

	return uploadModel, nil
}

// List 分页获取文件信息
func (e *fileEntity) List(ctx context.Context, req ListFileReq) ([]*FileModel, int64, error) {

	req.Normalize()

	var files []*FileModel
	var total int64

	db := e.conn.WithContext(ctx).Model(&FileModel{})

	// Filters
	if req.Name != "" {
		db = db.Where("file_name LIKE ?", "%"+req.Name+"%")
	}

	// Count AFTER filters
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Fixed order field, direction controlled by Desc
	order := "id"
	if req.Desc {
		order += " desc"
	}

	// Pagination + query
	if err := db.
		Order(order).
		Limit(req.PageSize).
		Offset(req.Offset()).
		Find(&files).Error; err != nil {
		return nil, 0, err
	}

	return files, total, nil
}

// Delete 删除文件
func (e *fileEntity) Delete(ctx context.Context, fileName string) error {
	if fileName == "" {
		return fmt.Errorf("fileName is empty")
	}

	uploadDir := global.TD27_CONFIG.File.Upload
	fullPath := filepath.Join(uploadDir, fileName)

	// Remove physical file
	if err := os.Remove(fullPath); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("delete file failed: %w", err)
		}
		// File not exist is acceptable, continue to clean DB
	}

	// Remove DB record
	if err := e.conn.WithContext(ctx).
		Where("file_name = ?", fileName).
		Delete(&FileModel{}).Error; err != nil {
		return fmt.Errorf("delete file record failed: %w", err)
	}

	return nil
}
