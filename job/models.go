package job

const (
	Endpoint  = "job"
	JobUpdate = "update"
	JobImport = "import"
)

type params struct {
	List              string `json:"list,omitempty"`
	Job               string `json:"job"`
	PostbackURL       string `json:"postback_url,omitempty"`
	ReportEmail       string `json:"report_email,omitempty"`
	IncludeSignupDate bool   `json:"signup_dates"`
	URL               string `json:"url"`
}

type Params struct {
	PostbackURL       string
	ReportEmail       string
	IncludeSignupDate bool
}
