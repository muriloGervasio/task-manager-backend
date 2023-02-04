package connection

import (
	"app/config"

	"github.com/andygrunwald/go-jira"
)

var Client *jira.Client

func init() {

	var tp = jira.BasicAuthTransport{
		Username: config.Config.Jira.Username,
		Password: config.Config.Jira.Password,
	}

	Client, _ = jira.NewClient(tp.Client(), "https://graodireto.atlassian.net/")
}
