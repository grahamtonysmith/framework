package config

type Config interface {
	// Env get any type config from env (or .env)
	Env(envName string, defaultValue ...any) any

	// Add configuration to application
	Add(name string, configuration map[string]any)

	// Get any type config from config
	Get(path string, defaultValue ...any) any

	// GetString string type config from config
	GetString(path string, defaultValue ...any) string

	// GetInt get int type config from config
	GetInt(path string, defaultValue ...any) int

	// GetBool get bool type config from config
	GetBool(path string, defaultValue ...any) bool
}
