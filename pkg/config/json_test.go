package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldLoadValidJSONandReadData(t *testing.T) {
	assert := assert.New(t)

	// should not return error when the json is invalid
	data, err := loadJSON([]byte(`{
		"person": {
			"name": "user",
			"age": 24,
			"weight": 65.45
		},
		"is_available": true
	}`))
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

func TestShouldNotReturnErrorForInvalidJSONFile(t *testing.T) {
	assert := assert.New(t)
	data, err := loadJSON([]byte(``))
	assert.NotNil(err)
	assert.Empty(data)
}
