module github.com/ESNFranceG33kTeam/CosmoAppy

go 1.18

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gorilla/mux v1.8.0
	github.com/stretchr/testify v1.8.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-openapi/analysis v0.21.4 // indirect
	github.com/go-openapi/errors v0.20.3 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/loads v0.21.2 // indirect
	github.com/go-openapi/runtime v0.24.1
	github.com/go-openapi/spec v0.20.7 // indirect
	github.com/go-openapi/strfmt v0.21.3 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-openapi/validate v0.22.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.10.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Packages system

replace github.com/ESNFranceG33kTeam/CosmoAppy/router => ../router

replace github.com/ESNFranceG33kTeam/CosmoAppy/cocas => ../cocas

replace github.com/ESNFranceG33kTeam/CosmoAppy/database => ../database

replace github.com/ESNFranceG33kTeam/CosmoAppy/helpers => ../helpers

replace github.com/ESNFranceG33kTeam/CosmoAppy/logger => ../logger

replace github.com/ESNFranceG33kTeam/CosmoAppy/docs => ../docs

replace github.com/ESNFranceG33kTeam/CosmoAppy/modules/launcher => ../launcher

replace github.com/ESNFranceG33kTeam/CosmoAppy/modules/health => ../health

// Packages applicatif

replace github.com/ESNFranceG33kTeam/CosmoAppy/modules/adherent => ../adherent

replace github.com/ESNFranceG33kTeam/CosmoAppy/modules/volunteer => ../volunteer

replace github.com/ESNFranceG33kTeam/CosmoAppy/modules/esncard => ../esncard

replace github.com/ESNFranceG33kTeam/CosmoAppy/modules/event => ../event

replace github.com/ESNFranceG33kTeam/CosmoAppy/modules/planning => ../planning

replace github.com/ESNFranceG33kTeam/CosmoAppy/modules/money => ../money

replace github.com/ESNFranceG33kTeam/CosmoAppy/modules/report => ../report
