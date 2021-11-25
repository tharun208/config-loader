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

	val := config.GetValue("cache")

	fmt.Printf("%v\n", val)

	val = config.GetValue("environment")

	fmt.Printf("%v\n", val)

	// GetConfigValue returns concrete go type value
	value := config.GetConfigValue("database.port")

	fmt.Printf("%v\n", value)

}
