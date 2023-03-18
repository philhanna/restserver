package restserver

import (
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
)

// ---------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------

const (
	PACKAGE_NAME   = "restserver"
	YAML_FILE_NAME = "config.yml"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

type Configuration struct {
	HOST   string `json:"host"`
	PORT   int    `json:"port"`
	DBNAME string `json:"dbname"`
	DBSQL  string `json:"dbsql"`
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewConfiguration creates a configuration structure from the YAML file
// in the user configuration directory.
func NewConfiguration() (*Configuration, error) {

	// Get the configuration file name
	configfile, err := ConfigurationFile()
	if err != nil {
		return nil, err
	}

	// Load its data
	yamlBlob, err := os.ReadFile(configfile)
	if err != nil {
		return nil, err
	}

	// Create a configuration structure from the YAML
	p := new(Configuration)
	err = yaml.Unmarshal(yamlBlob, p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// ConfigurationFile returns the fully qualified name of the YAML file.
func ConfigurationFile() (string, error) {

	// Start with the user configuration directory (on Unix, "$HOME/.config")
	dirname, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	// Add the path to the yaml
	configfile := filepath.Join(dirname, PACKAGE_NAME, YAML_FILE_NAME)

	return configfile, nil
}
