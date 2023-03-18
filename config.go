package restserver

import (
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

// NewConfiguration creates a configuration structure from the specified
// YAML file.
func NewConfiguration(configfile string) (Configuration, error) {
	yamlBlob, err := os.ReadFile(configfile)
	if err != nil {
		return Configuration{}, err
	}
	p := new(Configuration)
	err = yaml.Unmarshal(yamlBlob, p)
	if err != nil {
		return *p, err
	}
	return *p, nil
}
