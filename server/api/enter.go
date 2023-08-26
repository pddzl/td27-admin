package api

import (
	"server/api/fileM"
	"server/api/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
	FileApiGroup   fileM.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
