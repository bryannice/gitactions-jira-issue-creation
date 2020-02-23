package configuration

import (
	"errors"
	"os"
)

type Config struct {
	JiraApiToken         string
	JiraUsername         string
	JiraAccountUrl       string
	JiraProject          string
	JiraIssueType        string
	JiraIssueDescription string
	JiraIssueSummary     string
	JiraIssueAttachment  string
}

func (config *Config) Init() error {

	var err error

	// Checking and setting required environment variables
	if os.Getenv("JIRA_API_TOKEN") == "" {
		err = errors.New("JIRA_API_TOKEN must be set")
	} else {
		config.JiraApiToken = os.Getenv("JIRA_API_TOKEN")
	}

	if os.Getenv("JIRA_USERNAME") == "" {
		err = errors.New("JIRA_USERNAME must be set")
	} else {
		config.JiraUsername = os.Getenv("JIRA_USERNAME")
	}

	if os.Getenv("JIRA_ACCOUNT_URL") == "" {
		err = errors.New("JIRA_ACCOUNT_URL must be set")
	} else {
		config.JiraAccountUrl = os.Getenv("JIRA_ACCOUNT_URL")
	}

	if os.Getenv("JIRA_PROJECT") == "" {
		err = errors.New("JIRA_PROJECT must be set")
	} else {
		config.JiraProject = os.Getenv("JIRA_PROJECT")
	}

	if os.Getenv("JIRA_ISSUE_TYPE") == "" {
		err = errors.New("JIRA_ISSUE_TYPE must be set")
	} else {
		config.JiraIssueType = os.Getenv("JIRA_ISSUE_TYPE")
	}

	if os.Getenv("JIRA_ISSUE_DESCRIPTION") == "" {
		err = errors.New("JIRA_ISSUE_DESCRIPTION must be set")
	} else {
		config.JiraIssueDescription = os.Getenv("JIRA_ISSUE_DESCRIPTION")
	}

	if os.Getenv("JIRA_ISSUE_SUMMARY") == "" {
		err = errors.New("JIRA_ISSUE_SUMMARY must be set")
	} else {
		config.JiraIssueSummary = os.Getenv("JIRA_ISSUE_SUMMARY")
	}

	config.JiraIssueAttachment = os.Getenv("JIRA_ISSUE_ATTACHMENT")

	return err
}
