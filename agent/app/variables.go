package app

var (
	ConfigData     Config
	HostRouterList = make(map[string]HttpRouter, 1000) // 主机路由表
)

type HttpRouter struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type ServerData struct {
	Type   string      `json:"type"`
	UUID   string      `json:"uuid"`
	Object interface{} `json:"object"`
}

type Agent struct {
	Network    string `json:"network"`
	Serverip   string `json:"serverip"`
	Serverport int    `json:"serverport"`
	Secret     string `json:"secret"`
}

type Mapping struct {
	Name      string `ini:"name"`
	RemoteURL string `ini:"remoteurl"`
	LocalPort int    `ini:"localport"`
	LocalIP   string `ini:"LocalIP"`
}

type Config struct {
	Agent      Agent   `ini:"agent"`
	Mapping    Mapping `ini:"mapping"`
	PKIXPubKey string  `json:"pkixpubkey"`
}
