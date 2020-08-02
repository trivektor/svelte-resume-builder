package main

import (
  "context"
  "net/http"
  "os"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "github.com/trivektor/svelte-resume-builder/controllers/resumes"
  "github.com/trivektor/svelte-resume-builder/db"
)

func main() {
  client, _ := mongo.NewClient(options.Client().ApplyURI(os.Getenv("CHATTYY_DB_CONNECTION_STRING")))
  client.Connect(context.TODO())

  db.Client = client

  defer db.Client.Disconnect(context.TODO())

  r := mux.NewRouter()
  r.HandleFunc("/resumes", resumes.Create).Methods("POST")
  r.HandleFunc("/resumes", resumes.Index)
  r.HandleFunc("/resumes/{id}", resumes.Update).Methods("PUT")
  r.HandleFunc("/resumes/{id}", resumes.Delete).Methods("DELETE")
  r.HandleFunc("/resumes/{id}", resumes.Show)

  // https://stackoverflow.com/questions/40985920/making-golang-gorilla-cors-handler-work
  headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin"})
  originsOk := handlers.AllowedOrigins([]string{"*"})
  methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

  http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r))
}
