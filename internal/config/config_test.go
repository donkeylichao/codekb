package config

import "testing"

func TestDefaultIsValid(t *testing.T) {
	cfg := Default("example")

	if err := cfg.Validate(); err != nil {
		t.Fatalf("Default().Validate() error = %v", err)
	}
	if cfg.Embedding.Enabled {
		t.Fatal("embedding must be disabled by default")
	}
}

func TestValidateRejectsInvalidValues(t *testing.T) {
	tests := []struct {
		name   string
		change func(*Config)
	}{
		{"schema version", func(c *Config) { c.SchemaVersion++ }},
		{"project name", func(c *Config) { c.Project.Name = "" }},
		{"project root", func(c *Config) { c.Project.Root = "" }},
		{"include", func(c *Config) { c.Index.Include = nil }},
		{"file size", func(c *Config) { c.Index.MaxFileBytes = 0 }},
		{"search limits", func(c *Config) { c.Search.MaxLimit = c.Search.DefaultLimit - 1 }},
		{"embedding batch size", func(c *Config) { c.Embedding.BatchSize = 0 }},
		{"enabled embedding", func(c *Config) { c.Embedding.Enabled, c.Embedding.Model = true, "" }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := Default("example")
			tt.change(&cfg)
			if err := cfg.Validate(); err == nil {
				t.Fatal("Validate() error = nil, want error")
			}
		})
	}
}
