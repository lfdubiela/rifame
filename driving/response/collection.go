package response

type Collection struct {
	List   interface{} `json:"list"`
	Length int         `json:"length"`
}
