package params

type UpdateJob struct {
	Job               string `json:"job"`
	PostbackURL       string `json:"postback_url,omitempty"`
	ReportEmail       string `json:"report_email,omitempty"`
	IncludeSignupDate bool   `json:"signup_dates"`
	URL               string `json:"url"`
}

func (u *UpdateJob) GetEndpoint() string {
	return "job"
}
