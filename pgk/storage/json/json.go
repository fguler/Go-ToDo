package json

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path"
	"sync"

	"github.com/fguler/goToDo/pgk/config"
	"github.com/fguler/goToDo/pgk/models"
)

var lock = &sync.Mutex{} // there is also sync.Once option

var jsonDB *JsonDB

/* // TaskSaver represents an interface that must be implemented to persist tasks
type TaskSaver interface {
	Save(t models.Task) error
} */

type JsonDB struct {
	path  string
	tasks []models.Task
}

// NewDB returns an instance of JsonDB, which is a singleton
func NewDB(conf *config.AppConfig) (*JsonDB, error) {

	if jsonDB == nil {

		lock.Lock()
		defer lock.Unlock()

		jsonDB = &JsonDB{
			path: conf.ConnStr,
		}
		//loads existing tasks into memory
		if err := jsonDB.loadFromFile(); err != nil {
			return nil, err
		}
		return jsonDB, nil
	}

	return jsonDB, nil

}

//loadFromFile load tasks from json file
func (db *JsonDB) loadFromFile() error {

	//if the json DB file does not exist create one
	_, err := os.Stat(db.path)
	if err != nil && os.IsNotExist(err) {
		createDBFile(db.path)
	}

	f, err := os.Open(db.path)
	if err != nil {
		return err
	}
	defer f.Close()

	dec := json.NewDecoder(f)

	// read open bracket
	if _, err = dec.Token(); err != nil {
		return err
	}

	for dec.More() {
		var t models.Task
		if err = dec.Decode(&t); err != nil {
			return err
		}
		db.tasks = append(db.tasks, t)
	}

	// read closing bracket
	if _, err = dec.Token(); err != nil {
		return err
	}

	return nil

}

//saveToFile turns the tasks slice into json and saves it to the file
func (db *JsonDB) saveToFile() error {

	j, err := json.MarshalIndent(db.tasks, " ", "   ")

	if err != nil {
		return err
	}

	f, err := os.Create(db.path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(j)
	if err != nil {
		return err
	}

	return nil

}

//Add adds new task to the database
func (db *JsonDB) Add(t models.Task) error {

	db.tasks = append(db.tasks, t)

	return db.saveToFile()
}

//FindByID returns the task with given ID
func (db *JsonDB) FindByID(id string) (models.Task, error) {

	var ta models.Task

	for _, t := range db.tasks {

		if t.Id == id {
			return t, nil
		}
	}

	return ta, errors.New("no task with given ID")

}

//FindAll return all tasks
func (db *JsonDB) FindAll() []models.Task {
	return db.tasks
}

//Update updates to given task
func (db *JsonDB) Update(t models.Task) error {

	for i, tk := range db.tasks {

		if tk.Id == t.Id {
			db.tasks[i] = t
			return nil
		}

	}

	return errors.New("there is no task to update")

}

// getDBPath return path JSON DB
func GetDBPath(pa string) string {
	// using the function
	mydir, err := os.Getwd()

	if err != nil {
		panic(err)
	}
	mydir = path.Join(mydir, pa)
	return mydir
}

func createDBFile(pa string) {

	f, err := os.Create(pa)
	if err != nil {
		log.Fatal("Can't create json DB file")
	}

	defer f.Close()

	_, err = f.Write([]byte("[]"))
	if err != nil {
		log.Fatal("Can't write to json DB file")
	}

}

//since := time.Now().Add(-24 * time.Hour)
