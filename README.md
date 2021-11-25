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
	val := config.GetValue("cache")

	val = config.GetValue("environment")

	// GetConfigValue returns concrete go type value
	value := config.GetConfigValue("database.port")

}
```