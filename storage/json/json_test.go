package json

import (
	"log"
	"os"
	"testing"

	"github.com/fguler/goToDo/pgk/config"
)

var dbTest *JsonDB

func TestMain(m *testing.M) {

	conf := config.NewConfig()
	conf.ConnStr = "./db_test.json"

	f, err := os.Create(conf.ConnStr)
	if err != nil {
		log.Fatal("TestMain can't create db_test.json file!")
	}
	defer f.Close()

	_, err = f.Write([]byte("[]"))
	if err != nil {
		log.Fatal("TestMain can't write to db_test.json file!")
	}

	dbTest, _ = NewDB(conf)

	exitVal := m.Run()

	err = os.Remove(conf.ConnStr)
	if err != nil {
		log.Fatal("TestMain can't remove db_test.json file!")
	}

	os.Exit(exitVal)

}

func Test_NewDB(t *testing.T) {
	if dbTest == nil {
		t.Fatal("Can't create database: expected type *JsonDB but got", nil)
	}
}
