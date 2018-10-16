package cowrow

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const envPath = "COWROW"

// LoadYAMLFile provides loading of the yaml file
// by the name from path
func LoadYAMLFile(name string, cfg interface{}) error {
	path := os.Getenv(envPath)
	fileConfig, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", path, name))
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(fileConfig, &cfg)
	if err != nil {
		return err
	}

	return nil
}
