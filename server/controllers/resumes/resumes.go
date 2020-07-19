package resumes

import (
  "log"
  "context"
  "net/http"
  "encoding/json"
  "go.mongodb.org/mongo-driver/bson"
  "github.com/trivektor/svelte-resume-builder/db"
  "go.mongodb.org/mongo-driver/bson/primitive"
)

const DATABASE_NAME = "svelte_resumes_builder"
const COLLECTION_NAME = "resumes"

type Resume struct {
  Name string `json:name`
  Description string `json:description`
}

func Index(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Resumes"))
}

func Create(w http.ResponseWriter, r *http.Request) {
  var resume Resume

  json.NewDecoder(r.Body).Decode(&resume)

  collection := db.Client.Database(DATABASE_NAME).Collection(COLLECTION_NAME)

  res, err := collection.InsertOne(context.TODO(), bson.M{"name": resume.Name, "description": resume.Description})

  if err != nil {
    log.Fatal(err)
  }
  id := res.InsertedID.(primitive.ObjectID)

  m := make(map[string]string)
  m["id"] = id.Hex()
  result, _ := json.Marshal(m)

  w.Header().Set("Content-Type", "application/json")
  w.Write(result)
}
