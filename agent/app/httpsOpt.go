package app

type HttpsOpt struct {
	Object interface{} `json:"object"`
}

func (h *HttpsOpt) Analysis() ([]byte, map[string][]string, error) {
	return nil, nil, nil
}
