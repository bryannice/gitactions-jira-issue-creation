# Jira Issue Creation

Git Action to create an Issue in Jira. This action can be used to create an issue when a build fails in a Git Action workflow. 

## Usage 

This action can be used after any other action. Below is simple example on using it:

1\. Create a `.github/workflows/jira-issue-creation.yml`

2\. Add the following properties to `jira-issue-creation.yml` file

```yaml
on: push
name: Jira Issue Creation Demo
jobs:
  jiraIssueCreation:
    name: Jira Issue Creation Demo
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Jira Creation Demo
        uses: bryan-nice/jira-issue-creation@master
        env:
          JIRA_API_TOKEN: 'jiraApiToke'
          JIRA_USERNAME: 'user@email.com'
          JIRA_ACCOUNT_URL: 'https://someaccount.atlassian.net/'
          JIRA_PROJECT: 'jira-issue-creation'
          JIRA_ISSUE_TYPE: 'Demo''ing Jira Issue Creation'
          JIRA_ISSUE_DESCRIPTION: 'Demo''ing Jira Issue Creation'
          JIRA_ISSUE_SUMMARY: 'Demo''ing Jira Issue Creation'
```

Go [here](deployment/git-actions/template_jira_issue_creation.yml) for a template yml with all environment variables.

## Environment Variables
These are the environment variables that can be set to pass in additional information about the Git Action.

<table>
    <tr>
        <td>
            Variable Name
        </td>
        <td>
            Required
        </td>
        <td>
            Description
        </td>
    </tr>
    <tr>
        <td>
            JIRA_API_TOKEN
        </td>
        <td>
            Yes
        </td>
        <td>
            Jira API Token used instead of a password.
        </td>
    </tr>
    <tr>
        <td>
            JIRA_USERNAME
        </td>
        <td>
            Yes
        </td>
        <td>
            Jira user email.
        </td>
    </tr>
    <tr>
        <td>
            JIRA_ACCOUNT_URL
        </td>
        <td>
            Yes
        </td>
        <td>
            Base url to the Jira account.
        </td>
    </tr>
    <tr>
        <td>
            JIRA_PROJECT
        </td>
        <td>
            Yes
        </td>
        <td>
            Jira project the ticket will be filed under.
        </td>
    </tr>
    <tr>
        <td>
            JIRA_ISSUE_TYPE
        </td>
        <td>
            Yes
        </td>
        <td>
            Type of isse to be created [Bug|Task].
        </td>
    </tr>
    <tr>
        <td>
            JIRA_ISSUE_DESCRIPTION
        </td>
        <td>
            Yes
        </td>
        <td>
            Body of the Jira Issue.
        </td>
    </tr>
    <tr>
        <td>
            JIRA_ISSUE_SUMMARY
        </td>
        <td>
            Yes
        </td>
        <td>
            Title of the Jira Issue.
        </td>
    </tr>
</table>

## Reference
* [Creating an API toke for a user.](https://confluence.atlassian.com/cloud/api-tokens-938839638.html)
* [Jira Go Lang Library](https://github.com/andygrunwald/go-jira)

## License
[GPLv3](LICENSE)