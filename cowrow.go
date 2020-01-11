package cowrow

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var errPathNotFound = errors.New("path is not found")

// LoadYAMLFile provides loading of the yaml file
// by the name from path
func LoadYAMLFile(envPath, name string, cfg interface{}) error {
	path := os.Getenv(envPath)
	if path == "" {
		return errPathNotFound
	}

	return load(path, name, cfg)
}

// LoadYAMLByPath provides loading of the config by the path
func LoadYAMLByPath(path, name string, cfg interface{}) error {
	return load(path, name, cfg)
}

func load(path, name string, cfg interface{}) error {
	fileConfig, err := ioutil.ReadFile(fmt.Sprintf("%s/%s.yaml", path, name))
	if err != nil {
		return fmt.Errorf("unable to load file: %v", err)
	}
	err = yaml.Unmarshal(fileConfig, cfg)
	if err != nil {
		return fmt.Errorf("unable to unmarshal data: %v", err)
	}

	return nil
}
