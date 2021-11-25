config-loader
----

`config-loader` supports to load config files and convert to map values.

## Supported format

- json

### Usage

```go
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
```
## Running config-loader on local machine

You can run config-loader locally by running:
```bash
make run
```

## Running config-loader on local machine using docker image

Build the Image locally using:
```bash
docker build -t config-loader .
```

Run the image using:
```bash
docker run -it config-loader
```