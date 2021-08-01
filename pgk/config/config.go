package config

type AppConfig struct {
	ConnStr string
	Env     string
}

//NewConfig creates a new config an returns it
func NewConfig() *AppConfig {
	return &AppConfig{}
}
