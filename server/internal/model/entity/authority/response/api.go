package response

type ApiTree struct {
	List       interface{} `json:"list"`
	CheckedKey []string    `json:"checkedKey"`
}
