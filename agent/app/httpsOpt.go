package app

type HttpsOpt struct {
	Object interface{} `json:"object"`
}

func (h *HttpsOpt) Analysis() (interface{}, map[string][]string, error) {
	return h.Object, nil, nil
}
