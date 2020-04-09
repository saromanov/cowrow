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

	assert.Error(t, LoadByEnv("COWROW_TEST_NOT_FOUND", "test", &r))
	assert.Error(t, LoadByEnv("COWROW_TEST", "test2", &r))
	assert.Error(t, LoadByEnv("COWROW_TEST", "test_failed", &r))
}

func TestLoadByPath(t *testing.T) {
	var r testStruct
	assert.NoError(t, LoadByPath("./fixtures/test.yaml", &r))
	assert.Equal(t, "bar", r.Foo)
	assert.Equal(t, 42, r.Num)
	var r2 testStruct
	assert.NoError(t, LoadByPath("./fixtures/test", &r2))
	assert.Equal(t, "bar", r2.Foo)
	assert.Equal(t, 42, r2.Num)
	var r3 testStruct
	assert.NoError(t, LoadByPath("./fixtures/test.yml", &r3))
	assert.Equal(t, "bar", r3.Foo)
	assert.Equal(t, 42, r3.Num)
	assert.Error(t, LoadByPath("./fixtures/test22.yml", &r3))
}
