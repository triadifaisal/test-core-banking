package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type ConfigKafka struct {
	Host    string `env:"KAFKA_HOST,default:kafka:9092"`
	GroupID string `env:"KAFKA_GROUP_ID,default:core-banking"`
}

// NewKafkaConfig creates an instance of Config.
// It needs the path of the env file to be used.
func NewKafkaConfig(env string) (*ConfigKafka, error) {
	// just skip loading env files if it is not exists, env files only used in local dev
	_ = godotenv.Load(env)

	var config ConfigKafka
	if err := envdecode.Decode(&config); err != nil {
		return nil, errors.Wrap(err, "[NewConfig] error decoding env")
	}

	return &config, nil
}
