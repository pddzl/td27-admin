package sysTool

import (
	"context"
	"mime/multipart"

	"server/internal/global"
	"server/internal/model/sysTool"
)

type FileService struct {
	repository sysTool.FileRepository
	ctx        context.Context
}

func NewFileService() *FileService {
	return &FileService{
		repository: sysTool.NewFileEntity(global.TD27_DB),
		ctx:        context.Background(),
	}
}

// Upload 上传文件
func (s *FileService) Upload(file *multipart.FileHeader) (*sysTool.FileModel, error) {
	instance, err := s.repository.Upload(s.ctx, file)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

// List 分页获取文件信息
func (s *FileService) List(req sysTool.ListFileReq) ([]*sysTool.FileModel, int64, error) {
	list, count, err := s.repository.List(s.ctx, req)
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

// Delete 删除文件
func (s *FileService) Delete(fileName string) error {
	err := s.repository.Delete(s.ctx, fileName)
	if err != nil {
		return err
	}
	return nil
}
