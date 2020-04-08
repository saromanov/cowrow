package cowrow

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Foo  string          `yaml:"foo"`
	Num  int             `yaml:"num"`
	Data []testStructEmb `yaml:"data"`
}
type testStructEmb struct {
	Name string `yaml:"name"`
}

func TestLoadByEnv(t *testing.T) {
	os.Setenv("COWROW_TEST", "./fixtures")
	var r testStruct
	assert.NoError(t, LoadByEnv("COWROW_TEST", "test", &r))
	assert.Equal(t, "bar", r.Foo)
	assert.Equal(t, 42, r.Num)
}
