package domain

type Dentist struct {
	Id       int    `json:"id"`
	Name     string `json:"name" `
	LastName string `json:"last_name" `
	License  string `json:"license" `
}
