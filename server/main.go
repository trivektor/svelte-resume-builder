package main

import (
  "context"
  "net/http"
  "os"
  "github.com/gorilla/mux"
  "github.com/rs/cors"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "github.com/trivektor/svelte-resume-builder/controllers/resumes"
  "github.com/trivektor/svelte-resume-builder/db"
)

func main() {
  client, _ := mongo.NewClient(options.Client().ApplyURI(os.Getenv("CHATTYY_DB_CONNECTION_STRING")))
  client.Connect(context.Background())

  db.Client = client

  defer db.Client.Disconnect(context.Background())

  r := mux.NewRouter()
  r.HandleFunc("/resumes", resumes.Create).Methods("POST")
  r.HandleFunc("/resumes", resumes.Index)

  handler := cors.Default().Handler(r)
  http.ListenAndServe(":8080", handler)
}
