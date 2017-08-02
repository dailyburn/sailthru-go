package sailthru_job

const (
	Endpoint  = "job"
	JobUpdate = "update"
)

// Ensure that the params hash is ordered alphabetically so that
// the signature creation process "just works"
type params struct {
	Job               string `json:"job"`
	PostbackURL       string `json:"postback_url,omitempty"`
	IncludeSignupDate bool   `json:"signup_dates"`
	URL               string `json:"url"`
}

type Params struct {
	PostbackURL       string
	IncludeSignupDate bool
}
