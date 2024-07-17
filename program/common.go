package program

type Program struct {
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Config      map[string]string `json:"config,omitempty"`
	Operation   Operation         `json:"operation,omitempty"`
}

type Operation struct {
	Start   string `json:"start,omitempty"`
	Stop    string `json:"stop,omitempty"`
	Restart string `json:"restart,omitempty"`
	Status  string `json:"status,omitempty"`
}
