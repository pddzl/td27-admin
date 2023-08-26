package fileM

import "server/global"

type FileModel struct {
	global.TD27_MODEL
	FileName string `json:"fileName" gorm:"comment:文件名"`
	FullPath string `json:"fullPath" gorm:"comment:文件完整路径"`
	Mime     string `json:"mime" gorm:"comment:文件类型"`
}

func (FileModel) TableName() string {
	return "fileM_file"
}
