package config

import "testing"

func TestLoadUsesDevelopmentDefaults(t *testing.T) {
	t.Setenv("ZIPCRYPT_HTTP_ADDR", "")
	t.Setenv("ZIPCRYPT_ENVIRONMENT", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if cfg.HTTPAddr != ":8080" {
		t.Errorf("HTTPAddr = %q, want :8080", cfg.HTTPAddr)
	}
	if cfg.Environment != "development" {
		t.Errorf("Environment = %q, want development", cfg.Environment)
	}
}

func TestLoadRejectsMissingProductionDependencies(t *testing.T) {
	t.Setenv("ZIPCRYPT_ENVIRONMENT", "production")
	t.Setenv("ZIPCRYPT_DATABASE_URL", "")

	if _, err := Load(); err == nil {
		t.Fatal("Load() error = nil, want validation error")
	}
}
