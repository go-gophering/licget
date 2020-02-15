package licget

type License struct {
	Key            string   `json:"key"`
	Name           string   `json:"name"`
	Spdx_id        string   `json:"spdx_id"`
	Url            string   `json:"url"`
	Node_id        string   `json:"node_id"`
	Html_url       string   `json:"html_url"`
	Description    string   `json:"description"`
	Implementation string   `json:"implementation"`
	Permissions    []string `json:"permissions"`
	Conditions     []string `json:"conditions"`
	Limitations    []string `json:"limitations"`
	Body           string   `json:"body"`
	Featured       bool     `json:"featured"`
}

type Licenses []License
