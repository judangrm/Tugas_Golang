package common

type Employee struct {
	EmployeeID string `json:"EmployeeID"`
	LastName   string `json:"LastName"`
	FirstName  string `json:"FirstName"`
	Title      string `json:"Title"`
	Address    string `json:"Address"`
	City       string `json:"City"`
	Country    string `json:"Country"`
	HomePhone  string `json:"HomePhome"`
}
