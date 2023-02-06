package useCase

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"

	"github.com/andygrunwald/go-jira"
)

type Task struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	PlannerCode int    `json:"plannerCode"`
	Url         string `json:"url"`
	Stage       string `json:"status"`
}

func GetTdTasks(client jira.Client, project string) []Task {
	jql := fmt.Sprintf("project = \"%s\" and cf[10020] in openSprints() and updatedDate > startOfDay(-7)", project)

	chunk, _, _ := client.Issue.Search(jql, &jira.SearchOptions{})

	tasks := make([]Task, 0)

	r, error := regexp.Compile("(PLN-)([0-9]+)")

	if error != nil {
		panic("regex have failed")
	}

	for _, item := range chunk {

		stringified, error := json.Marshal(item)

		if error != nil {
			panic("not stringified")
		}

		plannerSearch := r.FindAllSubmatch(stringified, len(stringified))

		var plannerCode int = 0

		if len(plannerSearch) > 0 && len(plannerSearch[0]) > 0 {

			taskValue, error := strconv.Atoi(string(plannerSearch[0][2]))

			if error != nil {
				panic("error on converting string to int")
			}

			plannerCode = taskValue
		}

		tasks = append(tasks, Task{
			Key:         item.Key,
			Name:        item.Fields.Description,
			PlannerCode: plannerCode,
			Url:         item.Self,
			Stage:       item.Fields.Status.StatusCategory.Name,
		})

	}

	return tasks
}
