package config

type Config interface {
	// Returns config from env or dotenv file
	Env(envName string, defaultValue ...any) any

	// Adds configuration to the application
	Add(name string, configuration map[string]any)

	// Retuns config from the application
	Get(path string, defaultValue ...any) any

	// Returns config from the application
	GetString(path string, defaultValue ...any) string

	// Returns config from the application
	GetInt(path string, defaultValue ...any) int

	// Returns config from the application
	GetBool(path string, defaultValue ...any) bool
}
