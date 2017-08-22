package params

type ImportJob struct {
	Job               string `json:"job"`
	List              string `json:"list"`
	PostbackURL       string `json:"postback_url,omitempty"`
	ReportEmail       string `json:"report_email,omitempty"`
	IncludeSignupDate bool   `json:"signup_dates"`
	URL               string `json:"url"`
}

func (i *ImportJob) GetEndpoint() string {
	return "job"
}
