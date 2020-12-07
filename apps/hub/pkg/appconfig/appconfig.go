package appconfig

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/atastrophic/go-ms-with-eks/pkg/exception"
	"github.com/imdario/mergo"
	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	SQL SQL `yaml:"sql"`
}
type DSN struct {
	Readwrite bool   `yaml:"readwrite"`
	URL       string `yaml:"url"`
	Protocol  string `yaml:"protocol"`
	Port      int    `yaml:"port"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Database  string `yaml:"database"`
}
type SQL struct {
	DSN DSN `yaml:"dsn"`
}

func NewConfig(fallback, app string) *AppConfig {

	configFile, _ := os.Open(fmt.Sprintf("./conf/%s.yaml", fallback))
	configBytes, err := ioutil.ReadAll(configFile)
	exception.WithError(err)
	var config AppConfig
	err = yaml.Unmarshal(configBytes, &config)
	exception.WithError(err)

	appConfigFile, _ := os.Open(fmt.Sprintf("./conf/%s.yaml", app))
	appConfigBytes, err := ioutil.ReadAll(appConfigFile)
	exception.WithError(err)

	var appConfig AppConfig
	err = yaml.Unmarshal(appConfigBytes, &appConfig)
	exception.WithError(err)

	err = mergo.Merge(&appConfig, config)
	exception.WithError(err)

	return &appConfig
}
