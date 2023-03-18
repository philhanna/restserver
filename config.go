package webserver

import (
	"log"
	"os"

	"github.com/ghodss/yaml"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

type Configuration struct {
	HOST   string `json:"host"`
	PORT   int    `json:"port"`
	DBNAME string `json:"dbname"`
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

func NewConfiguration(configfile string) Configuration {
	yamlBlob, err := os.ReadFile(configfile)
	if err != nil {
		log.Fatal(err)
	}
	p := new(Configuration)
	err = yaml.Unmarshal(yamlBlob, p)
	if err != nil {
		log.Fatal(err)
	}
	return *p
}
