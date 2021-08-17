package config

import (
	"log"
	"os"
)

type Config struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

//NewAppConfig creates a app config an returns it
func NewAppConfig() *Config {

	ac := Config{}
	ac.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	ac.ErrorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return &ac
}
