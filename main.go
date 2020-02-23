package main

import (
	"fmt"
	jira "github.com/andygrunwald/go-jira"
	"github.com/bryan-nice/jira-issue-creation/configuration"
	"github.com/pkg/errors"
	"log"
	"os"
)

func main() {

	var issues []jira.Issue
	var err error
	var config *configuration.Config
	var jiraJql string
	var result string

	// Init configuration from environment variables
	config = new(configuration.Config)
	err = config.Init()
	if err != nil {
		log.Fatalf("Exception: %v", err)
	}

	jiraApiToken := config.JiraApiToken
	jiraUsername := config.JiraUsername
	jiraAccountUrl := config.JiraAccountUrl
	jiraProject := config.JiraProject
	jiraIssueType := config.JiraIssueType
	jiraIssueSummary := config.JiraIssueSummary
	jiraIssueDescription := config.JiraIssueDescription
	jiraIssueAttachment := config.JiraIssueAttachment

	tp := jira.BasicAuthTransport{
		Username: jiraUsername,
		Password: jiraApiToken,
	}

	jiraClient, err := jira.NewClient(tp.Client(), jiraAccountUrl)
	if err != nil {
		log.Fatalf("Exception: %v", err)
	}

	// appendFunc will append jira issues to []jira.Issue
	appendFunc := func(i jira.Issue) (err error) {
		issues = append(issues, i)
		return err
	}

	// SearchPages will page through results and pass each issue to appendFunc
	// In this example, we'll search for all the issues in the target project
	jiraJql = fmt.Sprintf("project = %s AND summary ~ '%s'", jiraProject, jiraIssueSummary)
	err = jiraClient.Issue.SearchPages(jiraJql, nil, appendFunc)
	if err != nil {
		log.Printf("%+v", errors.Wrap(err, "Exception"))
	}

	if len(issues) == 0 {
		i := jira.Issue{
			Fields: &jira.IssueFields{
				Description: jiraIssueDescription,
				Type: jira.IssueType{
					Name: jiraIssueType,
				},
				Project: jira.Project{
					Key: jiraProject,
				},
				Summary: jiraIssueSummary,
			},
		}
		issue, _, err := jiraClient.Issue.Create(&i)
		if err != nil {
			log.Printf("%+v", errors.Wrap(err, "Exception"))
		}

		if len(jiraIssueAttachment) != 0 {
			f, err := os.Open(jiraIssueAttachment)
			if err != nil {
				log.Printf("%+v", errors.Wrap(err, "Exception"))
			}
			jiraClient.Issue.PostAttachment(issue.Key, f, jiraIssueAttachment)
		}

		result = fmt.Sprintf("%sbrowse/%s", jiraAccountUrl, string(issue.Key))
	} else {
		for _, issue := range issues {
			result = fmt.Sprintf("%sbrowse/%s", jiraAccountUrl, string(issue.Key))
		}
	}
	fmt.Printf(fmt.Sprintf("::set-output name=jira_issue_url::%s", result))
}
