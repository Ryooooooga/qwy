package config

import (
	"os"
	"path"
)

const (
	QWY_CONFIG_FILE string = "QWY_CONFIG_FILE"
	XDG_CONFIG_HOME string = "XDG_CONFIG_HOME"
)

type ConfigPathLoader struct {
	lookupEnv   func(key string) (string, bool)
	userHomeDir func() (string, error)
}

var DefaultRulePathLoader = ConfigPathLoader{
	lookupEnv:   os.LookupEnv,
	userHomeDir: os.UserHomeDir,
}

/// ConfigPath returns config file path
/// 2. $QWY_CONFIG_FILE
/// 3. $XDG_CONFIG_HOME/qwy/config.yaml
/// 4. $HOME/.config/qwy/config.yaml
func (r *ConfigPathLoader) ConfigPath() (string, error) {
	if configPath, ok := r.lookupEnv(QWY_CONFIG_FILE); ok {
		return configPath, nil
	}
	if xdgConfigHome, ok := r.lookupEnv(XDG_CONFIG_HOME); ok {
		return path.Join(xdgConfigHome, "qwy", "config.yaml"), nil
	}

	home, err := r.userHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(home, ".config", "qwy", "config.yaml"), nil
}
