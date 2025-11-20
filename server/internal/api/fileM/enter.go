package fileM

import (
	"server/internal/service"
)

type ApiGroup struct {
	FileApi
}

var (
	fileService = service.ServiceGroupApp.FileM.FileService
)
