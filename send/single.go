package sailthru_send

import (
	"github.com/dailyburn/sailthru-go/client"
)

type Single struct {
	client *sailthru_client.Client
	Params Params
}

func NewSingle(client *sailthru_client.Client) *Single {
	return &Single{client: client}
}

func (s *Single) Deliver(email string, template string, vars map[string]string) error {
	p := params{
		Template: template,
		Email:    email,
		Options: Options{
			Headers: Headers{
				CC:  s.Params.CC,
				BCC: s.Params.BCC,
			},
			ReplyTo:  s.Params.ReplyTo,
			BehalfOf: s.Params.BehalfOf,
		},
		DataFeedURL: s.Params.DataFeedURL,
	}

	if s.Params.Test {
		p.Options.Test = 1
	}

	if s.Params.StartTime == "" && s.Params.EndTime == "" {
		p.ScheduleTime = s.Params.ScheduleTime
	} else {
		p.ScheduleTime = ScheduleWindow{
			StartTime: s.Params.StartTime,
			EndTime:   s.Params.EndTime,
		}
	}

	return s.client.Post(Endpoint, p)
}
