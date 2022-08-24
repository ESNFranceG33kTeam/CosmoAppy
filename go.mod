module github.com/ESNFranceG33kTeam/sAPI

go 1.18

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gorilla/mux v1.8.0
)

replace github.com/ESNFranceG33kTeam/sAPI/config => ../config

replace github.com/ESNFranceG33kTeam/sAPI/models => ../models

replace github.com/ESNFranceG33kTeam/sAPI/controllers => ../controllers
