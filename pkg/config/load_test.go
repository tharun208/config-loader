package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFiles(t *testing.T) {
	st := assert.New(t)

	config := NewConfig()

	// Should able to load multiple files
	err := config.LoadFiles("testdata/valid.json", "testdata/other.json")
	st.Nil(err)

	// Should not return error when loading a invalid content file
	err = config.LoadFiles("testdata/invalid.json")
	st.Nil(err)

	// Should return error when loading a file not present
	err = config.LoadFiles("testdata")
	st.NotNil(err)

	// Should return error when loading a different file format 
	err = config.LoadFiles("testdata/test.yaml")
	st.NotNil(err)
}
