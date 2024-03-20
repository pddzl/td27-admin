package request

type CronReq struct {
	ID          uint        `json:"id" binding:"required"`
	Name        string      `json:"name" binding:"required"`
	Method      string      `json:"method" binding:"required"`
	Expression  string      `json:"expression" binding:"required"`
	Strategy    string      `json:"strategy" binding:"omitempty,oneof=always once"`
	Open        bool        `json:"open"`
	ExtraParams ExtraParams `json:"extraParams"`
	Comment     string      `json:"comment"`
}

type ExtraParams struct {
	TableInfo []ClearTable `json:"tableInfo"` // for clearTable
	Command   string       `json:"command"`   // for shell
}

type ClearTable struct {
	TableName    string `json:"tableName"`
	CompareField string `json:"compareField"`
	Interval     string `json:"interval"`
}

type SwitchReq struct {
	Id   uint `json:"id" binding:"required"`
	Open bool `json:"open"`
}
