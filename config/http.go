package config

type Http struct {
	Addr   string `json:"addr,omitempty" yaml:"addr"`
	TlsCrt string `json:"tls_crt,omitempty" yaml:"tls_crt"`
	TlsKey string `json:"tls_key,omitempty" yaml:"tls_key"`
	Token  string `json:"token,omitempty" yaml:"token"`
}
