package params

type Event struct {
	ID           string            `json:"id"`
	Key          string            `json:"key,omitempty"`
	Event        string            `json:"event"`
	ScheduleTime string            `json:"schedule_time,omitempty"`
	Vars         map[string]string `json:"vars,omitempty"`
}

func (e *Event) GetEndpoint() string {
	return "event"
}
