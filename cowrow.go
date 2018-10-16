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
	fileConfig, err := ioutil.ReadFile(fmt.Sprintf("%s/%s.yaml", path, name))
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(fileConfig, &cfg)
	if err != nil {
		return err
	}

	return nil
}
