package main
type config struct {
	apiPort       string
	filename      string
	router        string
	domain        string
	service       string
	certResolver  string
	entryPoint    string
}

type Routers struct {
	HTTP struct {
		Routers map[string]Route `yaml:"routers"`
	} `yaml:"http"`
}

type Route struct{
	EntryPoints []string `yaml:"entryPoints"`
	Service     string   `yaml:"service"`
	Rule        string   `yaml:"rule"`
	TLS         struct {
		CertResolver string `yaml:"certResolver"`
	} `yaml:"tls"`
}

type requestBody struct {
	Subdomain string `json:"subdomain"`
}