package resumes

import (
  "log"
  "context"
  "net/http"
  "encoding/json"
  "go.mongodb.org/mongo-driver/bson"
  "github.com/trivektor/svelte-resume-builder/db"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "github.com/gorilla/mux"
)

const DATABASE_NAME = "svelte_resumes_builder"
const COLLECTION_NAME = "resumes"

type Section struct {
  Name string `json:"name"`
  Description string `json:"description"`
}

type Resume struct {
  ID primitive.ObjectID `bson:"_id" json:"id"` 
  Name string `json:"name"`
  Description string `json:"description"`
  Sections []Section `json:"sections"`
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

func Show(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  id := params["id"]
  var resume Resume
  docId, _ := primitive.ObjectIDFromHex(id)

  collection := db.Client.Database(DATABASE_NAME).Collection(COLLECTION_NAME)
  collection.FindOne(context.TODO(), bson.M{"_id": docId}).Decode(&resume)

  json, _ := json.Marshal(resume)

  w.Header().Set("Content-Type", "application/json")
  w.Write(json)
}
