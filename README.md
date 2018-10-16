# cowrow
Simple tool for loading of the yaml files from the directory

## Example
```go

import "github.com/saromanov/cowrow"

const envPath = "BASHTASKS"
func LoadYAMLByName(name string) (*Config, error) {
	cfg := &Config{}
	err := cowrow.LoadYAMLFile(envPath, name, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
```