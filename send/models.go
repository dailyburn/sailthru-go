package sailthru_send

const (
	Endpoint = "send"
)

// Ensure that the params hash is ordered alphabetically so that
// the signature creation process "just works"
type params struct {
	Template     string      `json:"template"`
	Email        string      `json:"email"`
	ScheduleTime interface{} `json:"schedule_time,omitempty"`
	Options      Options     `json:"options,omitempty"`
	DataFeedURL  string      `json:"data_feed_url,omitempty"`
}

type ScheduleWindow struct {
	StartTime string `json:"start_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
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

type Params struct {
	ScheduleTime string
	StartTime    string
	EndTime      string
	Vars         map[string]string
	CC           string
	BCC          string
	ReplyTo      string
	Test         bool
	DataFeedURL  string
	BehalfOf     string
}
