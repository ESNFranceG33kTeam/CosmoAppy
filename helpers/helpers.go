package helpers

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

//const configPath = "conf.yml"

type Cfg struct {
	ConfPath string
	Userdb   string `yaml:"userdb"`
	Passdb   string `yaml:"passdb"`
	Ipdb     string `yaml:"ipdb"`
	Portdb   string `yaml:"portdb"`
	Namedb   string `yaml:"namedb"`
	Extradb  string `yaml:"extradb"`
	Loglevel int    `yaml:"loglevel"`
}

var AppConfig Cfg

func InitFile(confpath string) {
	AppConfig.ConfPath = confpath
}

func ReadConfig() {
	f, err := os.Open(AppConfig.ConfPath)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&AppConfig)

	if err != nil {
		fmt.Println(err)
	}
}

func TheAppConfig() *Cfg {
	return &AppConfig
}