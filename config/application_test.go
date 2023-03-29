package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
	config *Application
}

func TestConfigTestSuite(t *testing.T) {
	f, err := os.Create(".env")
	if err != nil {
		panic(err)
	}
	f.Close()

	defer os.Remove(".env")

	suite.Run(t, &ConfigTestSuite{
		config: NewApplication(".env"),
	})
}

func (suite *ConfigTestSuite) TestEnv() {
	suite.Equal("framework", suite.config.Env("APP_NAME", "framework").(string))
}

func (suite *ConfigTestSuite) TestAdd() {
	suite.config.Add("app", map[string]any{
		"env": "local",
	})

	suite.Equal("local", suite.config.Get("app.env").(string))
}

func (suite *ConfigTestSuite) TestGet() {
	suite.Equal("framework", suite.config.Get("APP_NAME", "framework").(string))
}

func (suite *ConfigTestSuite) TestGetString() {
	suite.config.Add("database", map[string]any{
		"default": suite.config.Env("DB_CONNECTION", "mysql"),
		"connections": map[string]any{
			"mysql": map[string]any{
				"host": suite.config.Env("DB_HOST", "127.0.0.1"),
			},
		},
	})

	suite.Equal("mysql", suite.config.GetString("database.default"))
	suite.Equal("framework", suite.config.GetString("APP_NAME", "framework"))
	suite.Equal("127.0.0.1", suite.config.GetString("database.connections.mysql.host"))
}

func (suite *ConfigTestSuite) TestGetInt() {
	suite.Equal(1338, suite.config.GetInt("APP_PORT", 1338))
}

func (suite *ConfigTestSuite) TestGetBool() {
	suite.Equal(true, suite.config.GetBool("APP_DEBUG", true))
}
