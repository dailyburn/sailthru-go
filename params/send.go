package params

import "encoding/json"

type Send struct {
	Template     string            `json:"template"`
	Email        string            `json:"email"`
	Vars         map[string]string `json:"vars,omitempty"`
	Options      Options           `json:"options,omitempty"`
	ScheduleTime ScheduleTime      `json:"schedule_time,omitempty"`
}

func (s *Send) GetEndpoint() string {
	return "send"
}

type ScheduleTime struct {
	StartTime string
	EndTime   string
}

func (s *ScheduleTime) MarshalJSON() ([]byte, error) {
	if s.StartTime != "" && s.EndTime != "" {
		return json.Marshal(&struct {
			StartTime string `json:"start_time"`
			EndTime   string `json:"end_time"`
		}{
			StartTime: s.StartTime,
			EndTime:   s.EndTime,
		})
	}
	if s.StartTime != "" {
		return json.Marshal(s.StartTime)
	}
	return []byte(`""`), nil
}

type Options struct {
	Headers  Headers `json:"headers,omitempty"`
	ReplyTo  string  `json:"replyto,omitempty"`
	BehalfOf string  `json:"behalf_email,omitempty"`
	Test     int     `json:"test,omitempty"`
}

type Headers struct {
	CC  string `json:"Cc,omitempty"`
	BCC string `json:"Bcc,omitempty"`
}
