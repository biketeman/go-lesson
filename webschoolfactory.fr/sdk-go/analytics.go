package analytics

type Settings struct {
	URL string `json:"url"`
}

type Analytics struct {
	Settings *Settings `json:"settings"`
}

func New(inputs *Settings) (*Analytics, error) {

	if inputs.URL == "" {
		inputs.URL = "http://localhost:1337"
	}

	return &Analytics{
		Settings: inputs,
	}, nil
}
