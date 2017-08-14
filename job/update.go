package job

import (
	"github.com/dailyburn/sailthru-go/client"
)

type Update struct {
	client *client.Client
	Params Params
}

func NewUpdate(client *client.Client) *Update {
	return &Update{client: client}
}

func (u *Update) ProcessURL(dataURL string) error {
	p := params{
		Job:               JobUpdate,
		URL:               dataURL,
		PostbackURL:       u.Params.PostbackURL,
		ReportEmail:       u.Params.ReportEmail,
		IncludeSignupDate: u.Params.IncludeSignupDate,
	}

	return u.client.Post(Endpoint, p)
}
