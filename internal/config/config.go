// Package config loads and validates application configuration.
package config

import (
	"fmt"
	"os"
)

// Config contains values supplied by the environment. Keeping configuration in
// one typed struct makes required dependencies explicit at program startup.
type Config struct {
	HTTPAddr    string
	Environment string
	LogLevel    string

	DatabaseURL string
	RedisAddr   string

	S3Endpoint  string
	S3AccessKey string
	S3SecretKey string
	S3Bucket    string
	S3UseSSL    bool
}

// Load reads configuration. Development defaults make the first local run
// pleasant; production validates that no unsafe default has slipped through.
func Load() (Config, error) {
	cfg := Config{
		HTTPAddr:    value("ZIPCRYPT_HTTP_ADDR", ":8080"),
		Environment: value("ZIPCRYPT_ENVIRONMENT", "development"),
		LogLevel:    value("ZIPCRYPT_LOG_LEVEL", "info"),
		DatabaseURL: value("ZIPCRYPT_DATABASE_URL", ""),
		RedisAddr:   value("ZIPCRYPT_REDIS_ADDR", ""),
		S3Endpoint:  value("ZIPCRYPT_S3_ENDPOINT", ""),
		S3AccessKey: value("ZIPCRYPT_S3_ACCESS_KEY", ""),
		S3SecretKey: value("ZIPCRYPT_S3_SECRET_KEY", ""),
		S3Bucket:    value("ZIPCRYPT_S3_BUCKET", ""),
		S3UseSSL:    value("ZIPCRYPT_S3_USE_SSL", "false") == "true",
	}

	if cfg.Environment == "production" {
		for name, v := range map[string]string{
			"ZIPCRYPT_DATABASE_URL":  cfg.DatabaseURL,
			"ZIPCRYPT_REDIS_ADDR":    cfg.RedisAddr,
			"ZIPCRYPT_S3_ENDPOINT":   cfg.S3Endpoint,
			"ZIPCRYPT_S3_ACCESS_KEY": cfg.S3AccessKey,
			"ZIPCRYPT_S3_SECRET_KEY": cfg.S3SecretKey,
			"ZIPCRYPT_S3_BUCKET":     cfg.S3Bucket,
		} {
			if v == "" {
				return Config{}, fmt.Errorf("%s is required in production", name)
			}
		}
	}

	return cfg, nil
}

func value(name, fallback string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	return fallback
}
