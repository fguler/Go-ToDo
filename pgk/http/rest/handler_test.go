package rest_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/fguler/goToDo/pgk/config"
	"github.com/fguler/goToDo/pgk/http/rest"
	"github.com/fguler/goToDo/pgk/storage/json"
	"github.com/fguler/goToDo/pgk/task"
	"github.com/gorilla/mux"
)

var r *mux.Router

func TestMain(m *testing.M) {

	conf := config.NewConfig()
	conf.ConnStr = json.GetDBPath("db_test.json")

	db, err := json.NewDB(conf)

	if err != nil {
		log.Fatal("TestMain can't create a db_test.json file!")
	}

	ts := task.NewTaskService(db)

	r = mux.NewRouter()

	rest.RegisterRoutes(ts, r)

	exitVal := m.Run()
	log.Println("TODO: Clean up stuff after tests here!!!")
	os.Exit(exitVal)

}

func Test_AddTask(t *testing.T) {

	task := `{
		"topic":"Test 1",
		"status":false
	}`

	sr := strings.NewReader(task)

	req, err := http.NewRequest("POST", "/api/v1/task", sr)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Fatalf("Can't create request, got : %v", err)
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("handler fails, expected status code %d, but got : %d", http.StatusCreated, rr.Code)
	}

}
