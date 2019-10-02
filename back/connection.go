package main

import "database/sql"

var db *sql.DB

func GetConnection() *sql.DB {
  if db != nil {
    return db
  }

  var err error
  db, err = sql.Open("sqlite3", "data.sqlite")
  if err != nil {
    panic(err)
  }
  return db
}

func MakeMigrations() error {
  db := GetConnection()
  q := `CREATE TABLE IF NOT EXISTS posts (
          id INTEGER PRIMARY KEY AUTOINCREMENT,
          title VARCHAR(64) NULL,
          body TEXT NULL,
          created_at TIMESTAMP DEFAULT DATETIME,
          updated_at TIMESTAMP NOT NULL
        );`

  _, err := db.Exec(q)
  if err != nil {
    return err
  }
  return nil
}