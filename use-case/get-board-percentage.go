package useCase

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"

	"github.com/andygrunwald/go-jira"
)

func GetActiveTaskPercentage(client jira.Client, project string) float64 {
	jql := fmt.Sprintf("project = \"%s\" AND cf[10020] in openSprints()", project)

	chunk, _, _ := client.Issue.Search(jql, &jira.SearchOptions{})
	total, done := 0, 0

	r, _ := regexp.Compile("(customfield_10016\":)([0-9]+)")
	for i := 0; i < len(chunk); i++ {
		task := chunk[i]

		result, _ := json.Marshal(chunk[i])

		mat := r.FindAllSubmatch(result, len(result))

		if len(mat) > 0 {

			taskValue, _ := strconv.Atoi(string(mat[0][2]))
			total += taskValue

			if task.Fields.Status.StatusCategory.ID == 3 {
				done += taskValue
			}
		}

	}

	var percentage float64 = float64(done) / float64(total)

	return percentage
}
