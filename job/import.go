package sailthru_job

import (
	"github.com/dailyburn/sailthru-go/client"
)

type Import struct {
	client *sailthru_client.Client
	Params Params
}

func NewImport(client *sailthru_client.Client) *Import {
	return &Import{client: client}
}

func (i *Import) ProcessURL(list, dataURL string) error {
	p := params{
		List:              list,
		Job:               JobImport,
		URL:               dataURL,
		PostbackURL:       i.Params.PostbackURL,
		ReportEmail:       i.Params.ReportEmail,
		IncludeSignupDate: i.Params.IncludeSignupDate,
	}

	return i.client.Post(Endpoint, p)
}
