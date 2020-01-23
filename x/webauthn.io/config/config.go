package config

// Config represents the configuration information.
type Config struct {
	DBName       string `json:"db-name"       yaml:"db-name"       doc:"name is one of [sqlite3]`
	DBPath       string `json:"db-path"       yaml:"db-path"       doc:"path to the db"`
	HostAddress  string `json:"address"       yaml:"address"       doc:"host address to listen on."`
	RedirectPort string `json:"redirect-port" yaml:"redirect-port" doc:"redirects to secure port"`
	HostPort     string `json:"port"          yaml:"port"          doc:"tls secure port"`
	LogFile      string `json:"log-file"      yaml:"log-file"      doc:"log-file is an optional path"`
	RelyingParty string `json:"relying-party" yaml:"relying-party" doc:"webauthn relying party"`
	Cert         string `json:"cert"          yaml:"cert"          doc:"certificate file path"`
	Key          string `json:"key"           yaml:"key"           doc:"key file path"`
	Ca           string `json:"ca"            yaml:"ca"            doc:"cert authority file path"`
	Config       string `json:"config"        yaml:"config"        doc:"json or yaml config file"`
	Debug        bool   `json:"debug"         yaml:"debug"         doc:"be verbose"`
}
