package config

import (
    "log"
    "github.com/spf13/viper"
)

// Env holds the configuration settings for the application
type Env struct {
    AppEnv                 string `mapstructure:"APP_ENV"`
    ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
    ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
    MongoURI               string `mapstructure:"MONGO_URI"`
    DBName                 string `mapstructure:"DB_NAME"`
    AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
    RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
    AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
    RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
    SMTPUsername           string `mapstructure:"SMTP_USERNAME"`
    SMTPPassword           string `mapstructure:"SMTP_PASSWORD"`
    SMTPHost               string `mapstructure:"SMTP_HOST"`
    SMTPPort               string `mapstructure:"SMTP_PORT"`
}

// NewEnv initializes and returns a new Env instance with settings from .env file
func NewEnv() *Env {
    var env Env
    viper.SetConfigFile(".env")

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Can't find the file .env : %v", err)
    }

    if err := viper.Unmarshal(&env); err != nil {
        log.Fatalf("Environment can't be loaded: %v", err)
    }

    if env.AppEnv == "development" {
        log.Println("The App is running in development env")
    }

    return &env
}
