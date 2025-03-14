package entities

type ReviewEventPayload struct {
	Action      string      `json:"action"`
	Review      Review      `json:"review"`
	PullRequest PullRequest `json:"pull_request"`
	Repository  Repository  `json:"repository"`
}

type Review struct {
	ID       int    `json:"id"`
	Body     string `json:"body"`
	User     User   `json:"user"`
	State    string `json:"state"`
	CommitID string `json:"commit_id"`
}
