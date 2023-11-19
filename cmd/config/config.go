package config

import (
	"log"
	config "tfg/organization-service/pkg/config/common"
	"tfg/organization-service/pkg/libs/sql"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	CommonConfigs config.Config
	DBConfig      sql.DBConfig
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	configs := &Config{
		Port:     config.GetEnv("PORT", "8080"),
		DBConfig: sql.NewDBConfig(),
	}

	if commonConfigs := config.Common(); commonConfigs != nil {
		configs.CommonConfigs = *commonConfigs
	}

	return configs, nil
}
