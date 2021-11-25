package cmd

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigLoaderCmd(t *testing.T) {
	cmd := DefaultRootCmd()
	b := &bytes.Buffer{}
	cmd.SetOut(b)

	assert := assert.New(t)
	t.Run("Should able to load file(s)", func(t *testing.T) {
		cmd.SetArgs([]string{"load", filepath.Join("../..", "fixtures", "valid.json")})
		err := cmd.Execute()
		assert.Nil(err)
	})
	t.Run("Should able to get value from loaded configuration", func(t *testing.T) {
		cmd.SetArgs([]string{"get", "environment"})
		err := cmd.Execute()
		assert.Nil(err)
		out, err := ioutil.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal("production", strings.TrimSpace(string(out)))

		cmd.SetArgs([]string{"get", "database.host"})
		err = cmd.Execute()
		assert.Nil(err)
		out, err = ioutil.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal("mysql", strings.TrimSpace(string(out)))

	})
}
