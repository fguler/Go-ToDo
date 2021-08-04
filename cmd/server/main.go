package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/fguler/goToDo/pgk/config"
	"github.com/fguler/goToDo/pgk/http/rest"
	"github.com/fguler/goToDo/pgk/storage/json"
	"github.com/fguler/goToDo/pgk/task"
	"github.com/gorilla/mux"
)

func main() {

	conf := config.NewConfig()

	conf.ConnStr = getEnvValue("CONN_STRING", json.GetDBPath("/pgk/storage/json/db.json"))
	conf.Env = getEnvValue("ENV", "development")
	conf.Host = getEnvValue("HOST", "localhost")
	conf.Port = getEnvValue("PORT", "7070")

	if err := run(conf); err != nil {
		log.Fatal(err)
	}

}

func run(conf *config.AppConfig) error {

	address := net.JoinHostPort(conf.Host, conf.Port)

	db, err := json.NewDB(conf)
	if err != nil {
		return err
	}

	r := mux.NewRouter()

	ts := task.NewTaskService(db)

	rest.RegisterRoutes(ts, r)

	svr := http.Server{
		Handler: r,
		Addr:    address,
	}

	log.Printf("Starting aplication on %s \n", svr.Addr)
	err = svr.ListenAndServe()
	if err != nil {
		return err
	}

	return nil

}

func getEnvValue(envName string, defaultValue string) string {
	if val, ok := os.LookupEnv(envName); ok {
		return val
	}

	return defaultValue
}
