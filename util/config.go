package util

import (
	"time"

	"github.com/spf13/viper"
)

// POSTGRES_URL=postgresql://postgres:secret@localhost:5432/monke_bank?sslmode=disable
// ENVIRONMENT=development
// MIGRATION_URL=file://db/migration
// TOKEN_SYMMETRIC_KEY=12345678901234567890123456789012
// ACCESS_TOKEN_DURATION=15m
// REFRESH_TOKEN_DURATION=24h

type Config struct {
	Enviroment           string        `mapstructure:"ENVIROMENT"`
	PostgresURL          string        `mapstructure:"POSTGRES_URL"`
	MigrationURL         string        `mapstructure:"MIGRATION_URL"`
	TokenSymetricKey     string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return

}
