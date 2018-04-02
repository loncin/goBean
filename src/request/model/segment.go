package model

// 航段
type Segment struct {
	Carrier               string   `json:"carrier"`
	DepAirport            string   `json:"depAirport"`
	DepTime               string   `json:"depTime"`
	ArrAirport            string   `json:"arrAirport"`
	ArrTime               string   `json:"arrTime"`
	CodeShare             bool     `json:"codeShare"`
	OperatingCarrier      string   `json:"operatingCarrier"`
	OperatingFlightNumber string   `json:"operatingFlightNumber"`
	Cabin                 string   `json:"cabin"`
	AircraftCode          string   `json:"aircraftCode"`
	FlightNumber          string   `json:"flightNumber"`
	DepTerminal           string   `json:"depTerminal"`
	ArrTerminal           string   `json:"arrTerminal"`
	CabinGrade            string   `json:"cabinGrade"`
	Duration              int      `json:"duration"`
	StopTime              int      `json:"stopTime"`
	Stops                 []string `json:"stops"`
	SegmentNo             int      `json:"segmentNo"`
	FlightType            string   `json:"flightType"`
}
