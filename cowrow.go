package cowrow

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v2"
)

var errPathNotFound = errors.New("path is not found")

// LoadByEnv provides loading of the yaml file
// by the name from path
func LoadByEnv(envPath, name string, cfg interface{}) error {
	path := os.Getenv(envPath)
	if path == "" {
		return errPathNotFound
	}
	return load(fmt.Sprintf("%s/%s.yaml", path, name), cfg)
}

// LoadByPath provides loading of the config by the path
func LoadByPath(path string, cfg interface{}) error {
	return load(path, cfg)
}

// load provides loading of the config
func load(name string, cfg interface{}) error {
	fileConfig, err := loadFile(name)
	if err != nil {
		return fmt.Errorf("unable to load file: %v", err)
	}
	err = yaml.Unmarshal(fileConfig, cfg)
	if err != nil {
		return fmt.Errorf("unable to unmarshal data: %v", err)
	}

	return nil
}

// loadFile provides loading of file
// its trying to load by .yaml or .yml extension
func loadFile(name string) ([]byte, error) {
	if !strings.Contains(name, ".yaml") && !strings.Contains(name, ".yml") {
		name = name + ".yml"
	}
	fileConfig, err := ioutil.ReadFile(name)
	if err != nil {
		fname := strings.TrimSuffix(name, path.Ext(name))
		if strings.Contains(name, ".yaml") {
			fname = fname + ".yml"
		}
		if strings.Contains(name, ".yml") {
			fname = fname + ".yaml"
		}
		fileConfig, err = ioutil.ReadFile(fname)
	}
	if err != nil {
		return nil, fmt.Errorf("unable to load file: %v", err)
	}
	return fileConfig, nil
}
