package config

import (
	"fmt"
	"os"

	"github.com/gookit/color"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type Application struct {
	vip *viper.Viper
}

func NewApplication(envPath string) *Application {
	_, err := os.Stat(envPath)
	if err != nil {
		color.Redln("Please create .env and initialize it first.")
		os.Exit(0)
	}

	app := &Application{}
	app.vip = viper.New()
	app.vip.SetConfigName(envPath)
	app.vip.SetConfigType("env")
	app.vip.AddConfigPath(".")

	err = app.vip.ReadInConfig()
	if err != nil {
		color.Redln(fmt.Sprintf("invalid config error: %e", err))
	}

	app.vip.SetConfigName("framework")
	app.vip.AutomaticEnv()

	return app
}

// Env gets any type config from env (or .env)
func (app *Application) Env(envName string, defaultValue ...any) any {
	value := app.Get(envName, defaultValue...)

	if cast.ToString(value) == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}

		return nil
	}

	return value
}

// Add adds configuration to application
func (app *Application) Add(name string, configuration map[string]any) {
	app.vip.Set(name, configuration)
}

// Get gets any type config from config
func (app *Application) Get(path string, defaultValue ...any) any {
	if !app.vip.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}

		return nil
	}

	return app.vip.Get(path)
}

// GetString gets string type config from config
func (app *Application) GetString(path string, defaultValue ...any) string {
	value := app.Get(path, defaultValue...)

	if cast.ToString(value) == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0].(string)
		}

		return ""
	}

	return cast.ToString(value)
}

// GetInt gets int type config from config
func (app *Application) GetInt(path string, defaultValue ...any) int {
	value := app.Get(path, defaultValue...)

	if cast.ToString(value) == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0].(int)
		}

		return 0
	}

	return cast.ToInt(value)
}

// GetBool gets bool type config from config
func (app *Application) GetBool(path string, defaultValue ...any) bool {
	value := app.Get(path, defaultValue...)

	if cast.ToString(value) == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0].(bool)
		}

		return false
	}

	return cast.ToBool(value)
}
