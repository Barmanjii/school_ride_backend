package model

type Address struct {
	Flat     string `json:"flat"`
	Street   string `json:"street"`
	Landmark string `json:"landmark"`
	City     string `json:"city"`
	State    string `json:"state"`
	Pincode  string `json:"pincode"`
	Country  string `json:"country"`
}
