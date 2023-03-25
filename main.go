package main

import (
	"log"
	"os"

	"github.com/grahamtonysmith/framework/config"
	"github.com/grahamtonysmith/framework/facades"
)

func init() {
	os.Setenv("FRAMEWORK_GREETING", "Hello, from envVar")

	sp := &config.ServiceProvider{}
	sp.Register()

	facades.Config.Add("greeting", map[string]any{
		"standard": "Hello, from config",
	})
}

func main() {
	greetingFromEnvVar := facades.Config.Env("FRAMEWORK_GREETING", "Hello, from DEFAULT envVar")
	greetingFromDotEnv := facades.Config.Env("DOTENV_GREETING", "Hello, from DEFAULT dotEnv").(string)
	greetingFromConfig := facades.Config.GetString("greeting.standard", "Hello, from default config")

	log.Println("greetingFromEnvVar:", greetingFromEnvVar)
	log.Println("greetingFromDotEnv:", greetingFromDotEnv)
	log.Println("greetingFromConfig:", greetingFromConfig)
}
