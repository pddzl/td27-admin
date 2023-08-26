package fileM

import "server/service"

type ApiGroup struct {
	FileApi
}

var (
	fileService = service.ServiceGroupApp.FileMServiceGroup.FileService
)
