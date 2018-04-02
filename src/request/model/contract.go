package model

// 联系人
type Contract struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Postcode string `json:"postcode"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
}
