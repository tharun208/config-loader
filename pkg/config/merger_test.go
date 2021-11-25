package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldAbletoMergeTwoDifferentObject(t *testing.T) {
	assert := assert.New(t)

	expected := map[string]interface{}{
		"name":      "test-user",
		"available": true,
		"location": map[string]interface {
		}{"country": "unknown"},
		"phone": 456,
	}

	data := map[string]interface{}{
		"name":      "test-user",
		"available": true,
	}

	dataToMerge := map[string]interface{}{
		"location": map[string]interface {
		}{"country": "unknown"},
		"phone": 456,
	}

	actual := merge(data, dataToMerge)

	assert.Equal(expected, actual)
}

func TestShouldAbletoMergeAndUpdateValues(t *testing.T) {
	assert := assert.New(t)
	expected := map[string]interface{}{
		"name":      "test-user",
		"available": false,
		"code":      123,
		"dont_merge": map[string]interface{}{
			"val": "dont-merge",
		},
	}

	data := map[string]interface{}{
		"name":      "test-user-new",
		"available": true,
		"code":      123,
		"dont_merge": map[string]interface{}{
			"val": "dont-merge",
		},
	}

	dataToMergeAndUpdate := map[string]interface{}{
		"name":      "test-user",
		"available": false,
		"code":      123,
		"dont_merge": map[string]interface{}{
			"val": "dont-merge",
		},
	}

	actual := merge(data, dataToMergeAndUpdate)

	assert.Equal(expected, actual)
}
