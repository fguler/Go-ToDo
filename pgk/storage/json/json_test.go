package json

import (
	"log"
	"os"
	"testing"
)

var dbTest *JsonDB

func TestMain(m *testing.M) {

	dbUrl := GetDBPath("db_test.json")

	f, err := os.Create(dbUrl)
	if err != nil {
		log.Fatal("TestMain can't create db_test.json file!")
	}
	defer f.Close()

	_, err = f.Write([]byte("[]"))
	if err != nil {
		log.Fatal("TestMain can't write to db_test.json file!")
	}

	dbTest, _ = NewDB(dbUrl)

	exitVal := m.Run()

	err = os.Remove(dbUrl)
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
