package config

type AppConfig struct {
	ConnStr string
	Env     string
	Host    string
	Port    string
}

//NewConfig creates a new config an returns it
func NewConfig() *AppConfig {
	return &AppConfig{}
}
