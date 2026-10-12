package config

import (
	"errors"
	"fmt"
)

const SchemaVersion = 1

// Config is the versioned project configuration written to .codekb/config.yaml.
type Config struct {
	SchemaVersion int       `yaml:"schema_version"`
	Project       Project   `yaml:"project"`
	Index         Index     `yaml:"index"`
	Search        Search    `yaml:"search"`
	Embedding     Embedding `yaml:"embedding"`
}

type Project struct {
	Name string `yaml:"name"`
	Root string `yaml:"root"`
}

type Index struct {
	Include      []string `yaml:"include"`
	Exclude      []string `yaml:"exclude"`
	MaxFileBytes int64    `yaml:"max_file_bytes"`
}

type Search struct {
	DefaultLimit int `yaml:"default_limit"`
	MaxLimit     int `yaml:"max_limit"`
}

type Embedding struct {
	Enabled   bool   `yaml:"enabled"`
	Provider  string `yaml:"provider"`
	BaseURL   string `yaml:"base_url"`
	Model     string `yaml:"model"`
	BatchSize int    `yaml:"batch_size"`
}

// Default returns the initial offline-first project configuration.
func Default(projectName string) Config {
	return Config{
		SchemaVersion: SchemaVersion,
		Project: Project{
			Name: projectName,
			Root: "..",
		},
		Index: Index{
			Include: []string{"**/*.go", "**/*.md", "**/*.txt"},
			Exclude: []string{
				".git/**", ".codekb/**", "vendor/**", "node_modules/**",
				"**/.env*", "**/*secret*", "**/*.pem", "**/*.key", "**/*_generated.go",
			},
			MaxFileBytes: 1 << 20,
		},
		Search: Search{
			DefaultLimit: 10,
			MaxLimit:     50,
		},
		Embedding: Embedding{
			Provider:  "ollama",
			BaseURL:   "http://127.0.0.1:11434",
			Model:     "nomic-embed-text",
			BatchSize: 16,
		},
	}
}

// Validate checks values that do not require filesystem access.
func (c Config) Validate() error {
	if c.SchemaVersion != SchemaVersion {
		return fmt.Errorf("unsupported schema_version %d", c.SchemaVersion)
	}
	if c.Project.Name == "" {
		return errors.New("project.name is required")
	}
	if c.Project.Root == "" {
		return errors.New("project.root is required")
	}
	if len(c.Index.Include) == 0 {
		return errors.New("index.include must not be empty")
	}
	if c.Index.MaxFileBytes <= 0 {
		return errors.New("index.max_file_bytes must be positive")
	}
	if c.Search.DefaultLimit <= 0 || c.Search.MaxLimit < c.Search.DefaultLimit {
		return errors.New("search limits must be positive and max_limit must be at least default_limit")
	}
	if c.Embedding.BatchSize <= 0 {
		return errors.New("embedding.batch_size must be positive")
	}
	if c.Embedding.Enabled && (c.Embedding.Provider == "" || c.Embedding.BaseURL == "" || c.Embedding.Model == "") {
		return errors.New("enabled embedding requires provider, base_url, and model")
	}
	return nil
}
