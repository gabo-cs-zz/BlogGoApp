package main

import (
  "errors"
  "time"

  _ "github.com/mattn/go-sqlite3"
)

type Post struct {
  ID          int       `json:"id,omitempty"`
  Title       string    `json:"title"`
  Body        string    `json:"body"`
  CreatedAt   time.Time `json:"created_at,omitempty"`
  UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func (n Post) Create() error {
  db := GetConnection()

  q := `INSERT INTO posts (title, body, updated_at)
      VALUES(?, ?, ?)`

  stmt, err := db.Prepare(q)
  if err != nil {
    return err
  }
  defer stmt.Close()

  r, err := stmt.Exec(n.Title, n.Body, time.Now())
  if err != nil {
    return err
  }

  if i, err := r.RowsAffected(); err != nil || i != 1 {
    return errors.New("ERROR: At least a row was meant to be changed.")
  }

  return nil
}

func (n *Post) GetAll() ([]Post, error) {
  db := GetConnection()
  q := `SELECT
      id, title, body, created_at, updated_at
      FROM posts`

  rows, err := db.Query(q)
  if err != nil {
    return []Post{}, err
  }

  defer rows.Close()

  posts := []Post{}

  for rows.Next() {
    rows.Scan(&n.ID, &n.Title, &n.Body, &n.CreatedAt, &n.UpdatedAt)
    posts = append(posts, *n)
  }
  return posts, nil
}

func (n *Post) GetByID(id int) (Post, error) {
  db := GetConnection()
  q := `SELECT
    id, title, body, created_at, updated_at
    FROM posts WHERE id=?`

  err := db.QueryRow(q, id).Scan(
    &n.ID, &n.Title, &n.Body, &n.CreatedAt, &n.UpdatedAt,
  )
  if err != nil {
    return Post{}, err
  }

  return *n, nil
}

func (n Post) Update() error {
  db := GetConnection()
  q := `UPDATE posts set title=?, body=?, updated_at=?
    WHERE id=?`
  stmt, err := db.Prepare(q)
  if err != nil {
    return err
  }
  defer stmt.Close()

  r, err := stmt.Exec(n.Title, n.Body, time.Now(), n.ID)
  if err != nil {
    return err
  }
  if i, err := r.RowsAffected(); err != nil || i != 1 {
    return errors.New("ERROR: At least a row was meant to be changed.")
  }
  return nil
}

func (n Post) Delete(id int) error {
  db := GetConnection()

  q := `DELETE FROM posts
    WHERE id=?`
  stmt, err := db.Prepare(q)
  if err != nil {
    return err
  }
  defer stmt.Close()

  r, err := stmt.Exec(id)
  if err != nil {
    return err
  }
  if i, err := r.RowsAffected(); err != nil || i != 1 {
    return errors.New("ERROR: At least a row was meant to be changed.")
  }
  return nil
}