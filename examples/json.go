package main

import (
	"fmt"

	"github.com/tharun208/config-loader/pkg/config"
)

func main() {

	// Initialize config loader
	config := config.NewConfig()

	// Load configuration files

	err := config.LoadFiles("fixtures/valid.json")
	// config.LoadFiles("fixtures/valid.json", "testdata/valid_local.json")

	if err != nil {
		panic(err)
	}

	// GetValue method returns pretty version of return value

	val := config.GetConfigValueAsString("cache")

	fmt.Printf("%v\n", val)

	val = config.GetConfigValueAsString("environment")

	fmt.Printf("%v\n", val)

	// GetConfigValue returns concrete go type value
	value := config.GetConfigValue("database.port")

	fmt.Printf("%v\n", value)

}
