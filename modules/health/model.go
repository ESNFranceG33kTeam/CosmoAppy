package health

// swagger:model Health
type Health struct {
	// Title of the api page
	// in: string
	Title string `json:"title"`
	// Version of the api
	// in: string
	Version string `json:"version"`
	// Name of the api
	// in: string
	Name string `json:"name"`
	// Message of the api
	// in: string
	Message string `json:"message"`
}

func GetHealth() *Health {
	var hea Health

	hea.Title = "IT works !"
	hea.Version = "0.0.1"
	hea.Name = TheAppConfig().Nameapi
	hea.Message = "Gosmo is feeling API today :)"

	return &hea
}
