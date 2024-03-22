package request

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
	ID   uint `json:"id" binding:"required"`
	Open bool `json:"open"`
}
