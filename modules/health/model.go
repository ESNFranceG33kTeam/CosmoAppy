package health

// swagger:model Health
type Health struct {
	// Title of the api page
	// in: string
	Title string `json:"title"`
	// Message of the api
	// in: string
	Message string `json:"message"`
}

// swagger:model Status
type Status struct {
	// Title of the api page
	// in: string
	Title string `json:"title"`
	// Name of the api
	// in: string
	Name string `json:"name"`
	// Version of the api
	// in: string
	Version string `json:"version"`
	// Message of the api
	// in: string
	Message string `json:"message"`
	// DbStatus of the api
	// in: string
	DbStatus string `json:"dbstatus"`
}

func GetHealth() *Health {
	var hea Health

	hea.Title = "IT works !"
	hea.Message = "Gosmo is feeling API today :)"

	return &hea
}

func GetStatus() *Status {
	var sta Status

	sta.Title = "Status"
	sta.Name = TheAppConfig().Nameapi
	sta.Version = "0.0.1"
	sta.Message = "Gosmo is feeling API today :)"

	err := TheDb().Ping()
	if err != nil {
		TheLogger().LogCritical("health", "status db failed.", err, false)
		sta.DbStatus = "failed : " + err.Error()
	} else {
		sta.DbStatus = "everything is awesome <3"
	}

	return &sta
}
