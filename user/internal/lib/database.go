package lib

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"path/filepath"
	"runtime"
)

var DB = &sql.DB{}

func init() {
	var err error
	_, file, _, ok := runtime.Caller(0)

	sourcePath := ""
	if ok {
		sourcePath += filepath.Dir(file) + "/mydatabase.db"
	} else {
		sourcePath = "./mydatabase.db"
	}

	DB, err = sql.Open("sqlite3", sourcePath)
	if err != nil {
		log.Fatal(err)
	}
}
