package config

import (
	"encoding/json"
	"strings"
)

type Decoder func(data []byte) (map[string]interface{}, error)

type config struct {
	data     map[string]interface{}
	decoders map[string]Decoder
}

func NewConfig() *config {
	return &config{
		data:     make(map[string]interface{}),
		decoders: map[string]Decoder{"json": JSONDecoder},
	}
}

// GetConfigValueAsString returns the string represented value for the given key.
// The key can be a dot seperate path of nested config.
func (c *config) GetConfigValueAsString(key string) string {
	value := c.GetConfigValue(key)
	b, _ := json.MarshalIndent(value, "", "  ")
	return string(b)
}

// GetConfigValue returns the value for the given key.
// The key can be a dot seperate path of nested config.
func (c *config) GetConfigValue(key string) interface{} {
	return getValueFromKey(c.data, key)
}

func (c *config) merge(data map[string]interface{}) {
	c.data = merge(c.data, data)
}

func getValueFromKey(data map[string]interface{}, key string) interface{} {

	keys := strings.Split(key, ".")

	configuration := data
	for _, key := range keys {
		val, ok := configuration[key]
		if !ok {
			return nil
		}
		data, ok := val.(map[string]interface{})
		if ok {
			configuration = data
		} else {
			return configuration[key]
		}
	}
	return configuration
}
