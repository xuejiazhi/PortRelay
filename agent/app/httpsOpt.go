package app

type HttpsOpt struct {
	Object interface{} `json:"object"`
}

func (h *HttpsOpt) Analysis() (interface{}, error) {
	return h.Object, nil
}
