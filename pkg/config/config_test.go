package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadAndGetConfigFromKey(t *testing.T) {
	assert := assert.New(t)

	config := NewConfig()

	// Should able to load multiple files
	err := config.LoadFiles("testdata/valid.json", "testdata/other.json")
	assert.Nil(err)

	tests := []struct {
		given string
		want  interface{}
	}{{
		given: "is_available",
		want:  false,
	}, {
		given: "person.name",
		want:  "user",
	}, {
		given: "person",
		want: map[string]interface{}{
			"name":   "user",
			"age":    float64(24),
			"weight": 65.45,
		},
	}}

	for _, tt := range tests {
		got := config.GetConfigValue(tt.given)
		if !assert.Equal(tt.want, got) {
			t.Errorf("expected %v but got %v", tt.want, got)
		}
	}
}
