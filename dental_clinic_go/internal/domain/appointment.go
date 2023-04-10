package domain

type Appointment struct {
	Id          int     `json:"id"`
	Date        string  `json:"date"`
	Hour        string  `json:"hour"`
	Description string  `json:"description"`
	Patient     Patient `json:"patient"`
	Dentist     Dentist `json:"dentist"`
}
