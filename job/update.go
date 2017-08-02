package sailthru_job

import (
	"encoding/json"

	"github.com/dailyburn/sailthru-go/client"
)

type Update struct {
	client *sailthru_client.Client
	Params Params
}

func NewUpdate(client *sailthru_client.Client) *Update {
	return &Update{client: client}
}

func (u *Update) ProcessURL(dataURL string) error {
	dd, err := json.Marshal(params{
		Job:               JobUpdate,
		URL:               dataURL,
		PostbackURL:       u.Params.PostbackURL,
		IncludeSignupDate: u.Params.IncludeSignupDate,
	})
	if err != nil {
		return err
	}

	return u.client.Post(Endpoint, dd)
}
