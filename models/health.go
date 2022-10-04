package models

// swagger:model Adherent
type Health struct {
	// Title of the api page
	// in: string
	Title string `json:"title"`
	// Version of the api
	// in: string
	Version string `json:"version"`
}

func GetHealth() *Health {
	var hea Health

	hea.Title = "IT works !"
	hea.Version = "0.0.1"

	return &hea
}
