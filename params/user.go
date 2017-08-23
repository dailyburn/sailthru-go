package params

type User struct {
	ID     string         `json:"id"`
	Key    string         `json:"key,omitempty"`
	Fields map[string]int `json:"fields,omitempty"`
}

type SaveUser struct {
	User
	Keys            map[string]string `json:"keys,omitempty"`
	KeysConflict    string            `json:"keysconflict,omitempty"`
	Cookies         map[string]string `json:"cookies,omitempty"`
	Lists           map[string]int    `json:"lists,omitempty"`
	OptoutEmail     string            `json:"optout_email,omitempty"`
	OptoutTemplates map[string]int    `json:"optout_templates,omitempty"`
	Vars            map[string]string `json:"vars,omitempty"`
}

func (u *User) GetEndpoint() string {
	return "user"
}

func (u *SaveUser) GetEndpoint() string {
	return "user"
}
