package value_objects

type WorkflowRunEvent struct {
	Action      string      `json:"action"`
	WorkflowRun WorkflowRun `json:"workflow_run"`
	Repository  Repository `json:"repository"`
}

type WorkflowRun struct {
	Name      string `json:"name"`
	URL     string `json:"url"`
	Conclusion string `json:"conclusion"`
	HTMLURL   string `json:"html_url"`
}

type Repository struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
}