package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newDummyConfigPathLoader(envs map[string]string, home string) *ConfigPathLoader {
	return &ConfigPathLoader{
		lookupEnv: func(key string) (string, bool) {
			v, ok := envs[key]
			return v, ok
		},
		userHomeDir: func() (string, error) {
			return home, nil
		},
	}
}

func TestConfigPath(t *testing.T) {
	scenarios := []struct {
		envs     map[string]string
		home     string
		expected string
	}{
		{
			envs: map[string]string{
				"QWY_CONFIG_FILE": "/home/test/.qwy/config.yaml",
				"XDG_CONFIG_HOME": "/home/test/.config",
			},
			home:     "/home/test",
			expected: "/home/test/.qwy/config.yaml",
		},
		{
			envs: map[string]string{
				"XDG_CONFIG_HOME": "/home/test/config",
			},
			home:     "/home/test",
			expected: "/home/test/config/qwy/config.yaml",
		},
		{
			envs:     map[string]string{},
			home:     "/home/test",
			expected: "/home/test/.config/qwy/config.yaml",
		},
	}

	for _, s := range scenarios {
		t.Run(s.expected, func(t *testing.T) {
			r := newDummyConfigPathLoader(s.envs, s.home)

			actual, err := r.ConfigPath()
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, s.expected, actual)
		})
	}
}
