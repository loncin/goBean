package model

// 乘机人
type Passenger struct {
	Name           string `json:"name"`
	ProductType    string `json:"productType"`
	Surname        string `json:"surname"`
	GivenName      string `json:"givenName"`
	PassengerType  string `json:"passengerType"`
	Birthday       string `json:"birthday"`
	Gender         string `json:"gender"`
	CardNum        string `json:"cardNum"`
	CardType       string `json:"cardType"`
	CardIssuePlace string `json:"cardIssuePlace"`
	CardExpired    string `json:"cardExpired"`
	Nationality    string `json:"nationality"`
	Mobile         string `json:"mobile"`
}
