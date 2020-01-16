# cowrow
Simple tool for loading of the yaml files from the directory

### Loading of config by the env variable

```go

import "github.com/saromanov/cowrow"

const envPath = "BASHTASKS"
func LoadYAMLByName(name string) (*Config, error) {
	cfg := &Config{}
	err := cowrow.LoadByEnv(envPath, name, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
```

### Loading of config by the path

```go

import "github.com/saromanov/cowrow"

const envPath = "BASHTASKS"
func LoadYAMLByName(name string) (*Config, error) {
	cfg := &Config{}
	err := cowrow.LoadByPath(name, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
```
