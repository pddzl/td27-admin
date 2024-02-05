package request

type CronReq struct {
	ID          uint        `json:"id" validate:"required"`
	Name        string      `json:"name"`
	Method      string      `json:"method"`
	Expression  string      `json:"expression"`
	Strategy    string      `json:"strategy" validate:"omitempty,oneof=always once"`
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
	Id   uint `json:"id" validate:"required"`
	Open bool `json:"open"`
}
