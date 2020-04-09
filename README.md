# cowrow
![Go](https://github.com/saromanov/cowrow/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/saromanov/cowrow)](https://goreportcard.com/report/github.com/saromanov/cowrow)
[![Coverage Status](https://coveralls.io/repos/github/saromanov/cowrow/badge.svg?branch=master)](https://coveralls.io/github/saromanov/cowrow?branch=master)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/d2ec3aa3607a469f95061bf6fd9f51d0)](https://app.codacy.com/manual/saromanov/cowrow?utm_source=github.com&utm_medium=referral&utm_content=saromanov/cowrow&utm_campaign=Badge_Grade_Dashboard)


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

LICENSE
MIT
