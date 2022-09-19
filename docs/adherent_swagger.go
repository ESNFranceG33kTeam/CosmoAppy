package docs

import "github.com/ESNFranceG33kTeam/sAPI/models"

// swagger:model GetAdherents
type GetAdherents struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the response
	// in: string
	Message string            `json:"message"`
	Data    *models.Adherents `json:"data"`
}

// swagger:model GetAdherent
type GetAdherent struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the response
	// in: string
	Message string           `json:"message"`
	Data    *models.Adherent `json:"data"`
}

// swagger:parameters adherent showAdherent
type ShowAdherent struct {
	// name: id
	// in: path
	// type: integer
	// required: true
	Id int64 `json:"id"`
}
