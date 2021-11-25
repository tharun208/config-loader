package config

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
