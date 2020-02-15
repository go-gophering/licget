package licget

type Text struct {
	Media_type string `json:"media_type"`
	Title      string `json:"title"`
	Url        string `json:"url"`
}

type Identifier struct {
	Identifier string `json:"identifier"`
	Scheme     string `json:"scheme"`
}

type Link struct {
	Note string `json:"note"`
	Url  string `json:"url"`
}

type License struct {
	Id            string       `json:"id"`
	Identifiers   []Identifier `json:"identifiers"`
	Links         []Link       `json:"link"`
	Name          string       `json:"name"`
	Other_names   []string     `json:"other_names"`
	Suberseded_by string       `json:"suberseded_by"`
	Keywords      []string     `json:"keywords"`
	Text          Text         `json:"text"`
}
