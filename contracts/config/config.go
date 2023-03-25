package config

type Config interface {
	// Env gets any type config from env (or .env)
	Env(envName string, defaultValue ...any) any

	// Add adds configuration to application
	Add(name string, configuration map[string]any)

	// Get gets any type config from config
	Get(path string, defaultValue ...any) any

	// GetString gets string type config from config
	GetString(path string, defaultValue ...any) string

	// GetInt gets int type config from config
	GetInt(path string, defaultValue ...any) int

	// GetBool gets bool type config from config
	GetBool(path string, defaultValue ...any) bool
}
