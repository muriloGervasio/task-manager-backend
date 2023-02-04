package config

import (
	"os"

	"github.com/joho/godotenv"
)

type JiraVariables struct {
	Username string
	Password string
}

type SConfig struct {
	Jira JiraVariables
}

var Config SConfig

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("no env file")
	}

	Config = SConfig{
		Jira: JiraVariables{
			Username: getEnvVariable("JIRA_USERNAME"),
			Password: getEnvVariable("JIRA_PASSWORD"),
		},
	}
}

func getEnvVariable(key string) string {
	return os.Getenv(key)
}
