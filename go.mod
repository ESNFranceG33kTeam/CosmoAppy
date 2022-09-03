module github.com/ESNFranceG33kTeam/sAPI

go 1.18

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gorilla/mux v1.8.0
	github.com/stretchr/testify v1.8.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/ESNFranceG33kTeam/sAPI/config => ../config

replace github.com/ESNFranceG33kTeam/sAPI/models => ../models

replace github.com/ESNFranceG33kTeam/sAPI/controllers => ../controllers

replace github.com/ESNFranceG33kTeam/sAPI/logger => ../logger

replace github.com/ESNFranceG33kTeam/sAPI/helpers => ../helpers
