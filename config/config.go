package config

import (
	"fmt"
	"os"

	"github.com/gookit/color"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type Config struct {
	vip *viper.Viper
}

func NewConfig(envPath string) *Config {
	_, err := os.Stat(envPath)
	if err != nil {
		color.Redln("Please create .env and initialize it first.")
		os.Exit(0)
	}

	config := &Config{}
	config.vip = viper.New()
	config.vip.SetConfigName(envPath)
	config.vip.SetConfigType("env")
	config.vip.AddConfigPath(".")

	err = config.vip.ReadInConfig()
	if err != nil {
		color.Redln(fmt.Sprintf("invalid config error: %e", err))
	}

	config.vip.SetConfigName("framework")
	config.vip.AutomaticEnv()

	return config
}

// Env gets any type config from env (or .env)
func (config *Config) Env(envName string, defaultValue ...any) any {
	value := config.Get(envName, defaultValue...)

	if cast.ToString(value) == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}

		return nil
	}

	return value
}

// Add adds configuration to application
func (config *Config) Add(name string, configuration map[string]any) {
	config.vip.Set(name, configuration)
}

// Get gets any type config from config
func (config *Config) Get(path string, defaultValue ...any) any {
	if !config.vip.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}

		return nil
	}

	return config.vip.Get(path)
}

// GetString gets string type config from config
func (config *Config) GetString(path string, defaultValue ...any) string {
	value := config.Get(path, defaultValue...)

	if cast.ToString(value) == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0].(string)
		}

		return ""
	}

	return cast.ToString(value)
}

// GetInt gets int type config from config
func (config *Config) GetInt(path string, defaultValue ...any) int {
	value := config.Get(path, defaultValue...)

	if cast.ToString(value) == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0].(int)
		}

		return 0
	}

	return cast.ToInt(value)
}

// GetBool gets bool type config from config
func (config *Config) GetBool(path string, defaultValue ...any) bool {
	value := config.Get(path, defaultValue...)

	if cast.ToString(value) == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0].(bool)
		}

		return false
	}

	return cast.ToBool(value)
}
