package common

//Struct API
// Order struct (Model) ...

type Message struct {
	Code    int     `json:"code"`
	Remark  string  `json:"remark"`
	OrderID string  `json:"orderID"`
	Orders  *Orders `json:"orders,omitempty"`
	Result  *Result `json:"result,omitempty"`
}

type Orders struct {
	OrderID    string         `json:"orderID"`
	CustomerID string         `json:"customerID"`
	EmployeeID string         `json:"employeeID"`
	OrderDate  string         `json:"orderDate"`
	OrdersDet  []OrdersDetail `json:"ordersDetail"`
}

type OrdersDetail struct {
	OrderID     string  `json:"orderID"`
	ProductID   string  `json:"ProductID"`
	ProductName string  `json:"ProductName"`
	UnitPrice   float64 `json:"UnitPrice"`
	Quantity    int     `json:"Quantity"`
}

type Result struct {
	Code   int    `json:"code"`
	Remark string `json:"remark,omitempty"`
}
type Customers struct {
	CustomerID   string `json:"CustomerID"`
	CompanyName  string `json:"CompanyName"`
	ContactName  string `json:"ContactName"`
	ContactTitle string `json:"ContactTitle"`
	Address      string `json:"Address"`
	City         string `json:"City"`
	Country      string `json:"Country"`
	Phone        string `json:"Phone"`
	PostalCode   string `json:"PostalCode"`
}
type Products struct {
	ProductID       string `json:"ProductID"`
	ProductName     string `json:"ProductName"`
	SupplierID      string `json:"SupplierID"`
	CategoryID      string `json:"CategoryID"`
	QuantityPerUnit string `json:"QuantityPerUnit"`
	UnitPrice       string `json:"UnitPrice"`
	UnitsInStock    string `json:"UnitsInStock"`
}

//End Struct API
type FastPayRequest struct {
	TransId    string `json:"TransId"`
	Merchant   string `json:"merchant"`
	MerchantID string `json:"merchant_id"`
	Request    string `json:"request"`
	Signature  string `json:"signature"`
}

type FastPayResponse struct {
	TransId       string `json:"TransId"`
	Merchant      string `json:"Merchant"`
	Merchant_id   string `json:"Merchant_id"`
	Code          string `json:"Code"`
	Name          string `json:"Name"`
	Response      string `json:"Response"`
	Response_Code string `json:"Response_Code"`
	Response_Desc string `json:"Response_Desc"`
}

//mytrip

type MyTripsrequest struct {
	DepatureDate1 string `json:"depature_date_1"`
	DepatureDate2 string `json:"depature_date_2"`
	Provinsi      int64  `json:"provinsi"`
}

type MytripsResponse struct {
	Message    string       `json:"message"`
	Status     string       `json:"status"`
	TripDetail []TripDetail `json:"data"`
}

type TripDetail struct {
	AirlineName      string `json:"AirlineName,omitempty"`
	AirportName      string `json:"AirportName,omitempty"`
	CityName         string `json:"CityName,omitempty"`
	Currency         string `json:"Currency,omitempty"`
	DepartureDate    string `json:"DepartureDate,omitempty"`
	Description      string `json:"Description,omitempty"`
	Destination      string `json:"Destination,omitempty"`
	DetailTransit    string `json:"DetailTransit,omitempty"`
	DoubleType       string `json:"DoubleType,omitempty"`
	Duration         string `json:"Duration,omitempty"`
	Goods            string `json:"Goods,omitempty"`
	HotelName        string `json:"HotelName,omitempty"`
	HotelRating      string `json:"HotelRating,omitempty"`
	Lat              string `json:"Lat,omitempty"`
	LicenseNumber    string `json:"LicenseNumber,omitempty"`
	Logo             string `json:"Logo,omitempty"`
	Long             string `json:"Long,omitempty"`
	Origin           string `json:"Origin,omitempty"`
	OriginCity       string `json:"OriginCity,omitempty"`
	Price            string `json:"Price,omitempty"`
	PromoCode        string `json:"PromoCode,omitempty"`
	PromoDescription string `json:"PromoDescription,omitempty"`
	Provinsi         string `json:"Provinsi,omitempty"`
	QuadType         string `json:"QuadType,omitempty"`
	Rating           string `json:"Rating,omitempty"`
	ReturnDate       string `json:"ReturnDate,omitempty"`
	TermCondition    string `json:"TermCondition,omitempty"`
	Transit          string `json:"Transit,omitempty"`
	TravelID         string `json:"TravelID,omitempty"`
	TravelName       string `json:"TravelName,omitempty"`
	TripID           string `json:"TripID,omitempty"`
	TripleType       string `json:"TripleType,omitempty"`
}
