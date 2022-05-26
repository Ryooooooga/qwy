package config

import (
	"io"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	FinderCommand string        `yaml:"finder-command"`
	Finder        FinderOptions `yaml:"finder"`
	Completions   Completions   `yaml:"completions"`
}

type Completions []*Completion

type Completion struct {
	Description   string        `yaml:"description"`
	Patterns      []string      `yaml:"patterns"`
	Source        string        `yaml:"source"`
	Finder        FinderOptions `yaml:"finder"`
	UnescapeQuery bool          `yaml:"unescape-query"`
	Callback      string        `yaml:"callback"`
}

type FinderOptions map[string]any

func defaultConfig() *Config {
	return &Config{
		FinderCommand: "fzf",
		Finder:        FinderOptions{},
		Completions:   Completions{},
	}
}

func LoadConfig() (*Config, error) {
	configPath, err := DefaultRulePathLoader.ConfigPath()
	if err != nil {
		return nil, err
	}

	config, err := loadConfigFromFile(configPath)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func loadConfig(r io.Reader) (*Config, error) {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	config := defaultConfig()
	if err := yaml.Unmarshal(bytes, config); err != nil {
		return nil, err
	}

	for _, c := range config.Completions {
		c.Finder = mergeMap(config.Finder, c.Finder)
	}

	return config, nil
}

func loadConfigFromFile(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return loadConfig(f)
}

func LoadConfigFromText(text string) (*Config, error) {
	r := strings.NewReader(text)
	return loadConfig(r)
}

func mergeMap[K comparable, V any](a, b map[K]V) map[K]V {
	dest := make(map[K]V, len(a)+len(b))
	for k, v := range a {
		dest[k] = v
	}
	for k, v := range b {
		dest[k] = v
	}
	return dest
}
