package app

// 超传输
func ProtoTransfer(proto string, data interface{}) ProtoInterface {
	switch proto {
	case "http":
		return &HttpOpt{
			Object: data,
		}
	case "https":
		return &HttpsOpt{
			Object: data,
		}
	default:
		return nil
	}
}

type ProtoInterface interface {
	Analysis() (interface{}, error)
}
