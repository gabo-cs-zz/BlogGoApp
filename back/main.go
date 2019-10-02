package main

import (
  "encoding/json"
  "flag"
  "fmt"
  "log"
  "net/http"
  "strconv"
  
  "github.com/gorilla/handlers"
)

var post Post

func main() {
  mux := http.NewServeMux()

  headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
  methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"})
  origins := handlers.AllowedOrigins([]string{"*"})

  migrate := flag.Bool("migrate", false, "Creates database tables.")
  flag.Parse()
  if *migrate {
    if err := MakeMigrations(); err != nil {
      log.Fatal(err)
    }
  }
  mux.HandleFunc("/", IndexHandler)
  mux.HandleFunc("/posts", PostsHandler)
  log.Println("Running on http://localhost:3000")
  http.ListenAndServe(":3000", handlers.CORS(headers, methods, origins)(mux))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello world")
}

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
  n := new(Post)
  posts, err := n.GetAll()
  if err != nil {
    http.Error(w, err.Error(), http.StatusNotFound)
    return
  } 
  j, err := json.Marshal(posts)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }
  w.WriteHeader(http.StatusOK)
  w.Write(j)
}

func CreatePostsHandler(w http.ResponseWriter, r *http.Request) {
  var post Post
  err := json.NewDecoder(r.Body).Decode(&post)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }
  err = post.Create()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.WriteHeader(http.StatusOK)
}


func UpdatePostsHandler(w http.ResponseWriter, r *http.Request) {
  var post Post
  err := json.NewDecoder(r.Body).Decode(&post)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }
  err = post.Update()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.WriteHeader(http.StatusOK)
}

func DeletePostsHandler(w http.ResponseWriter, r *http.Request) {
  idStr := r.URL.Query().Get("id")
  if idStr == "" {
    http.Error(w, "Query id is required", http.StatusBadRequest)
    return
  }
  id, err := strconv.Atoi(idStr)
  if err != nil {
    http.Error(w, "Query id must be a number", http.StatusBadRequest)
    return
  }
  var post Post
  err = post.Delete(id)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  
  w.WriteHeader(http.StatusOK)
}

func PostsHandler(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
  case http.MethodGet:
    GetPostsHandler(w, r)
  case http.MethodPost:
    CreatePostsHandler(w, r)
  case http.MethodPut:
    UpdatePostsHandler(w, r)
  case http.MethodDelete:
    DeletePostsHandler(w, r)
  default:
    http.Error(w, "Method not allowed", http.StatusBadRequest)
    return
}
}