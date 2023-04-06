package domain

type Patient struct {
	Id            int    `json:"id"`
	Name          string `json:"name" `
	LastName      string `json:"last_name" `
	Domicilio     string `json:"domicilio" `
	Dni           int    `json:"dni" `
	Email         string `json:"email" `
	AdmissionDate string `json:"admission_date" `
}
