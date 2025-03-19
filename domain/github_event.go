package domain

// GitHubEvent representa un evento de GitHub
type GitHubEvent struct {
	Action      string `json:"action"`
	PullRequest struct {
		Title   string `json:"title"`
		HTMLURL string `json:"html_url"`
		User    struct {
			Login string `json:"login"`
		} `json:"user"`
		Merged bool `json:"merged"`
	} `json:"pull_request"`
	WorkflowRun struct {
		Status     string `json:"status"`
		Conclusion string `json:"conclusion"`
		Name       string `json:"name"`
		HTMLURL    string `json:"html_url"`
	} `json:"workflow_run"`
}
