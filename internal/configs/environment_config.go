package configs

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type EnvironmentConfig struct {
	IrisClient IrisClientConfig
	Kafka      KafkaConfig
}

type IrisClientConfig struct {
	Target string `envconfig:"IRIS_TARGET"`
}

type KafkaConfig struct {
	Target string `envconfig:"KAFKA_TARGET"`
}

var config EnvironmentConfig

func GetEnvironmentConfig() *EnvironmentConfig {
	if config != (EnvironmentConfig{}) {
		return &config
	}

	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal("Could not load Configs: ", err)
	}

	return &config
}
