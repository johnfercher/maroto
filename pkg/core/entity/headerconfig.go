package entity

// HeaderConfig represents configs from a header.
type HeaderConfig struct {
	// StartPage define start page of the header, start from 0
	StartPage int
}

func (h *HeaderConfig) MakeValid() {
	if h.StartPage < 0 {
		h.StartPage = 0
	}
}

func (h *HeaderConfig) AppendMap(m map[string]interface{}) map[string]interface{} {
	if h.StartPage > 0 {
		m["config_header_start_page"] = h.StartPage
	}

	return m
}
