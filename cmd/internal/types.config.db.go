package internal

// RedisConfig redis module
type RedisConfig struct {
	Endpoint string
	Timeout  int
	MaxIdle  int
}

// ElasticConfig elastic module
type ElasticConfig struct {
	Endpoint string
}
