package cowrow

import (
	"os"
	"path/filepath"
	"testing"
)

var rootTestDir = "./tmp"

func removeContentFromRoot(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

type testStruct struct {
	Foo  string          `yaml:"foo"`
	Data []testStructEmb `yaml:"data"`
}
type testStructEmb struct {
	Name string `yaml:"name"`
}

func TestLoadByEnv(t *testing.T) {
	os.Setenv("COWROW_TEST", "./fixtures")
	LoadByEnv("COWROW_TEST", "test.yaml")
}
