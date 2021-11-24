package config_test

import (
	"testing"

	"config-loader/pkg/config"

	"github.com/stretchr/testify/assert"
)

func TestShouldLoadValidJSONandReadData(t *testing.T) {
	assert := assert.New(t)

	// should not return error when the valid file is present
	data, err := config.LoadJSON("testdata/valid.json")
	assert.Nil(err)
	assert.NotEmpty(data)

	// should able to map the data
	person, ok := data["person"].(map[string]interface{})

	isAvailable, ok := data["is_available"]
	assert.True(ok)

	// should able to return the data

	assert.Equal(float64(65.45), person["weight"])
	assert.Equal(float64(24), person["age"])
	assert.Equal(true, isAvailable)
}

func TestShouldReturnErrorForNoFilePresent(t *testing.T) {
	assert := assert.New(t)
	data, err := config.LoadJSON("testdata/nofile.json")
	assert.NotNil(err)
	assert.Empty(data)
}

func TestShouldNotReturnErrorForInvalidJSONFile(t *testing.T) {
	assert := assert.New(t)
	data, err := config.LoadJSON("testdata/invalid.json")
	assert.Nil(err)
	assert.Empty(data)
}
