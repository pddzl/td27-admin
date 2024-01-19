package response

type Menu struct {
	List    interface{} `json:"list"`
	MenuIds []uint      `json:"menuIds"`
}
