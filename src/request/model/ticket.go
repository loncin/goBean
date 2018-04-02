package model

// 退改规则
type Ticket struct {
	ReservationType int    `json:"reservationType"`
	IsChangeCode    int    `json:"isChangeCode"`
	FareType        int    `json:"fareType"`
	IsAutoTicket    int    `json:"isAutoTicket"`
	TicketAgencyFee string `json:"ticketAgencyFee"`
	SpecialOrder    string `json:"specialOrder"`
	CabinChange     string `json:"cabinChange"`
	TicketOffice    string `json:"ticketOffice"`
}
