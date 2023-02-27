package cocas

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/ESNFranceG33kTeam/CosmoAppy/helpers"
	"github.com/ESNFranceG33kTeam/CosmoAppy/logger"
	"gopkg.in/cas.v2"
	"gopkg.in/yaml.v2"
)

var cocas *http.Server
var cocasmwd *http.ServeMux

type myHandler struct{}

var MyHandler = &myHandler{}

// swagger:model GalaxyConf
type GalaxyConf struct {
	ConfUrl       string
	CASServer     string   `yaml:"cas_server"`
	Country       string   `yaml:"country"`
	G33kTeam      []string `yaml:"galaxy_username"`
	ExtraUsername []string `yaml:"extra_username"`
}

var GalaxyUsers GalaxyConf

func InitCas() {
	GalaxyUsers.ConfUrl = helpers.TheAppConfig().G33kTeam

	// G33kTeam username
	g33k_username, err := http.Get(GalaxyUsers.ConfUrl)
	if err != nil {
		logger.GetLogger().LogError("cocas", "g33kteam url file unavailable", err)
	}
	defer g33k_username.Body.Close()
	decoder := yaml.NewDecoder(g33k_username.Body)
	err = decoder.Decode(&GalaxyUsers)
	if err != nil {
		fmt.Println(err)
	}
	// Extra username
	GalaxyUsers.ExtraUsername = helpers.TheAppConfig().ExtraAdmin[:]

	// Server init
	cocasmwd = http.NewServeMux()
	cocasmwd.Handle("/", MyHandler)

	url, _ := url.Parse(GalaxyUsers.CASServer)
	client := cas.NewClient(&cas.Options{
		URL: url,
	})

	cocas = &http.Server{
		Addr:    ":" + helpers.TheAppConfig().PortCas,
		Handler: client.Handle(cocasmwd),
	}

	logger.GetLogger().LogInfo("cocas", "Conf up !")
}

func GetCasServer() *http.Server {
	return cocas
}
