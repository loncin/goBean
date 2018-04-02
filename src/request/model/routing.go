package model

type Routing struct {
	Data              string    `json:"data"`
	ReservationType   string    `json:"reservationType"`
	AdultPrice        int       `json:"adultPrice"`
	AdultTax          int       `json:"adultTax"`
	ChildPrice        int       `json:"childPrice"`
	ChildTax          int       `json:"childTax"`
	ValidatingCarrier string    `json:"validatingCarrier"`
	InvoiceType       int       `json:"invoiceType"`
	IsLowCabin        string    `json:"isLowCabin"`
	PolicyType        string    `json:"policyType"`
	TicketParamInfo   Ticket    `json:"ticketParamInfo"`
	Rule              Rule      `json:"rule"`
	FromSegments      []Segment `json:"fromSegments"`
	RetSegments       []Segment `json:"retSegments"`
	Segments          []Segment `json:"segments"`
}
