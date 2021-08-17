package main

import (
	"net"
	"net/http"
	"os"

	"github.com/fguler/goToDo/pgk/config"
	"github.com/fguler/goToDo/pgk/http/rest"
	"github.com/fguler/goToDo/pgk/storage/json"
	"github.com/fguler/goToDo/pgk/task"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type application struct {
	conf   *config.Config
	server *http.Server
}

func main() {

	app := application{}
	appConf := config.NewAppConfig()
	app.conf = appConf

	// load .env vars
	if err := godotenv.Load(".env"); err != nil {
		app.conf.ErrorLog.Fatal(err)
	}

	if err := app.start(); err != nil {
		app.conf.ErrorLog.Fatal(err)
	}

}

//start runs the server
func (app *application) start() error {

	address := net.JoinHostPort(os.Getenv("HOST"), os.Getenv("PORT"))

	db, err := json.NewDB(os.Getenv("DB_URL"))

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
	app.server = &svr

	app.conf.InfoLog.Printf("Starting aplication on %s \n", svr.Addr)
	err = svr.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

/*
func getEnvValue(envName string, defaultValue string) string {
	if val, ok := os.LookupEnv(envName); ok {
		return val
	}

	return defaultValue
}
*/
