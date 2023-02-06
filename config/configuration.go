package config

import (
	"os"

	"github.com/joho/godotenv"
)

type JiraVariables struct {
	Username string
	Password string
	Url      string
}

type SConfig struct {
	Jira JiraVariables
	Port string
}

var Config SConfig

func init() {
	godotenv.Load(".env")

	Config = SConfig{
		Jira: JiraVariables{
			Username: getEnvVariable("JIRA_USERNAME"),
			Password: getEnvVariable("JIRA_PASSWORD"),
			Url:      getEnvVariable("JIRA_URL"),
		},
		Port: getEnvVariable("PORT"),
	}
}

func getEnvVariable(key string) string {
	return os.Getenv(key)
}
